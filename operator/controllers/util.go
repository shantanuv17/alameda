package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	machinegrouprepository "github.com/containers-ai/alameda/operator/datahub/client/machinegroup"
	machinesetrepository "github.com/containers-ai/alameda/operator/datahub/client/machineset"
	"github.com/containers-ai/alameda/operator/pkg/machineset"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	alamedaScalerTypeAnnotationKey = "alamedascalers.autoscaling.containers.ai/type"
)

// setAlamedaScalerControllerName sets the AlamedaScaler controller's type into the object's annotation.
func setAlamedaScalerControllerName(obj metav1.Object, name string) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[alamedaScalerTypeAnnotationKey] = name

	obj.SetAnnotations(annotations)
}

// IsMonitoredByAlamedaScalerController returns if the object is monitored by the provided name of AlamdeScaler controller.
func IsMonitoredByAlamedaScalerController(
	obj metav1.ObjectMeta, name string) bool {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return false
	}
	if annotations[alamedaScalerTypeAnnotationKey] == name {
		return true
	}
	return false
}

func isLabelsSelectedBySelector(
	selector metav1.LabelSelector, labels map[string]string) bool {
	for k, v := range selector.MatchLabels {
		v2, exist := labels[k]
		if v != v2 || !exist {
			return false
		}
	}

	return true
}

func getFirstCreatedObjectMeta(objs []metav1.ObjectMeta) metav1.ObjectMeta {
	t := time.Now()
	firstCreatedObe := metav1.ObjectMeta{}
	for _, obj := range objs {
		if obj.GetCreationTimestamp().UnixNano() < t.UnixNano() {
			firstCreatedObe = obj
			t = obj.GetCreationTimestamp().Time
		}
	}
	return firstCreatedObe
}

func getTotalResourceFromContainers(
	containers []corev1.Container) corev1.ResourceRequirements {
	total := corev1.ResourceRequirements{
		Limits:   corev1.ResourceList{},
		Requests: corev1.ResourceList{},
	}
	for _, c := range containers {
		for resourceName, quantity := range c.Resources.Limits {
			q := total.Limits[resourceName]
			q.Add(quantity)
			total.Limits[resourceName] = q
		}
		for resourceName, quantity := range c.Resources.Requests {
			q := total.Requests[resourceName]
			q.Add(quantity)
			total.Requests[resourceName] = q
		}
	}

	return total
}

func getFirstTime(times []time.Time) time.Time {
	min := time.Now()
	for _, t := range times {
		if min.After(t) {
			min = t
		}
	}

	return min
}

