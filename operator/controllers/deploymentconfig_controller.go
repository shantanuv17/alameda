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
	"context"
	"time"

	appsv1 "github.com/openshift/api/apps/v1"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	deploymentConfigFirstSynced = false
)

// DeploymentConfigReconciler reconciles a DeploymentConfig object
type DeploymentConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ClusterUID string
}

func (r *DeploymentConfigReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	requeueDuration := 1 * time.Second
	instance := appsv1.DeploymentConfig{}
	err := r.Get(context.Background(), req.NamespacedName, &instance)
	if err != nil && k8serrors.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		scope.Warnf("Get DeploymentConfig(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	deploymentConfig := appsv1.DeploymentConfig{}
	instance.DeepCopyInto(&deploymentConfig)
	scope.Infof("Reconciling DeploymentConfig(%s/%s).", req.Namespace, req.Name)

	name, err := GetAlamedaScalerControllerName(context.TODO(), r.Client, deploymentConfig.ObjectMeta)
	if err != nil {
		scope.Warnf("Get AlamedaScaler controller name and type failed: %+v", err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}, nil
	}
	setAlamedaScalerControllerName(&deploymentConfig.ObjectMeta, name)
	if err := r.Client.Update(context.TODO(), &deploymentConfig); err != nil {
		scope.Errorf("Update DeploymentConfig(%s/%s) falied: %s", deploymentConfig.GetNamespace(), deploymentConfig.GetName(), err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}, nil
	}

	return ctrl.Result{}, nil
}

func (r *DeploymentConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.DeploymentConfig{}).
		Complete(r)
}
