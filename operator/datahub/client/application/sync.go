package application

import (
	"context"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha2 "github.com/containers-ai/alameda/operator/api/v1alpha2"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func RemoveOutOfDate(client client.Client, datahubClient *datahubpkg.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	scalersList := autoscalingv1alpha2.AlamedaScalerList{}
	if err := client.List(ctx, &scalersList); err != nil {
		return errors.Errorf(
			"Sync applications with datahub failed due to list applications from cluster failed: %s", err.Error())
	}

	resourceApps := []entities.ResourceClusterStatusApplication{}
	if err := datahubClient.List(&resourceApps); err != nil {
		return err
	}

	oodApp := []entities.ResourceClusterStatusApplication{}
	for _, resourceApp := range resourceApps {
		isOOD := true
		for _, scaler := range scalersList.Items {
			if scaler.Spec.ClusterName == resourceApp.ClusterName &&
				scaler.GetNamespace() == resourceApp.Namespace &&
				scaler.GetName() == resourceApp.Name {
				isOOD = false
				break
			}
		}
		if isOOD {
			oodApp = append(oodApp, resourceApp)
		}
	}

	if len(oodApp) > 0 {
		if err := datahubClient.Delete(&oodApp); err != nil {
			return err
		}
	}

	return nil
}
