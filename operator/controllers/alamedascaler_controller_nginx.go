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

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"

	alamedascaler_reconciler "github.com/containers-ai/alameda/operator/pkg/reconciler/alamedascaler"
	utilsresource "github.com/containers-ai/alameda/operator/pkg/utils/resources"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	"github.com/containers-ai/alameda/pkg/utils/log"
	appsapi_v1 "github.com/openshift/api/apps/v1"
	routeapi_v1 "github.com/openshift/api/route/v1"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

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
	err := k8sClient.List(
		ctx, &alamedaScalerList, &client.ListOptions{
			Namespace: objectMeta.Namespace,
		})
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
		if ok := isLabelsSelectedBySelector(
			*alamedaScaler.Spec.Nginx.Selector, objectMeta.GetLabels()); ok {
			candidates = append(candidates, alamedaScaler)
		}
	}
	return candidates, nil
}

func init() {
	RegisterAlamedaScalerController(autoscalingv1alpha1.AlamedaScalerTypeNginx,
		listCandidatesNginxAlamedaScaler)
}

// AlamedaScalerNginxReconciler reconciles a AlamedaScaler object
type AlamedaScalerNginxReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ClusterUID            string
	Logger                *log.Scope
	ReconcileTimeout      time.Duration
	DatahubClient         *datahubpkg.Client
	HasOpenShiftAPIAppsv1 bool
}

