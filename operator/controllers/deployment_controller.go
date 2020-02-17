/*
Copyright 2019 The Alameda Authors.

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
	"time"

	"golang.org/x/net/context"

	appsv1 "k8s.io/api/apps/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DeploymentReconciler reconciles a Deployment object
type DeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ClusterUID string
}

func (r *DeploymentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	requeueDuration := 1 * time.Second

	instance := appsv1.Deployment{}
	err := r.Get(context.Background(), req.NamespacedName, &instance)
	if err != nil && k8sErrors.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		scope.Warnf("Get Deployment(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	deployment := appsv1.Deployment{}
	instance.DeepCopyInto(&deployment)
	scope.Infof("Reconciling Deployment(%s/%s).", req.Namespace, req.Name)

	name, err := GetAlamedaScalerControllerName(context.TODO(), r.Client, deployment.ObjectMeta)
	if err != nil {
		scope.Warnf("Get AlamedaScaler controller name and type failed: %+v", err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}, nil
	}
	setAlamedaScalerControllerName(&deployment.ObjectMeta, name)
	if err := r.Client.Update(context.TODO(), &deployment); err != nil {
		scope.Errorf("Update Deployment(%s/%s) falied: %s", deployment.GetNamespace(), deployment.GetName(), err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}, nil
	}

	return ctrl.Result{}, nil
}

func (r *DeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		Complete(r)
}
