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
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"

	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
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
	if r.isMachineScaleUping(req.Name) {
		if updateErr := r.updateScaleUpExecutionNodeName(
			req.Namespace, machineSetName, req.Name); updateErr != nil {
			scope.Errorf("update execution failed for machine (%s/%s): %s",
				req.Namespace, req.Name, updateErr.Error())
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

func (r *MachineReconciler) updateScaleUpExecutionNodeName(
	msNS, msName, machineName string) error {
	msExecution := []entities.ExecutionClusterAutoscalerMachineset{}

	err := r.DatahubClient.ListTS(&msExecution, &datahubpkg.TimeRange{
		Order: datahubpkg.Desc,
		Limit: 1,
	}, nil, nil, datahubpkg.Option{
		Entity: entities.ExecutionClusterAutoscalerMachineset{
			ClusterName: r.ClusterUID,
			Namespace:   msNS,
			Name:        msName,
			NodeName:    "",
		},
		Fields: []string{"ClusterName", "Namespace", "Name", "NodeName"},
	})
	if err != nil {
		return err
	}

	if len(msExecution) == 1 && msExecution[0].ReplicasFrom < msExecution[0].ReplicasTo {
		msExecution[0].NodeName = machineName
		if err := r.DatahubClient.Create(&msExecution); err != nil {
			return err
		} else {
			scope.Infof("update node name %s to machineset (%s/%s) execution (time: %d) successfully",
				machineName, msNS, msName, msExecution[0].Time.Unix())
		}
	}
	return nil
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
	nodes := []entities.ResourceClusterStatusNode{}
	err := r.DatahubClient.List(&nodes, datahubpkg.Option{
		Entity: entities.ResourceClusterStatusNode{
			ClusterName: r.ClusterUID,
			Name:        machineName,
		},
		Fields: []string{"ClusterName", "Name"},
	})
	if err == nil && len(nodes) == 0 {
		return true
	}
	return false
}