// Reconcile reads that state of the cluster for a AlamedaScaler object and makes changes based on the state read
// and what is in the AlamedaScaler .Spec
func (r *AlamedaScalerNginxReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ReconcileTimeout)
	defer cancel()
	cachedAlamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	err := r.Client.Get(ctx, client.ObjectKey{
		Namespace: req.Namespace, Name: req.Name,
	}, &cachedAlamedaScaler)
	if err != nil && k8serrors.IsNotFound(err) {
		r.Logger.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleDeletion(ctx, req.Namespace, req.Name); err != nil {
			r.Logger.Warnf("Handle deletion of AlamedaScaler(%s/%s) failed, retry reconciling: %s",
				req.Namespace, req.Name, err)
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
	err = r.Get(context.Background(), client.ObjectKey{
		Namespace: req.Namespace, Name: alamedaScaler.Spec.Nginx.Service,
	}, &svcIns)
	if err != nil && k8serrors.IsNotFound(err) {
		scope.Errorf("service (%s/%s) does not exist",
			req.Namespace, alamedaScaler.Spec.Nginx.Service)
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		return ctrl.Result{Requeue: false}, nil
	}

	if cachedAlamedaScaler.GetDeletionTimestamp() != nil {
		r.Logger.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleDeletion(ctx, req.Namespace, req.Name); err != nil {
			r.Logger.Warnf("Handle deletion of AlamedaScaler(%s/%s) failed, retry reconciling: %s",
				req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		r.Logger.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}

	r.Logger.Infof("Reconciling AlamedaScaler(%s/%s)...", req.Namespace, req.Name)

	alamedaScaler = r.setDefaultAlamedaScaler(alamedaScaler)
	alamedaScaler.Status.AlamedaController = autoscalingv1alpha1.NewAlamedaController()
	if alamedaScaler, err = r.listAndAddDeploymentsIntoAlamedaScalerStatus(
		context.TODO(), alamedaScaler, svcIns); err != nil {
		scope.Warnf("List and add Deployments into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v",
			req.Namespace, req.Name, requeueAfter.Seconds(), err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	if hasOpenshiftAPIAppsV1 {
		if alamedaScaler, err = r.listAndAddDeploymentConfigsIntoAlamedaScalerStatus(
			context.TODO(), alamedaScaler, svcIns); err != nil {
			scope.Warnf("List and add DeploymentConfigs into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v",
				req.Namespace, req.Name, requeueAfter.Seconds(), err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
	}
	if alamedaScaler, err = r.listAndAddStatefulSetsIntoAlamedaScalerStatus(
		context.TODO(), alamedaScaler, svcIns); err != nil {
		scope.Warnf("List and add StatefulSets into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v",
			req.Namespace, req.Name, requeueAfter.Seconds(), err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	nMatchedDC := len(alamedaScaler.Status.AlamedaController.DeploymentConfigs)
	nMatchedSts := len(alamedaScaler.Status.AlamedaController.StatefulSets)
	nMatchedDeploy := len(alamedaScaler.Status.AlamedaController.Deployments)
	if (nMatchedDC + nMatchedSts + nMatchedDeploy) > 1 {
		alamedaScaler.Status.Nginx.Message = fmt.Sprintf("More than one controller mapped by service %s in AlamedaScaler(%s/%s): %v",
			svcIns.GetName(), req.Namespace, req.Name, alamedaScaler.Status.AlamedaController)
		if err := r.Update(context.TODO(), &alamedaScaler); err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s",
				req.Namespace, req.Name, err)
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

	nginxs := r.prepareNginxs(ctx, alamedaScaler)
	err = r.syncWithDatahub(ctx, alamedaScaler, nginxs)
	if err != nil {
		r.Logger.Warnf("Synchornize nginx with remote of AlamedaScaler(%s/%s) failed, retry reconciling: %s",
			req.Namespace, req.Name, err)

		alamedaScaler.Status.Nginx.Message = err.Error()
		if err := r.Update(context.TODO(), &alamedaScaler); err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s",
				req.Namespace, req.Name, err)
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

func (r AlamedaScalerNginxReconciler) isAlamedaScalerTypeNeedToBeReconciled(
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) bool {
	if alamedaScaler.GetType() != autoscalingv1alpha1.AlamedaScalerTypeNginx {
		return false
	}
	return true
}

func (r AlamedaScalerNginxReconciler) handleDeletion(
	ctx context.Context, namespace, name string) error {
	wg := errgroup.Group{}
	wg.Go(func() error {
		err := r.DatahubClient.DeleteByOpts(
			&entities.ApplicationNginx{}, datahubpkg.Option{
				Entity: entities.ApplicationNginx{
					ClusterName:            r.ClusterUID,
					AlamedaScalerNamespace: namespace,
					AlamedaScalerName:      name,
				},
				Fields: []string{
					"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
			})
		if err != nil {
			return fmt.Errorf("Delete nginxs for alamedascaler (%s/%s) failed: %s",
				namespace, name, err.Error())
		}
		return nil
	})
	return wg.Wait()
}

func (r AlamedaScalerNginxReconciler) listAndAddDeploymentsIntoAlamedaScalerStatus(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, svcIns corev1.Service) (
	autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(
		r.Client, &alamedaScaler)
	svcDeployments, err := listResources.ListDeploymentsByNamespaceLabels(
		alamedaScaler.Namespace, svcIns.Spec.Selector)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list Deployments for service failed")
	}
	deployments, err := listResources.ListDeploymentsByNamespaceLabels(
		alamedaScaler.Namespace, alamedaScaler.Spec.Nginx.Selector.MatchLabels)
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
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(
			ctx, deployment.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if Deployment can be monitored failed")
		}

		//ok = ok && IsMonitoredByAlamedaScalerController(deployment.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeNginx)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByDeployment(&deployment)
			if err != nil {
				return alamedaScaler,
					errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by Deployment failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(deployment.ObjectMeta)
			r.addAlamedaResourceIntoStatus(alamedaScaler, autoscalingv1alpha1.DeploymentController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerNginxReconciler) listAndAddStatefulSetsIntoAlamedaScalerStatus(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, svcIns corev1.Service) (
	autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	svcStatefulSets, err := listResources.ListStatefulSetsByNamespaceLabels(
		alamedaScaler.Namespace, svcIns.Spec.Selector)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list StatefulSets for service failed")
	}
	statefulSets, err := listResources.ListStatefulSetsByNamespaceLabels(
		alamedaScaler.Namespace, alamedaScaler.Spec.Nginx.Selector.MatchLabels)
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
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(
			context.TODO(), statefulSet.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if StatefulSet can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(
			statefulSet.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeNginx)
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

func (r AlamedaScalerNginxReconciler) listAndAddDeploymentConfigsIntoAlamedaScalerStatus(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, svcIns corev1.Service) (
	autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	svcDcs, err := listResources.ListDeploymentConfigsByNamespaceLabels(
		alamedaScaler.Namespace, svcIns.Spec.Selector)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list DeploymentConfigs for service failed")
	}
	dploymentConfigs, err := listResources.ListDeploymentConfigsByNamespaceLabels(
		alamedaScaler.Namespace, alamedaScaler.Spec.Nginx.Selector.MatchLabels)
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
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(
			context.TODO(), deploymentConfig.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if DeploymentConfig can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(
			deploymentConfig.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeNginx)
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

func (r AlamedaScalerNginxReconciler) isWorkloadControllerCanBeMonitoredByAlamedaScaler(
	ctx context.Context, workloadController metav1.ObjectMeta, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (
	bool, error) {
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

func (r AlamedaScalerNginxReconciler) getIneffectiveAlamedaResource(
	workloadController metav1.ObjectMeta) autoscalingv1alpha1.AlamedaResource {
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

func (r AlamedaScalerNginxReconciler) setDefaultAlamedaScaler(
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) autoscalingv1alpha1.AlamedaScaler {
	alamedaScaler.SetDefaultValue()
	alamedaScaler.Spec.Type = autoscalingv1alpha1.AlamedaScalerTypeNginx
	return alamedaScaler
}

func (r AlamedaScalerNginxReconciler) addAlamedaResourceIntoStatus(
	as autoscalingv1alpha1.AlamedaScaler, arType autoscalingv1alpha1.AlamedaControllerType,
	ar autoscalingv1alpha1.AlamedaResource) {
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

func (r AlamedaScalerNginxReconciler) syncWithDatahub(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, nginxs []entities.ApplicationNginx) error {
	r.Logger.Debugf("Synchronize with Datahub. Nginx: %+v", nginxs)
	wg := errgroup.Group{}
	wg.Go(func() error {
		if err := r.DatahubClient.Create(&nginxs); err != nil {
			return errors.Wrap(err, "creae nginxs to Datahub failed")
		}
		delNginxs := r.listNginxsToDelete(ctx, alamedaScaler, nginxs)
		if err := r.DatahubClient.Delete(&delNginxs); err != nil {
			return errors.Wrap(err, "delete nginxs failed")
		}
		return nil
	})

	return wg.Wait()
}

func (r AlamedaScalerNginxReconciler) prepareNginxs(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) []entities.ApplicationNginx {
	nginxs := []entities.ApplicationNginx{}
	routeList := &routeapi_v1.RouteList{}
	routePodList := &corev1.PodList{}
	if r.HasOpenShiftAPIAppsv1 {
		err := r.List(ctx, routeList, &client.ListOptions{Namespace: alamedaScaler.GetNamespace()})
		if err != nil {
			scope.Errorf("List routes for nginx application in alamedascaler (%s/%s) failed: %s",
				alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
		}
		err = r.List(ctx, routePodList, &client.ListOptions{
			Namespace: alamedaScaler.Spec.Nginx.ExporterNamespace,
		})
		if err != nil {
			scope.Errorf(
				"List route pods in namespace %s for nginx application in alamedascaler (%s/%s) failed: %s",
				alamedaScaler.Spec.Nginx.ExporterNamespace, alamedaScaler.GetNamespace(),
				alamedaScaler.GetName(), err.Error())
		}
	}
	routeName := ""
	for _, route := range routeList.Items {
		if route.Spec.To.Kind == "Service" &&
			route.Spec.To.Name == alamedaScaler.Spec.Nginx.Service {
			routeName = route.Name
		}
	}

	exporterPods := ""
	for _, routePod := range routePodList.Items {
		if exporterPods == "" {
			exporterPods = routePod.GetName()
		} else {
			exporterPods = fmt.Sprintf("%s,%s", exporterPods, routePod.GetName())
		}
	}

	for _, deploy := range alamedaScaler.Status.Nginx.AlamedaController.Deployments {
		nginx := entities.ApplicationNginx{
			ResourceK8sMinReplicas:      *alamedaScaler.Spec.Nginx.MinReplicas,
			ResourceK8sMaxReplicas:      *alamedaScaler.Spec.Nginx.MaxReplicas,
			Namespace:                   alamedaScaler.Spec.Nginx.ExporterNamespace,
			ExporterNamespace:           alamedaScaler.Spec.Nginx.ExporterNamespace,
			ClusterName:                 r.ClusterUID,
			AlamedaScalerName:           alamedaScaler.GetName(),
			AlamedaScalerNamespace:      alamedaScaler.GetNamespace(),
			ExporterPods:                exporterPods,
			Policy:                      alamedaScaler.Spec.Policy,
			EnableExecution:             alamedaScaler.IsEnableExecution(),
			ResourceK8sNamespace:        deploy.Namespace,
			ResourceK8sName:             deploy.Name,
			ResourceK8sKind:             "Deployment",
			ResourceK8sSpecReplicas:     *deploy.SpecReplicas,
			ResourceK8sReplicas:         int32(len(deploy.Pods)),
			ResourceK8sServiceNamespace: alamedaScaler.GetNamespace(),
			ResourceK8sServiceName:      alamedaScaler.Spec.Nginx.Service,
			ResourceK8sRouteNamespace:   alamedaScaler.GetNamespace(),
			ResourceK8sRouteName:        routeName,
			ReplicaMarginPercentage:     *alamedaScaler.Spec.Nginx.ReplicaMarginPercentage,
		}
		if alamedaScaler.Spec.Nginx.TargetResponseTime != nil {
			nginx.HttpResponseTime = *alamedaScaler.Spec.Nginx.TargetResponseTime
		}
		nginxs = append(nginxs, nginx)
	}
	for _, dc := range alamedaScaler.Status.Nginx.AlamedaController.DeploymentConfigs {
		nginx := entities.ApplicationNginx{
			ResourceK8sMinReplicas:      *alamedaScaler.Spec.Nginx.MinReplicas,
			ResourceK8sMaxReplicas:      *alamedaScaler.Spec.Nginx.MaxReplicas,
			Namespace:                   alamedaScaler.Spec.Nginx.ExporterNamespace,
			ExporterNamespace:           alamedaScaler.Spec.Nginx.ExporterNamespace,
			ClusterName:                 r.ClusterUID,
			AlamedaScalerName:           alamedaScaler.GetName(),
			AlamedaScalerNamespace:      alamedaScaler.GetNamespace(),
			ExporterPods:                exporterPods,
			Policy:                      alamedaScaler.Spec.Policy,
			EnableExecution:             alamedaScaler.IsEnableExecution(),
			ResourceK8sNamespace:        dc.Namespace,
			ResourceK8sName:             dc.Name,
			ResourceK8sKind:             "DeploymentConfig",
			ResourceK8sSpecReplicas:     *dc.SpecReplicas,
			ResourceK8sReplicas:         int32(len(dc.Pods)),
			ResourceK8sServiceNamespace: alamedaScaler.GetNamespace(),
			ResourceK8sServiceName:      alamedaScaler.Spec.Nginx.Service,
			ResourceK8sRouteNamespace:   alamedaScaler.GetNamespace(),
			ResourceK8sRouteName:        routeName,
			ReplicaMarginPercentage:     *alamedaScaler.Spec.Nginx.ReplicaMarginPercentage,
		}
		if alamedaScaler.Spec.Nginx.TargetResponseTime != nil {
			nginx.HttpResponseTime = *alamedaScaler.Spec.Nginx.TargetResponseTime
		}
		nginxs = append(nginxs, nginx)
	}
	for _, sts := range alamedaScaler.Status.Nginx.AlamedaController.StatefulSets {
		nginx := entities.ApplicationNginx{
			ResourceK8sMinReplicas:      *alamedaScaler.Spec.Nginx.MinReplicas,
			ResourceK8sMaxReplicas:      *alamedaScaler.Spec.Nginx.MaxReplicas,
			Namespace:                   alamedaScaler.Spec.Nginx.ExporterNamespace,
			ExporterNamespace:           alamedaScaler.Spec.Nginx.ExporterNamespace,
			ClusterName:                 r.ClusterUID,
			AlamedaScalerName:           alamedaScaler.GetName(),
			AlamedaScalerNamespace:      alamedaScaler.GetNamespace(),
			ExporterPods:                exporterPods,
			Policy:                      alamedaScaler.Spec.Policy,
			EnableExecution:             alamedaScaler.IsEnableExecution(),
			ResourceK8sNamespace:        sts.Namespace,
			ResourceK8sName:             sts.Name,
			ResourceK8sKind:             "StatefulSet",
			ResourceK8sSpecReplicas:     *sts.SpecReplicas,
			ResourceK8sReplicas:         int32(len(sts.Pods)),
			ResourceK8sServiceNamespace: alamedaScaler.GetNamespace(),
			ResourceK8sServiceName:      alamedaScaler.Spec.Nginx.Service,
			ResourceK8sRouteNamespace:   alamedaScaler.GetNamespace(),
			ResourceK8sRouteName:        routeName,
			ReplicaMarginPercentage:     *alamedaScaler.Spec.Nginx.ReplicaMarginPercentage,
		}
		if alamedaScaler.Spec.Nginx.TargetResponseTime != nil {
			nginx.HttpResponseTime = *alamedaScaler.Spec.Nginx.TargetResponseTime
		}
		nginxs = append(nginxs, nginx)
	}
	return nginxs
}

func (r AlamedaScalerNginxReconciler) listNginxsToDelete(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	createNginxs []entities.ApplicationNginx) []entities.ApplicationNginx {
	delNginxs := []entities.ApplicationNginx{}
	nginxs := []entities.ApplicationNginx{}
	err := r.DatahubClient.List(&nginxs, datahubpkg.Option{
		Entity: entities.ApplicationNginx{
			ClusterName:            r.ClusterUID,
			Namespace:              alamedaScaler.Spec.Nginx.ExporterNamespace,
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			AlamedaScalerName:      alamedaScaler.GetName(),
		},
		Fields: []string{
			"ClusterName", "Namespace", "AlamedaScalerNamespace", "AlamedaScalerName"},
	})
	if err != nil {
		scope.Errorf("list nginxs to delete for alamedascaler (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}
	for _, nginx := range nginxs {
		toDel := true
		for _, createNginx := range createNginxs {
			if createNginx.ClusterName == nginx.ClusterName &&
				createNginx.ResourceK8sNamespace == nginx.ResourceK8sNamespace &&
				createNginx.ResourceK8sName == nginx.ResourceK8sName &&
				createNginx.ResourceK8sKind == nginx.ResourceK8sKind &&
				createNginx.Namespace == nginx.Namespace {
				toDel = false
			}
		}
		if toDel {
			delNginxs = append(delNginxs, nginx)
		}
	}
	return delNginxs
}
