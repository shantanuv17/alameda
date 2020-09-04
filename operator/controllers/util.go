package controllers

import (
	"context"
	"fmt"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	operatorutils "github.com/containers-ai/alameda/operator/pkg/utils"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetTotalResourceFromContainers(containers []corev1.Container) corev1.ResourceRequirements {
	return operatorutils.GetTotalResourceFromContainers(containers)
}

func SyncCAInfoWithScalerAndMachineGroup(ctx context.Context,
	clusterUID string, clnt client.Client, datahubClient *datahubpkg.Client,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	mgIns autoscalingv1alpha1.AlamedaMachineGroupScaler) error {

	mgsInScaler := []entities.ResourceClusterAutoscalerMachinegroup{}
	err := datahubClient.List(
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
		err = datahubClient.DeleteByOpts(&entities.ResourceClusterAutoscalerMachineset{}, datahubpkg.Option{
			Entity: entities.ResourceClusterAutoscalerMachineset{
				ClusterName:           clusterUID,
				MachinegroupNamespace: mgInScaler.Namespace,
				MachinegroupName:      mgInScaler.Name,
			},
			Fields: []string{"ClusterName", "MachinegroupNamespace", "MachinegroupName"},
		})
		if err != nil {
			return fmt.Errorf("Delete machinesets with machinegroup (%s/%s) failed: %s",
				mgInScaler.Namespace, mgInScaler.Name, err.Error())
		}
	}

	err = datahubClient.DeleteByOpts(&entities.ResourceClusterAutoscalerMachinegroup{}, datahubpkg.Option{
		Entity: entities.ResourceClusterAutoscalerMachinegroup{
			ClusterName:            clusterUID,
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			AlamedaScalerName:      alamedaScaler.GetName(),
		},
		Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
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

	err = datahubClient.Create(&mgs)
	if err != nil {
		return fmt.Errorf("Create machinegroup (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}

	err = datahubClient.DeleteByOpts(&entities.ResourceClusterAutoscalerMachineset{}, datahubpkg.Option{
		Entity: entities.ResourceClusterAutoscalerMachineset{
			ClusterName:           clusterUID,
			MachinegroupNamespace: mgIns.Namespace,
			MachinegroupName:      mgIns.Name,
		},
		Fields: []string{"ClusterName", "MachinegroupNamespace", "MachinegroupName"},
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

	mss := []entities.ResourceClusterAutoscalerMachineset{}
	for _, ms := range msList.Items {
		for _, msSpec := range mgIns.Spec.MachineSets {
			if ms.GetName() == msSpec.Name && ms.GetNamespace() == msSpec.Namespace {
				mss = append(mss, entities.ResourceClusterAutoscalerMachineset{
					ClusterName:             clusterUID,
					Namespace:               ms.Namespace,
					Name:                    ms.Name,
					MachinegroupName:        mgIns.GetName(),
					MachinegroupNamespace:   mgIns.GetNamespace(),
					ResourceK8sReplicas:     ms.Status.ReadyReplicas,
					ResourceK8sSpecReplicas: *ms.Spec.Replicas,
					EnableExecution:         alamedaScaler.IsEnableExecution(),
					ResourceK8sMaxReplicas:  *msSpec.MaxReplicas,
					ResourceK8sMinReplicas:  *msSpec.MinReplicas,
					ScaleUpDelay:            *msSpec.ScaleUpDelay,
					ScaleDownDelay:          *msSpec.ScaleDownDelay,
				})
				break
			}
		}
	}
	err = datahubClient.Create(&mss)
	if err != nil {
		return fmt.Errorf("Create machineset (%s/%s) failed: %s",
			alamedaScaler.GetNamespace(), alamedaScaler.GetName(), err.Error())
	}
	return nil
}
