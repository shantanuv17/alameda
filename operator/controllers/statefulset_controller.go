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

	appsv1 "k8s.io/api/apps/v1"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// StatefulSetReconciler reconciles a StatefulSet object
type StatefulSetReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ClusterUID string
}

func (r *StatefulSetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	requeueDuration := 1 * time.Second

	instance := appsv1.StatefulSet{}
	err := r.Get(context.Background(), req.NamespacedName, &instance)
	if err != nil && k8serrors.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		scope.Warnf("Get StatefulSet(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	statefulSet := appsv1.StatefulSet{}
	instance.DeepCopyInto(&statefulSet)
	scope.Infof("Reconciling StatefulSet(%s/%s).", req.Namespace, req.Name)

	name, err := GetAlamedaScalerControllerName(context.TODO(), r.Client, statefulSet.ObjectMeta)
	if err != nil {
		scope.Warnf("Get AlamedaScaler controller name and type failed: %+v", err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}, nil
	}
	setAlamedaScalerControllerName(&statefulSet.ObjectMeta, name)
	if err := r.Client.Update(context.TODO(), &statefulSet); err != nil {
		scope.Errorf("Update StatefulSet(%s/%s) falied: %s", statefulSet.GetNamespace(), statefulSet.GetName(), err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueDuration}, nil
	}

	return ctrl.Result{}, nil
}

func (r *StatefulSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.StatefulSet{}).
		Complete(r)
}
