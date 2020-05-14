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
	"fmt"
	"time"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	"github.com/containers-ai/alameda/operator/pkg/nginx"
	nginxmodel "github.com/containers-ai/alameda/operator/pkg/nginx"
	alamedascaler_reconciler "github.com/containers-ai/alameda/operator/pkg/reconciler/alamedascaler"
	utilsresource "github.com/containers-ai/alameda/operator/pkg/utils/resources"
	appsapi_v1 "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	nginxrepository "github.com/containers-ai/alameda/operator/datahub/client/nginx"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahubschemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var listCandidatesNginxAlamedaScaler = func(
	ctx context.Context,
	k8sClient client.Client,
	objectMeta metav1.ObjectMeta,
) ([]autoscalingv1alpha1.AlamedaScaler, error) {

	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := k8sClient.List(ctx, &alamedaScalerList, &client.ListOptions{Namespace: objectMeta.Namespace})
	if err != nil {
		return nil, errors.Wrap(err, "list AlamedaScalers failed")
	}

	candidates := make([]autoscalingv1alpha1.AlamedaScaler, 0)
	for _, alamedaScaler := range alamedaScalerList.Items {
		if alamedaScaler.GetType() != autoscalingv1alpha1.AlamedaScalerTypeNginx {
			continue
		}
		if alamedaScaler.Spec.Nginx == nil {
			continue
		}
		if ok := isLabelsSelectedBySelector(*alamedaScaler.Spec.Nginx.Selector, objectMeta.GetLabels()); ok {
			candidates = append(candidates, alamedaScaler)
		}
	}
	return candidates, nil
}

func init() {
	RegisterAlamedaScalerController(autoscalingv1alpha1.AlamedaScalerTypeNginx, listCandidatesNginxAlamedaScaler)
}

// AlamedaScalerNginxReconciler reconciles a AlamedaScaler object
type AlamedaScalerNginxReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ClusterUID                         string
	NginxRepository                    nginxrepository.NginxRepository
	DatahubApplicationNginxSchema      datahubschemas.Schema
	DatahubApplicationNginxMeasurement datahubschemas.Measurement
	Logger                             *log.Scope
	ReconcileTimeout                   time.Duration
}

