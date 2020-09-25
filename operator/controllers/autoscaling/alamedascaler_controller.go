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

package autoscaling

import (
	"context"
	"time"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"prophetstor.com/alameda/internal/pkg/message-queue/kafka"
	autoscalingv1alpha2 "prophetstor.com/alameda/operator/apis/autoscaling/v1alpha2"
	datahubscaler "prophetstor.com/alameda/operator/datahub/api/scaler"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	logUtil "prophetstor.com/alameda/pkg/utils/log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AlamedaScalerReconciler reconciles a AlamedaScaler object
type AlamedaScalerReconciler struct {
	client.Client
	EnabledDA     bool
	DatahubClient *datahubpkg.Client
	Scheme        *runtime.Scheme
	IsOpenshift   bool
	KafkaClient   kafka.Client
}

var (
	scope        = logUtil.RegisterScope("alamedascaler_controllers", "alamedascaler controllers", 0)
	requeueAfter = 3 * time.Second
)

func (r *AlamedaScalerReconciler) updateDatahubByScaler(
	scaler *autoscalingv1alpha2.AlamedaScaler) error {

	if err := datahubscaler.DeleteV1Alpha2Scaler(r.DatahubClient, r.Client,
		scaler.GetNamespace(), scaler.GetName(), r.EnabledDA); err != nil {
		return err
	}
	if err := datahubscaler.CreateV1Alpha2Scaler(r.DatahubClient, r.Client,
		r.KafkaClient, scaler, r.EnabledDA, r.IsOpenshift); err != nil {
		return err
	}
	return nil
}

// +kubebuilder:rbac:groups=autoscaling.containers.ai,resources=alamedascalers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=autoscaling.containers.ai,resources=alamedascalers/status,verbs=get;update;patch

func (r *AlamedaScalerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {

	ctx := context.TODO()
	instance := &autoscalingv1alpha2.AlamedaScaler{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: req.Namespace, Name: req.Name,
	}, instance)
	if err != nil && k8sErrors.IsNotFound(err) {
		if err := datahubscaler.DeleteV1Alpha2Scaler(r.DatahubClient, r.Client,
			req.Namespace, req.Name, r.EnabledDA); err != nil {
			scope.Errorf("Remove AlamedaScaler(%s/%s) from datahub failed: %s",
				req.Namespace, req.Name, err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		scope.Errorf("Get AlamedaScaler(%s/%s) failed: %s",
			req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	if err := r.updateDatahubByScaler(instance); err != nil {
		scope.Errorf("Update datahub for AlamedaScaler(%s/%s) failed: %s",
			req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	return ctrl.Result{}, nil
}

func (r *AlamedaScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autoscalingv1alpha2.AlamedaScaler{}).
		Complete(r)
}
