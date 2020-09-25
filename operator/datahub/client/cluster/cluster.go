package cluster

import (
	"github.com/pkg/errors"
	"prophetstor.com/alameda/operator/datahub/client"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	datahub_resources "prophetstor.com/api/datahub/resources"
)

type ClusterRepository struct {
	datahubClient *datahubpkg.Client

	clusterUID string
}

// NewClusterRepository return ClusterRepository instance
func NewClusterRepository(datahubClient *datahubpkg.Client, clusterUID string) *ClusterRepository {
	return &ClusterRepository{
		datahubClient: datahubClient,
		clusterUID:    clusterUID,
	}
}

// CreateClusters creates clusters to datahub
func (repo *ClusterRepository) CreateClusters(arg interface{}) error {
	clusters := []*datahub_resources.Cluster{}
	if apps, ok := arg.([]*datahub_resources.Cluster); ok {
		clusters = apps
	}

	req := datahub_resources.CreateClustersRequest{
		Clusters: clusters,
	}

	if resp, err := repo.datahubClient.CreateClusters(&req); err != nil {
		return errors.Wrap(err, "create clusters to datahub failed")
	} else if _, err := client.IsResponseStatusOK(resp); err != nil {
		return errors.Wrap(err, "create clusters to datahub failed")
	}
	return nil
}

func (repo *ClusterRepository) ListClusters() ([]*datahub_resources.Cluster, error) {
	req := datahub_resources.ListClustersRequest{
		ObjectMeta: []*datahub_resources.ObjectMeta{
			{
				ClusterName: repo.clusterUID,
			},
		},
	}

	resp, err := repo.datahubClient.ListClusters(&req)
	if err != nil {
		return nil, errors.Wrap(err, "list clusters from Datahub failed")
	} else if resp == nil {
		return nil, errors.Errorf("list clusters from Datahub failed, receive nil response")
	} else if _, err := client.IsResponseStatusOK(resp.Status); err != nil {
		return nil, errors.Wrap(err, "list clusters from Datahub failed")
	}
	return resp.Clusters, nil
}
