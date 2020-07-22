package scaler

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func DeleteV1Alpha2Scaler(
	datahubClient *datahubpkg.Client, k8sClient client.Client,
	scalerNamespace, scalerName string, enabledDA bool) error {

	namespaceEntityList := []entities.TargetClusterStatusCluster{}
	err := datahubClient.List(&namespaceEntityList, datahubpkg.Option{
		Entity: entities.TargetClusterStatusCluster{
			AlamedaScalerNamespace: scalerNamespace,
			AlamedaScalerName:      scalerName,
		},
		Fields: []string{"AlamedaScalerNamespace", "AlamedaScalerName"},
	})
	scalerClusterName := ""
	if len(namespaceEntityList) == 1 {
		scalerClusterName = namespaceEntityList[0].Name
	}

	err = datahubClient.DeleteByOpts(&entities.TargetClusterStatusCluster{},
		datahubpkg.Option{
			Entity: entities.TargetClusterStatusCluster{
				Name:                   scalerClusterName,
				AlamedaScalerNamespace: scalerNamespace,
				AlamedaScalerName:      scalerName,
			},
			Fields: []string{"Name", "AlamedaScalerNamespace", "AlamedaScalerName"},
		})
	if err != nil {
		return err
	}
	err = datahubClient.DeleteByOpts(&entities.TargetClusterStatusController{},
		datahubpkg.Option{
			Entity: entities.TargetClusterStatusController{
				ClusterName:            scalerClusterName,
				AlamedaScalerNamespace: scalerNamespace,
				AlamedaScalerName:      scalerName,
			},
			Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
		})
	if err != nil {
		return err
	}
	err = datahubClient.DeleteByOpts(&entities.TargetKafkaConsumerGroup{},
		datahubpkg.Option{
			Entity: entities.TargetKafkaConsumerGroup{
				ClusterName:            scalerClusterName,
				AlamedaScalerNamespace: scalerNamespace,
				AlamedaScalerName:      scalerName,
			},
			Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
		})
	if err != nil {
		return err
	}
	err = datahubClient.DeleteByOpts(&entities.TargetKafkaTopic{},
		datahubpkg.Option{
			Entity: entities.TargetKafkaTopic{
				ClusterName:            scalerClusterName,
				AlamedaScalerNamespace: scalerNamespace,
				AlamedaScalerName:      scalerName,
			},
			Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
		})
	if err != nil {
		return err
	}

	if !enabledDA {
		err = datahubClient.DeleteByOpts(&entities.ResourceClusterStatusContainer{},
			datahubpkg.Option{
				Entity: entities.ResourceClusterStatusContainer{
					Name:                   scalerClusterName,
					AlamedaScalerNamespace: scalerNamespace,
					AlamedaScalerName:      scalerName,
				},
				Fields: []string{"Name", "AlamedaScalerNamespace", "AlamedaScalerName"},
			})
		if err != nil {
			return err
		}
		err = datahubClient.DeleteByOpts(&entities.ResourceClusterStatusPod{},
			datahubpkg.Option{
				Entity: entities.ResourceClusterStatusPod{
					ClusterName:            scalerClusterName,
					AlamedaScalerNamespace: scalerNamespace,
					AlamedaScalerName:      scalerName,
				},
				Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
			})
		if err != nil {
			return err
		}
		err = datahubClient.DeleteByOpts(&entities.ResourceClusterStatusController{},
			datahubpkg.Option{
				Entity: entities.ResourceClusterStatusController{
					ClusterName:            scalerClusterName,
					AlamedaScalerNamespace: scalerNamespace,
					AlamedaScalerName:      scalerName,
				},
				Fields: []string{"ClusterName", "AlamedaScalerName", "AlamedaScalerName"},
			})
		if err != nil {
			return err
		}
	}

	// finally delete application
	err = datahubClient.DeleteByOpts(&entities.ResourceClusterStatusApplication{},
		datahubpkg.Option{
			Entity: entities.ResourceClusterStatusApplication{
				ClusterName: scalerClusterName,
				Namespace:   scalerNamespace,
				Name:        scalerName,
			},
			Fields: []string{"ClusterName", "Namespace", "Name"},
		})
	if err != nil {
		return err
	}

	return nil
}
