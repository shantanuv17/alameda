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

	machinesetrepository "github.com/containers-ai/alameda/operator/datahub/client/machineset"
	"github.com/containers-ai/alameda/operator/pkg/machineset"
	datahubschemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MachineSetReconciler reconciles a MachineSet object
type MachineSetReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	DatahubMachineSetRepo          *machinesetrepository.MachineSetRepository
	DatahubCAMachineSetSchema      datahubschemas.Schema
	DatahubCAMachineSetMeasurement datahubschemas.Measurement
	ClusterUID                     string
	ReconcileTimeout               time.Duration
}

func (r *MachineSetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ReconcileTimeout)
	defer cancel()
	mss := []machineset.MachineSet{
		machineset.MachineSet{
			ClusterName: r.ClusterUID,
			ResourceMeta: machineset.ResourceMeta{
				KubernetesMeta: machineset.KubernetesMeta{
					Namespace: req.Namespace,
					Name:      req.Name,
				},
			},
		},
	}

	instance := mahcinev1beta1.MachineSet{}
	err := r.Get(context.Background(), req.NamespacedName, &instance)
	if err != nil && k8serrors.IsNotFound(err) {
		err = r.DatahubMachineSetRepo.DeleteMachineSets(ctx, mss)
		if err != nil {
			scope.Errorf("Delete MachineSet (%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	} else if err != nil {
		scope.Warnf("Get MachineSet(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	return ctrl.Result{}, nil
}

func (r *MachineSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mahcinev1beta1.MachineSet{}).
		Complete(r)
}
