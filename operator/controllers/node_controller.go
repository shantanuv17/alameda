/*
Copyright 2020 The Alameda Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"time"

	datahub_node "github.com/containers-ai/alameda/operator/datahub/client/node"
	nodeinfo "github.com/containers-ai/alameda/operator/pkg/nodeinfo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	conn            *grpc.ClientConn
	DatahubNodeRepo datahub_node.AlamedaNodeRepository
	DatahubClient   *datahubpkg.Client

	Cloudprovider string
	RegionName    string
	ClusterUID    string
}

// Reconcile reads that state of the cluster for a Node object and makes changes based on the state read
// and what is in the Node.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=nodes/status,verbs=get;update;patch
func (r *NodeReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	requeueInterval := 3 * time.Second
	instance := &corev1.Node{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)

	nodeIsDeleted := false
	if err != nil && k8sErrors.IsNotFound(err) {
		nodeIsDeleted = true
		instance.Namespace = request.Namespace
		instance.Name = request.Name
	} else if err != nil {
		scope.Error(err.Error())
	}

	datahubNode, err := r.getNodeFromDatahub(request.Name)
	if err != nil {
		scope.Errorf("Get node %s from Datahub failed: %s", request.Name, err.Error())
		return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
	}

	msExecution := []entities.ExecutionClusterAutoscalerMachineset{}
	if datahubNode != nil {
		err = r.DatahubClient.ListTS(&msExecution, &datahubpkg.TimeRange{
			Order: datahubpkg.Desc,
			Limit: 1,
		}, nil, nil, datahubpkg.Option{
			Entity: entities.ExecutionClusterAutoscalerMachineset{
				ClusterName: r.ClusterUID,
				Namespace:   datahubNode.MachinesetNamespace,
				Name:        datahubNode.MachinesetName,
			},
			Fields: []string{"ClusterName", "Namespace", "Name"},
		})
		if err != nil {
			scope.Errorf("Get last execution of machineset %s/%s for node %s from Datahub failed: %s",
				datahubNode.MachinesetNamespace, datahubNode.MachinesetName, request.Name, err.Error())
			return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
		}
	}

	nodes := make([]*corev1.Node, 1)
	nodes[0] = instance

	if nodeIsDeleted {
		if len(msExecution) == 1 && msExecution[0].NodeName == "" &&
			msExecution[0].ReplicasFrom > msExecution[0].ReplicasTo {
			msExecution[0].DeltaDownTime = time.Now().Unix() - msExecution[0].Time.Unix()
			msExecution[0].NodeName = datahubNode.Name
			if err := r.DatahubClient.Create(&msExecution); err != nil {
				scope.Errorf(
					"Update delta down time for machineset %s/%s at execution time %v for node %s failed: %s",
					datahubNode.MachinesetNamespace, datahubNode.MachinesetName,
					msExecution[0].Time, request.Name, err.Error())
				return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
			} else {
				scope.Infof(
					"Update delta down time to %v seconds for machineset %s/%s at execution time %v for node %s",
					msExecution[0].DeltaDownTime, datahubNode.MachinesetNamespace,
					datahubNode.MachinesetName, msExecution[0].Time, request.Name)
			}
		}
		if err := r.deleteNodesFromDatahub(nodes); err != nil {
			scope.Errorf("Delete nodes from Datahub failed: %s", err.Error())
			return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
		}
	} else {
		if err := r.createNodesToDatahub(nodes); err != nil {
			scope.Errorf("Create node to Datahub failed failed: %s", err.Error())
			return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
		}
		if len(msExecution) == 1 && msExecution[0].NodeName == request.Name &&
			msExecution[0].ReplicasFrom < msExecution[0].ReplicasTo {
			if err = r.updateExecutionDeltaUpTime(msExecution[0], datahubNode); err != nil {
				scope.Errorf("Update delta up time for node %s failed failed: %s",
					request.Name, err.Error())
				return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
			}
		}
	}
	return reconcile.Result{}, nil
}

func (r *NodeReconciler) getNodeFromDatahub(nodeName string) (
	*entities.ResourceClusterStatusNode, error) {
	nodes := []entities.ResourceClusterStatusNode{}
	err := r.DatahubClient.List(&nodes, datahubpkg.Option{
		Entity: entities.ResourceClusterStatusNode{
			ClusterName: r.ClusterUID,
			Name:        nodeName,
		},
		Fields: []string{"ClusterName", "Name"},
	})
	if len(nodes) == 0 {
		return nil, err
	}
	return &nodes[0], err
}

func (r *NodeReconciler) createNodesToDatahub(nodes []*corev1.Node) error {

	nodeInfos, err := r.createNodeInfos(nodes)
	if err != nil {
		return errors.Wrap(err, "create nodeInfos failed")
	}

	return r.DatahubClient.Create(&nodeInfos)
}

func (r *NodeReconciler) deleteNodesFromDatahub(nodes []*corev1.Node) error {

	nodeInfos, err := r.createNodeInfos(nodes)
	if err != nil {
		return errors.Wrap(err, "create nodeInfos failed")
	}
	return r.DatahubClient.Delete(&nodeInfos)
}

func (r *NodeReconciler) createNodeInfos(nodes []*corev1.Node) (
	[]entities.ResourceClusterStatusNode, error) {
	nodeInfos := make([]entities.ResourceClusterStatusNode, len(nodes))
	for i, node := range nodes {
		n, err := r.createNodeInfo(node)
		if err != nil {
			return nodeInfos, errors.Wrap(err, "create nodeInfos failed")
		}
		nodeInfos[i] = *n
	}
	return nodeInfos, nil
}

func (r *NodeReconciler) createNodeInfo(node *corev1.Node) (
	*entities.ResourceClusterStatusNode, error) {
	n, err := nodeinfo.NewNodeInfo(*node, r.Client, r.ClusterUID)
	if err != nil {
		return nil, errors.Wrap(err, "new NodeInfo failed")
	}

	// set default info for node
	if n.IOProvider == "" {
		n.IOProvider = r.Cloudprovider
	}
	if n.IORegion == "" {
		n.IORegion = r.RegionName
	}
	return &n, nil
}

func (r *NodeReconciler) updateExecutionDeltaUpTime(
	theExecution entities.ExecutionClusterAutoscalerMachineset,
	datahubNode *entities.ResourceClusterStatusNode) error {
	deltaUpTimeMax := viper.GetInt64("ca.deltaUpTimeMax")
	theExecution.DeltaUpTime = datahubNode.CreateTime - datahubNode.MachineCreateTime
	if theExecution.DeltaUpTime > deltaUpTimeMax {
		scope.Infof(
			"Delta up time of node %s is %d seconds is larger than max delta up time limitation %d seconds",
			datahubNode.Name, theExecution.DeltaUpTime, deltaUpTimeMax)
		scope.Infof("Delta up time of node %s is set from %d seconds to %d seconds",
			datahubNode.Name, theExecution.DeltaUpTime, deltaUpTimeMax)
		theExecution.DeltaUpTime = deltaUpTimeMax
	}
	if err := r.DatahubClient.Create(
		&[]entities.ExecutionClusterAutoscalerMachineset{theExecution}); err != nil {
		return err
	}
	return nil
}

func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Node{}).
		Complete(r)
}
