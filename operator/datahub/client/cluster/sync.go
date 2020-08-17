package cluster

import (
	"fmt"

	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(client client.Client, datahubClient *datahubpkg.Client) error {

	clusterUID, err := k8sutils.GetClusterUID(client)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	datahubClusterRepo := NewClusterRepository(datahubClient, clusterUID)

	if err := datahubClusterRepo.CreateClusters([]*datahub_resources.Cluster{
		&datahub_resources.Cluster{
			ObjectMeta: &datahub_resources.ObjectMeta{
				Name: clusterUID,
			},
		},
	}); err != nil {
		return fmt.Errorf(
			"Sync cluster with datahub failed due to register cluster failed: %s",
			err.Error())
	}

	return nil
}
