package application

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/autoscaling/v1alpha1"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(client client.Client, datahubClient *datahubpkg.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	applicationList := autoscalingv1alpha1.AlamedaScalerList{}
	if err := client.List(ctx, &applicationList); err != nil {
		return errors.Errorf(
			"Sync applications with datahub failed due to list applications from cluster failed: %s", err.Error())
	}

	clusterUID, err := k8sutils.GetClusterUID(client)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	if len(applicationList.Items) > 0 {
		apps := []entities.ResourceClusterStatusApplication{}
		for idx := range applicationList.Items {
			entity := entities.ResourceClusterStatusApplication{
				ClusterName: clusterUID,
				Namespace:   applicationList.Items[idx].Namespace,
				Name:        applicationList.Items[idx].Name,
				ScalingTool: GetAlamedaScalerDatahubScalingTypeStr(applicationList.Items[idx]),
				Type:        applicationList.Items[idx].GetType(),
			}
			if applicationList.Items[idx].GetType() == autoscalingv1alpha1.AlamedaScalerTypeKafka {
				appSpecBin, _ := json.Marshal(applicationList.Items[idx].Spec.Kafka)
				entity.AppSpec = string(appSpecBin)
			}
			apps = append(apps, entity)
		}
		if err := datahubClient.Create(&apps); err != nil {
			return fmt.Errorf(
				"Sync applications with datahub failed due to register application failed: %s", err.Error())
		}
	}

	// Clean up unexisting applications from Datahub
	existingApplicationMap := make(map[string]bool)
	for _, application := range applicationList.Items {
		existingApplicationMap[fmt.Sprintf("%s/%s",
			application.GetNamespace(), application.GetName())] = true
	}

	existingApps := []entities.ResourceClusterStatusApplication{}
	err = datahubClient.List(&existingApps)
	if err != nil {
		return fmt.Errorf(
			"Sync applications with datahub failed due to list applications from datahub failed: %s",
			err.Error())
	}
	applicationsNeedDeleting := []entities.ResourceClusterStatusApplication{}
	for _, n := range existingApps {
		if _, exist := existingApplicationMap[fmt.Sprintf("%s/%s",
			n.Namespace, n.Name)]; exist {
			continue
		}
		applicationsNeedDeleting = append(applicationsNeedDeleting, n)
	}
	if len(applicationsNeedDeleting) > 0 {
		err = datahubClient.Delete(&applicationsNeedDeleting)
		if err != nil {
			return errors.Wrap(err, "delete applications from Datahub failed")
		}
	}

	return nil
}

func GetAlamedaScalerDatahubScalingTypeStr(
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) string {
	scalingType := datahub_resources.ScalingTool_name[int32(datahub_resources.ScalingTool_SCALING_TOOL_UNDEFINED)]
	switch alamedaScaler.Spec.ScalingTool.Type {
	case autoscalingv1alpha1.ScalingToolTypeVPA:
		scalingType = datahub_resources.ScalingTool_name[int32(datahub_resources.ScalingTool_VPA)]
	case autoscalingv1alpha1.ScalingToolTypeHPA:
		scalingType = datahub_resources.ScalingTool_name[int32(datahub_resources.ScalingTool_HPA)]
	case autoscalingv1alpha1.ScalingToolTypeDefault:
		scalingType = datahub_resources.ScalingTool_name[int32(datahub_resources.ScalingTool_NONE)]
	}
	return scalingType
}