// Reconcile reads that state of the cluster for a AlamedaScaler object and makes changes based on the state read
// and what is in the AlamedaScaler .Spec
func (r *AlamedaScalerNginxReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ReconcileTimeout)
	defer cancel()
	cachedAlamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	err := r.Client.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: req.Name}, &cachedAlamedaScaler)
	if err != nil && k8serrors.IsNotFound(err) {
		r.Logger.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleDeletion(ctx, req.Namespace, req.Name); err != nil {
			r.Logger.Warnf("Handle deletion of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		r.Logger.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}

	alamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	cachedAlamedaScaler.DeepCopyInto(&alamedaScaler)
	if alamedaScaler.Status.Nginx == nil {
		alamedaScaler.Status.Nginx = &autoscalingv1alpha1.NginxStatus{}
	}

	if !r.isAlamedaScalerTypeNeedToBeReconciled(alamedaScaler) {
		return ctrl.Result{Requeue: false}, nil
	}

	if alamedaScaler.Spec.Nginx == nil {
		scope.Errorf("nginx spec of alamedascaler (%s/%s) does not configure", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}

	svcIns := corev1.Service{}
	err = r.Get(context.Background(), client.ObjectKey{Namespace: req.Namespace, Name: alamedaScaler.Spec.Nginx.Service}, &svcIns)
	if err != nil && k8serrors.IsNotFound(err) {
		scope.Errorf("service (%s/%s) does not exist", req.Namespace, alamedaScaler.Spec.Nginx.Service)
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		return ctrl.Result{Requeue: false}, nil
	}

	if cachedAlamedaScaler.GetDeletionTimestamp() != nil {
		r.Logger.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleDeletion(ctx, req.Namespace, req.Name); err != nil {
			r.Logger.Warnf("Handle deletion of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		r.Logger.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}

	r.Logger.Infof("Reconciling AlamedaScaler(%s/%s)...", req.Namespace, req.Name)

	alamedaScaler = r.setDefaultAlamedaScaler(alamedaScaler)
	alamedaScaler.Status.AlamedaController = autoscalingv1alpha1.NewAlamedaController()
	if alamedaScaler, err = r.listAndAddDeploymentsIntoAlamedaScalerStatus(context.TODO(), alamedaScaler, svcIns); err != nil {
		scope.Warnf("List and add Deployments into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v", req.Namespace, req.Name, requeueAfter.Seconds(), err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	if hasOpenshiftAPIAppsV1 {
		if alamedaScaler, err = r.listAndAddDeploymentConfigsIntoAlamedaScalerStatus(context.TODO(), alamedaScaler, svcIns); err != nil {
			scope.Warnf("List and add DeploymentConfigs into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v", req.Namespace, req.Name, requeueAfter.Seconds(), err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
	}
	if alamedaScaler, err = r.listAndAddStatefulSetsIntoAlamedaScalerStatus(context.TODO(), alamedaScaler, svcIns); err != nil {
		scope.Warnf("List and add StatefulSets into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v", req.Namespace, req.Name, requeueAfter.Seconds(), err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	nMatchedDC := len(alamedaScaler.Status.AlamedaController.DeploymentConfigs)
	nMatchedSts := len(alamedaScaler.Status.AlamedaController.StatefulSets)
	nMatchedDeploy := len(alamedaScaler.Status.AlamedaController.Deployments)
	if (nMatchedDC + nMatchedSts + nMatchedDeploy) > 1 {
		alamedaScaler.Status.Nginx.Message = fmt.Sprintf("More than one controller mapped by service %s in AlamedaScaler(%s/%s): %v", svcIns.GetName(), req.Namespace, req.Name, alamedaScaler.Status.AlamedaController)
		if err := r.Update(context.TODO(), &alamedaScaler); err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	nginxController := &autoscalingv1alpha1.AlamedaController{}
	alamedaScaler.Status.AlamedaController.DeepCopyInto(nginxController)
	alamedaScaler.Status.Nginx.AlamedaController = *nginxController
	alamedaScaler.Status.AlamedaController = autoscalingv1alpha1.AlamedaController{}
	if err := r.Update(context.TODO(), &alamedaScaler); err != nil {
		scope.Errorf("Update AlamedaScaler(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	nginxs := r.prepareNginxs(alamedaScaler)
	err = r.syncWithDatahub(ctx, alamedaScaler, nginxs)
	if err != nil {
		r.Logger.Warnf("Synchornize nginx with remote of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)

		alamedaScaler.Status.Nginx.Message = err.Error()
		if err := r.Update(context.TODO(), &alamedaScaler); err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	scope.Infof("Reconcile AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
	return ctrl.Result{}, nil
}

func (r *AlamedaScalerNginxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autoscalingv1alpha1.AlamedaScaler{}).
		Complete(r)
}

func (r AlamedaScalerNginxReconciler) isAlamedaScalerTypeNeedToBeReconciled(alamedaScaler autoscalingv1alpha1.AlamedaScaler) bool {
	if alamedaScaler.GetType() != autoscalingv1alpha1.AlamedaScalerTypeNginx {
		return false
	}
	return true
}

func (r AlamedaScalerNginxReconciler) handleDeletion(ctx context.Context, namespace, name string) error {
	wg := errgroup.Group{}
	wg.Go(func() error {
		opt := nginxrepository.DeleteNginxsOption{
			ClusterName:            r.ClusterUID,
			AlamedaScalerNamespace: namespace,
			AlamedaScalerName:      name,
		}
		if err := r.NginxRepository.DeleteNginxsByOption(ctx, opt); err != nil {
			return errors.Wrap(err, "delete nginxs from Datahub failed")
		}
		return nil
	})
	return wg.Wait()
}

func (r AlamedaScalerNginxReconciler) listAndAddDeploymentsIntoAlamedaScalerStatus(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, svcIns corev1.Service) (autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	svcDeployments, err := listResources.ListDeploymentsByNamespaceLabels(alamedaScaler.Namespace, svcIns.Spec.Selector)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list Deployments for service failed")
	}
	deployments, err := listResources.ListDeploymentsByNamespaceLabels(alamedaScaler.Namespace, alamedaScaler.Spec.Nginx.Selector.MatchLabels)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list Deployments failed")
	}

	matchedDeploys := []appsv1.Deployment{}
	for _, deploy := range deployments {
		for _, svcDeploy := range svcDeployments {
			if deploy.GetName() == svcDeploy.GetName() {
				matchedDeploys = append(matchedDeploys, deploy)
			}
		}
	}

	for _, deployment := range matchedDeploys {
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(ctx, deployment.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if Deployment can be monitored failed")
		}

		//ok = ok && IsMonitoredByAlamedaScalerController(deployment.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeNginx)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByDeployment(&deployment)
			if err != nil {
				return alamedaScaler, errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by Deployment failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(deployment.ObjectMeta)
			r.addAlamedaResourceIntoStatus(alamedaScaler, autoscalingv1alpha1.DeploymentController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerNginxReconciler) listAndAddStatefulSetsIntoAlamedaScalerStatus(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, svcIns corev1.Service) (autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	svcStatefulSets, err := listResources.ListStatefulSetsByNamespaceLabels(alamedaScaler.Namespace, svcIns.Spec.Selector)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list StatefulSets for service failed")
	}
	statefulSets, err := listResources.ListStatefulSetsByNamespaceLabels(alamedaScaler.Namespace, alamedaScaler.Spec.Nginx.Selector.MatchLabels)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list StatefulSets failed")
	}

	matchedStss := []appsv1.StatefulSet{}
	for _, sts := range statefulSets {
		for _, svcSts := range svcStatefulSets {
			if sts.GetName() == svcSts.GetName() {
				matchedStss = append(matchedStss, sts)
			}
		}
	}

	for _, statefulSet := range matchedStss {
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(context.TODO(), statefulSet.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if StatefulSet can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(statefulSet.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeNginx)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByStatefulSet(&statefulSet)
			if err != nil {
				return alamedaScaler, errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by StatefulSet failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(statefulSet.ObjectMeta)
			r.addAlamedaResourceIntoStatus(alamedaScaler, autoscalingv1alpha1.StatefulSetController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerNginxReconciler) listAndAddDeploymentConfigsIntoAlamedaScalerStatus(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, svcIns corev1.Service) (autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	svcDcs, err := listResources.ListDeploymentConfigsByNamespaceLabels(alamedaScaler.Namespace, svcIns.Spec.Selector)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list DeploymentConfigs for service failed")
	}
	dploymentConfigs, err := listResources.ListDeploymentConfigsByNamespaceLabels(alamedaScaler.Namespace, alamedaScaler.Spec.Nginx.Selector.MatchLabels)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list DeploymentConfigs failed")
	}

	matchedDcs := []appsapi_v1.DeploymentConfig{}
	for _, dc := range dploymentConfigs {
		for _, svcDc := range svcDcs {
			if dc.GetName() == svcDc.GetName() {
				matchedDcs = append(matchedDcs, dc)
			}
		}
	}

	for _, deploymentConfig := range matchedDcs {
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(context.TODO(), deploymentConfig.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if DeploymentConfig can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(deploymentConfig.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeNginx)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByDeploymentConfig(&deploymentConfig)
			if err != nil {
				return alamedaScaler, errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by DeploymentConfig failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(deploymentConfig.ObjectMeta)
			r.addAlamedaResourceIntoStatus(alamedaScaler, autoscalingv1alpha1.DeploymentConfigController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerNginxReconciler) isWorkloadControllerCanBeMonitoredByAlamedaScaler(ctx context.Context, workloadController metav1.ObjectMeta, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (bool, error) {
	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := r.List(ctx, &alamedaScalerList, &client.ListOptions{Namespace: workloadController.GetNamespace()})
	if err != nil {
		return false, errors.Wrap(err, "list AlamedaScalers failed")
	}

	alamedaScalersObjectMeta := make([]metav1.ObjectMeta, 0, len(alamedaScalerList.Items))
	for _, alamedaScaler := range alamedaScalerList.Items {
		if alamedaScaler.Spec.Nginx.Selector == nil {
			continue
		}
		if ok := isLabelsSelectedBySelector(*alamedaScaler.Spec.Nginx.Selector, workloadController.GetLabels()); !ok {
			continue
		}
		alamedaScalersObjectMeta = append(alamedaScalersObjectMeta, alamedaScaler.ObjectMeta)
	}

	oldestAlamedaScaler := getFirstCreatedObjectMeta(alamedaScalersObjectMeta)
	if oldestAlamedaScaler.GetNamespace() != alamedaScaler.GetNamespace() ||
		oldestAlamedaScaler.GetName() != alamedaScaler.GetName() {
		return false, nil
	}
	return true, nil
}

func (r AlamedaScalerNginxReconciler) getIneffectiveAlamedaResource(workloadController metav1.ObjectMeta) autoscalingv1alpha1.AlamedaResource {
	empty := int32(0)
	return autoscalingv1alpha1.AlamedaResource{
		Namespace:    workloadController.GetNamespace(),
		Name:         workloadController.GetName(),
		UID:          string(workloadController.GetUID()),
		SpecReplicas: &empty,
		Effective:    false,
		Message:      "Is monitoring by other AlamedaScaler.",
	}
}

func (r AlamedaScalerNginxReconciler) setDefaultAlamedaScaler(alamedaScaler autoscalingv1alpha1.AlamedaScaler) autoscalingv1alpha1.AlamedaScaler {
	alamedaScaler.SetDefaultValue()
	alamedaScaler.Spec.Type = autoscalingv1alpha1.AlamedaScalerTypeNginx
	return alamedaScaler
}

func (r AlamedaScalerNginxReconciler) addAlamedaResourceIntoStatus(as autoscalingv1alpha1.AlamedaScaler, arType autoscalingv1alpha1.AlamedaControllerType, ar autoscalingv1alpha1.AlamedaResource) {
	if (&as.Status.AlamedaController) == nil {
		as.Status.AlamedaController = autoscalingv1alpha1.AlamedaController{}
	}
	ac := as.Status.AlamedaController
	switch arType {
	case autoscalingv1alpha1.DeploymentController:
		if ac.Deployments == nil {
			ac.Deployments = map[string]autoscalingv1alpha1.AlamedaResource{}
		}
		ac.Deployments[ar.GetNamespacedName()] = ar
	case autoscalingv1alpha1.DeploymentConfigController:
		if ac.DeploymentConfigs == nil {
			ac.DeploymentConfigs = map[string]autoscalingv1alpha1.AlamedaResource{}
		}
		ac.DeploymentConfigs[ar.GetNamespacedName()] = ar
	case autoscalingv1alpha1.StatefulSetController:
		if ac.StatefulSets == nil {
			ac.StatefulSets = map[string]autoscalingv1alpha1.AlamedaResource{}
		}
		ac.StatefulSets[ar.GetNamespacedName()] = ar
	}
}

func (r AlamedaScalerNginxReconciler) syncWithDatahub(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, nginxs []nginxmodel.Nginx) error {
	r.Logger.Debugf("Synchronize with Datahub. Nginx: %+v", nginxs)
	wg := errgroup.Group{}
	wg.Go(func() error {
		if err := r.NginxRepository.CreateNginxs(ctx, nginxs); err != nil {
			return errors.Wrap(err, "creae nginxs to Datahub failed")
		}
		delNginxs := r.listNginxsToDelete(ctx, alamedaScaler, nginxs)
		if err := r.NginxRepository.DeleteNginxs(ctx, delNginxs); err != nil {
			return errors.Wrap(err, "delete nginxs failed")
		}
		return nil
	})

	return wg.Wait()
}

func (r AlamedaScalerNginxReconciler) prepareNginxs(alamedaScaler autoscalingv1alpha1.AlamedaScaler) []nginxmodel.Nginx {
	nginxs := []nginxmodel.Nginx{}
	for _, deploy := range alamedaScaler.Status.Nginx.AlamedaController.Deployments {
		nginxs = append(nginxs, nginxmodel.Nginx{
			MinReplicas:            *alamedaScaler.Spec.Nginx.MinReplicas,
			MaxReplicas:            *alamedaScaler.Spec.Nginx.MaxReplicas,
			ExporterNamespace:      alamedaScaler.Spec.Nginx.ExporterNamespace,
			ClusterName:            r.ClusterUID,
			AlamedaScalerName:      alamedaScaler.GetName(),
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			Policy:                 alamedaScaler.Spec.Policy,
			EnableExecution:        alamedaScaler.IsEnableExecution(),
			ResourceMeta: nginx.ResourceMeta{
				KubernetesMeta: nginx.KubernetesMeta{
					Namespace:        deploy.Namespace,
					Name:             deploy.Name,
					Kind:             "Deployment",
					SpecReplicas:     *deploy.SpecReplicas,
					ServiceName:      alamedaScaler.Spec.Nginx.Service,
					ServiceNamespace: alamedaScaler.GetNamespace(),
					ReadyReplicas:    int32(len(deploy.Pods)),
				},
			},
			ReplicaMarginPercentage: *alamedaScaler.Spec.Nginx.ReplicaMarginPercentage,
		})
	}
	for _, dc := range alamedaScaler.Status.Nginx.AlamedaController.DeploymentConfigs {
		nginxs = append(nginxs, nginxmodel.Nginx{
			MinReplicas:            *alamedaScaler.Spec.Nginx.MinReplicas,
			MaxReplicas:            *alamedaScaler.Spec.Nginx.MaxReplicas,
			ExporterNamespace:      alamedaScaler.Spec.Nginx.ExporterNamespace,
			ClusterName:            r.ClusterUID,
			AlamedaScalerName:      alamedaScaler.GetName(),
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			Policy:                 alamedaScaler.Spec.Policy,
			EnableExecution:        alamedaScaler.IsEnableExecution(),
			ResourceMeta: nginx.ResourceMeta{
				KubernetesMeta: nginx.KubernetesMeta{
					Namespace:        dc.Namespace,
					Name:             dc.Name,
					Kind:             "DeploymentConfig",
					ServiceName:      alamedaScaler.Spec.Nginx.Service,
					ServiceNamespace: alamedaScaler.GetNamespace(),
					SpecReplicas:     *dc.SpecReplicas,
					ReadyReplicas:    int32(len(dc.Pods)),
				},
			},
			ReplicaMarginPercentage: *alamedaScaler.Spec.Nginx.ReplicaMarginPercentage,
		})
	}
	for _, sts := range alamedaScaler.Status.Nginx.AlamedaController.StatefulSets {
		nginxs = append(nginxs, nginxmodel.Nginx{
			MinReplicas:            *alamedaScaler.Spec.Nginx.MinReplicas,
			MaxReplicas:            *alamedaScaler.Spec.Nginx.MaxReplicas,
			ExporterNamespace:      alamedaScaler.Spec.Nginx.ExporterNamespace,
			ClusterName:            r.ClusterUID,
			AlamedaScalerName:      alamedaScaler.GetName(),
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			Policy:                 alamedaScaler.Spec.Policy,
			EnableExecution:        alamedaScaler.IsEnableExecution(),
			ResourceMeta: nginx.ResourceMeta{
				KubernetesMeta: nginx.KubernetesMeta{
					Namespace:        sts.Namespace,
					Name:             sts.Name,
					Kind:             "StatefulSet",
					ServiceName:      alamedaScaler.Spec.Nginx.Service,
					ServiceNamespace: alamedaScaler.GetNamespace(),
					SpecReplicas:     *sts.SpecReplicas,
					ReadyReplicas:    int32(len(sts.Pods)),
				},
			},
			ReplicaMarginPercentage: *alamedaScaler.Spec.Nginx.ReplicaMarginPercentage,
		})
	}
	return nginxs
}

func (r AlamedaScalerNginxReconciler) listNginxsToDelete(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, createNginxs []nginxmodel.Nginx) []nginxmodel.Nginx {
	delNginxs := []nginxmodel.Nginx{}
	nginxs, err := r.NginxRepository.ListNginxs(ctx, nginxrepository.ListNginxsOption{
		ClusterName:            r.ClusterUID,
		ExporterNamespace:      alamedaScaler.Spec.Nginx.ExporterNamespace,
		AlamedaScalerName:      alamedaScaler.GetName(),
		AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
	})
	if err != nil {
		scope.Errorf("list nginxs to delete for alamedascaler (%s/%s) failed: %s", alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}
	for _, nginx := range nginxs {
		toDel := true
		for _, createNginx := range createNginxs {
			if createNginx.ClusterName == nginx.ClusterName &&
				createNginx.KubernetesMeta.Namespace == nginx.KubernetesMeta.Namespace &&
				createNginx.KubernetesMeta.Name == nginx.KubernetesMeta.Name &&
				createNginx.KubernetesMeta.Kind == nginx.KubernetesMeta.Kind &&
				createNginx.ExporterNamespace == nginx.ExporterNamespace {
				toDel = false
			}
		}
		if toDel {
			delNginxs = append(delNginxs, nginx)
		}
	}
	return delNginxs
}
