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

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	ca_client "github.com/containers-ai/alameda/operator/datahub/client/ca"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MachineReconciler reconciles a Machine object
type MachineReconciler struct {
	client.Client
	Scheme           *runtime.Scheme
	ClusterUID       string
	ReconcileTimeout time.Duration
	DatahubClient    *datahubpkg.Client
}

func (r *MachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	machineIns := mahcinev1beta1.Machine{}
	err := r.Get(context.Background(), req.NamespacedName, &machineIns)
	if err != nil && k8serrors.IsNotFound(err) {
		// not ready node removed does not trigger node reconcile function
		// delete node here to remove not ready node record
		delErr := r.DatahubClient.DeleteByOpts(
			&entities.ResourceClusterStatusNode{}, datahubpkg.Option{
				Entity: entities.ResourceClusterStatusNode{
					ClusterName: r.ClusterUID,
					Name:        req.Name,
				},
				Fields: []string{"ClusterName", "Name"},
			})
		if delErr != nil {
			scope.Errorf("remove node %s from datahub failed: %s", req.Name, delErr.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		return ctrl.Result{}, nil
	} else if err != nil {
		scope.Warnf("Get Machine(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	machineSetName := r.getMachineSetName(&machineIns)
	if machineSetName == "" {
		scope.Infof("machine (%s/%s) is not created by machineset", req.Namespace, req.Name)
		return ctrl.Result{}, nil
	}
	machinesetEntities := []entities.ResourceClusterAutoscalerMachineset{}
	if err := r.DatahubClient.List(&machinesetEntities, datahubpkg.Option{
		Entity: entities.ResourceClusterAutoscalerMachineset{
			ClusterName: r.ClusterUID,
			Namespace:   req.Namespace,
			Name:        machineSetName,
		},
		Fields: []string{"ClusterName", "Namespace", "Name"},
	}); err != nil {
		scope.Errorf("get machineset (%s/%s) failed", req.Namespace, machineSetName)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	if len(machinesetEntities) == 0 {
		scope.Infof("machine (%s/%s) is not watched by scaler", req.Namespace, req.Name)
		return ctrl.Result{}, nil
	}

	hasScaleUpExecution, hasScaleUpErr := r.hasScaleUpExecution(req.Name)
	if hasScaleUpErr != nil {
		scope.Errorf("scale up execution check for node %s failed: %s",
			req.Name, hasScaleUpErr.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	isScalingUp := r.isMachineScaleUping(req.Name)
	if isScalingUp && !hasScaleUpExecution {
		execTime := machineIns.GetCreationTimestamp().Time
		if execErr := r.sendExecutionTime(&execTime,
			req.Namespace, machineSetName, req.Name, isScalingUp); execErr != nil {
			scope.Errorf("add execution (scale up) failed for machine (%s/%s): %s",
				req.Namespace, req.Name, execErr.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
	}

	return ctrl.Result{}, nil
}

func (r *MachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mahcinev1beta1.Machine{}).
		Complete(r)
}

func (r *MachineReconciler) getMachineSetName(machine *mahcinev1beta1.Machine) string {
	for idx := range machine.OwnerReferences {
		if machine.OwnerReferences[idx].Kind == "MachineSet" {
			return machine.OwnerReferences[idx].Name
		}
	}
	return ""
}

// machine is created by machineset
func (r *MachineReconciler) isMachineScaleUping(machineName string) bool {

	nodeIns := corev1.Node{}
	err := r.Get(context.Background(), client.ObjectKey{
		Name: machineName,
	}, &nodeIns)
	if err != nil && k8serrors.IsNotFound(err) {
		executionEntity := []entities.ExecutionClusterAutoscalerMachineset{}
		err := r.DatahubClient.List(&executionEntity, datahubpkg.Option{
			Entity: entities.ExecutionClusterAutoscalerMachineset{
				ClusterName: r.ClusterUID,
				NodeName:    machineName,
			},
			Fields: []string{"ClusterName", "NodeName"},
		})
		if err != nil {
			scope.Errorf("cannot check is node %s scaling up: %s", machineName, err.Error())
			return false
		}
		if err == nil && len(executionEntity) == 0 {
			return true
		}
	} else if err != nil {
		scope.Errorf("cannot check is node %s scaling up: %s", machineName, err.Error())
		return false
	}
	return false
}

func (r *MachineReconciler) sendExecutionTime(execTime *time.Time,
	machinesetNamespace, machinesetName, machineName string, isScalingUp bool) error {
	return ca_client.SendExecutionTime(r.ClusterUID, execTime, r.Client,
		r.DatahubClient, machinesetNamespace, machinesetName, machineName, isScalingUp)
}

func (r *MachineReconciler) hasScaleUpExecution(machineName string) (bool, error) {
	msExecution := []entities.ExecutionClusterAutoscalerMachineset{}
	err := r.DatahubClient.ListTS(&msExecution, &datahubpkg.TimeRange{
		Order: datahubpkg.Desc,
	}, nil, nil, datahubpkg.Option{
		Entity: entities.ExecutionClusterAutoscalerMachineset{
			ClusterName: r.ClusterUID,
			NodeName:    machineName,
		},
		Fields: []string{"ClusterName", "NodeName"},
	})
	if err != nil {
		scope.Errorf("Get executions of for node %s from Datahub failed: %s",
			machineName, err.Error())
		return false, err
	}
	for idx := range msExecution {
		if msExecution[idx].ReplicasFrom < msExecution[idx].ReplicasTo {
			return true, nil
		}
	}
	return false, nil
}
