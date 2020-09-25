package cluster

import (
	"fmt"

	"github.com/pkg/errors"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	k8sutils "prophetstor.com/alameda/pkg/utils/kubernetes"
	datahub_resources "prophetstor.com/api/datahub/resources"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(client client.Client, datahubClient *datahubpkg.Client) error {

	clusterUID, err := k8sutils.GetClusterUID(client)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	datahubClusterRepo := NewClusterRepository(datahubClient, clusterUID)

	if err := datahubClusterRepo.CreateClusters([]*datahub_resources.Cluster{{
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
