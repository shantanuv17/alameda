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

	tenantv1alpha1 "github.com/containers-ai/alameda/operator/apis/tenant/v1alpha1"
	datahubtenant "github.com/containers-ai/alameda/operator/datahub/api/tenant"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AlamedaOrganizationReconciler reconciles a AlamedaOrganization object
type AlamedaOrganizationReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DatahubClient *datahubpkg.Client
}

var (
	scope = logUtil.RegisterScope("alamedaorganization_controllers", "alamedaorganization controllers", 0)
)

func (r *AlamedaOrganizationReconciler) updateDatahubByOrg(
	org *tenantv1alpha1.AlamedaOrganization) error {
	if err := datahubtenant.DeleteOrganization(r.DatahubClient, org.GetName()); err != nil {
		return err
	}
	if err := datahubtenant.CreateOrganization(r.DatahubClient, org); err != nil {
		return err
	}
	return nil
}

// +kubebuilder:rbac:groups=tenant.containers.ai,resources=alamedaorganizations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tenant.containers.ai,resources=alamedaorganizations/status,verbs=get;update;patch

func (r *AlamedaOrganizationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.TODO()
	instance := &tenantv1alpha1.AlamedaOrganization{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: req.Namespace, Name: req.Name,
	}, instance)
	if err != nil && k8sErrors.IsNotFound(err) {
		if err := datahubtenant.DeleteOrganization(r.DatahubClient, req.Name); err != nil {
			scope.Errorf("Remove AlamedaOrganization %s from datahub failed: %s",
				req.Name, err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		scope.Errorf("Get AlamedaOrganization %s failed: %s",
			req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	if err := r.updateDatahubByOrg(instance); err != nil {
		scope.Errorf("Update datahub for AlamedaOrganization %s failed: %s",
			req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	return ctrl.Result{}, nil
}

func (r *AlamedaOrganizationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tenantv1alpha1.AlamedaOrganization{}).
		Complete(r)
}