func SyncCAInfoWithScalerAndMachineGroup(ctx context.Context,
	clusterUID string, clnt client.Client,
	datahubMachineSetRepo *machinesetrepository.MachineSetRepository,
	datahubMachineGroupRepo *machinegrouprepository.MachineGroupRepository,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	mgIns autoscalingv1alpha1.AlamedaMachineGroupScaler) error {

	mgsInScaler := []entities.ResourceClusterAutoscalerMachinegroup{}
	err := datahubMachineGroupRepo.GetDatahubClient().List(
		&mgsInScaler, datahubpkg.Option{
			Entity: entities.ResourceClusterAutoscalerMachinegroup{
				ClusterName:            clusterUID,
				AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
				AlamedaScalerName:      alamedaScaler.GetName(),
			},
			Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
		})
	if err != nil {
		return fmt.Errorf("List machinegroups with alamedascaler (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}

	for _, mgInScaler := range mgsInScaler {
		err = datahubMachineSetRepo.DeleteMachineSetsByOption(ctx,
			machinesetrepository.DeleteMachineSetsOption{
				ClusterName:           clusterUID,
				MachineGroupNamespace: mgInScaler.Namespace,
				MachineGroupName:      mgInScaler.Name,
			})
		if err != nil {
			return fmt.Errorf("Delete machinesets with machinegroup (%s/%s) failed: %s",
				mgInScaler.Namespace, mgInScaler.Name, err.Error())
		}
	}

	err = datahubMachineGroupRepo.DeleteMachineGroupsByOption(ctx,
		machinegrouprepository.DeleteMachineGroupsOption{
			ClusterName:            clusterUID,
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			AlamedaScalerName:      alamedaScaler.GetName(),
		})

	if err != nil {
		return fmt.Errorf("Delete machinegroups with alamedascaler (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}

	newMg := entities.ResourceClusterAutoscalerMachinegroup{
		ClusterName:            clusterUID,
		AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
		AlamedaScalerName:      alamedaScaler.GetName(),
		Name:                   mgIns.Name,
		Namespace:              mgIns.Namespace,
	}
	if mgIns.Spec.Metrics["cpu"].UtilizationTarget != nil {
		newMg.CPUMetricUtilizationTarget = *mgIns.Spec.Metrics["cpu"].UtilizationTarget
	}
	if mgIns.Spec.Metrics["cpu"].ScaleUpGap != nil {
		newMg.CPUMetricScaleUpGap = *mgIns.Spec.Metrics["cpu"].ScaleUpGap
	}
	if mgIns.Spec.Metrics["cpu"].ScaleDownGap != nil {
		newMg.CPUMetricScaleDownGap = *mgIns.Spec.Metrics["cpu"].ScaleDownGap
	}
	if mgIns.Spec.Metrics["cpu"].DurationUpThresholdPercentage != nil {
		newMg.CPUDurationUpThresholdPercentage = *mgIns.Spec.Metrics["cpu"].DurationUpThresholdPercentage
	}
	if mgIns.Spec.Metrics["cpu"].DurationDownThresholdPercentage != nil {
		newMg.CPUDurationDownThresholdPercentage = *mgIns.Spec.Metrics["cpu"].DurationDownThresholdPercentage
	}
	if mgIns.Spec.Metrics["memory"].UtilizationTarget != nil {
		newMg.MemoryMetricUtilizationTarget = *mgIns.Spec.Metrics["memory"].UtilizationTarget
	}
	if mgIns.Spec.Metrics["memory"].ScaleUpGap != nil {
		newMg.MemoryMetricScaleUpGap = *mgIns.Spec.Metrics["memory"].ScaleUpGap
	}
	if mgIns.Spec.Metrics["memory"].ScaleDownGap != nil {
		newMg.MemoryMetricScaleDownGap = *mgIns.Spec.Metrics["memory"].ScaleDownGap
	}
	if mgIns.Spec.Metrics["memory"].DurationUpThresholdPercentage != nil {
		newMg.MemoryDurationUpThresholdPercentage = *mgIns.Spec.Metrics["memory"].DurationUpThresholdPercentage
	}
	if mgIns.Spec.Metrics["memory"].DurationDownThresholdPercentage != nil {
		newMg.MemoryDurationDownThresholdPercentage = *mgIns.Spec.Metrics["memory"].DurationDownThresholdPercentage
	}
	mgs := []entities.ResourceClusterAutoscalerMachinegroup{
		newMg,
	}

	err = datahubMachineGroupRepo.CreateMachineGroups(mgs)
	if err != nil {
		return fmt.Errorf("Create machinegroup (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}

	err = datahubMachineSetRepo.DeleteMachineSetsByOption(ctx,
		machinesetrepository.DeleteMachineSetsOption{
			ClusterName:           clusterUID,
			MachineGroupNamespace: mgIns.Namespace,
			MachineGroupName:      mgIns.Name,
		})
	if err != nil {
		return fmt.Errorf("Delete machinesets with machinegroup (%s/%s) failed: %s",
			mgIns.Namespace, mgIns.Name, err.Error())
	}

	msList := mahcinev1beta1.MachineSetList{}
	err = clnt.List(ctx, &msList, &client.ListOptions{})
	if err != nil {
		return fmt.Errorf("Get MachineSet Liist in namespace %s failed: %s",
			alamedaScaler.GetNamespace(), err.Error())
	}

	mss := []machineset.MachineSet{}
	for _, ms := range msList.Items {
		for _, msSpec := range mgIns.Spec.MachineSets {
			if ms.GetName() == msSpec.Name && ms.GetNamespace() == msSpec.Namespace {
				mss = append(mss, machineset.MachineSet{
					ClusterName:           clusterUID,
					MachineGroupName:      mgIns.GetName(),
					MachineGroupNamespace: mgIns.GetNamespace(),
					ResourceMeta: machineset.ResourceMeta{
						KubernetesMeta: machineset.KubernetesMeta{
							Namespace:     ms.Namespace,
							Name:          ms.Name,
							ReadyReplicas: ms.Status.ReadyReplicas,
							SpecReplicas:  *ms.Spec.Replicas,
						},
					},
					EnableExecution: alamedaScaler.IsEnableExecution(),
					MaxReplicas:     *msSpec.MaxReplicas,
					MinReplicas:     *msSpec.MinReplicas,
				})
				break
			}
		}
	}
	err = datahubMachineSetRepo.CreateMachineSets(ctx, mss)
	if err != nil {
		return fmt.Errorf("Create machineset (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}
	return nil
}
