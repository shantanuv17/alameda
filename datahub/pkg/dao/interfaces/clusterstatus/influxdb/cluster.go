package influxdb

import (
	DaoClusterTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	RepoInfluxCluster "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	Metadata "prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
)

type Cluster struct {
	InfluxDBConfig InfluxDB.Config
}

func NewClusterWithConfig(config InfluxDB.Config) DaoClusterTypes.ClusterDAO {
	return &Cluster{InfluxDBConfig: config}
}

func (p *Cluster) CreateClusters(clusters []*DaoClusterTypes.Cluster) error {
	clusterRepo := RepoInfluxCluster.NewClusterRepository(p.InfluxDBConfig)
	err := clusterRepo.CreateClusters(clusters)
	if err != nil {
		scope.Error(err.Error())
		return err
	}
	return nil
}

func (p *Cluster) ListClusters(request *DaoClusterTypes.ListClustersRequest) ([]*DaoClusterTypes.Cluster, error) {
	clusterRepo := RepoInfluxCluster.NewClusterRepository(p.InfluxDBConfig)
	clusters, err := clusterRepo.ListClusters(request)
	if err != nil {
		scope.Error(err.Error())
		return make([]*DaoClusterTypes.Cluster, 0), err
	}
	return clusters, nil
}

func (p *Cluster) DeleteClusters(request *DaoClusterTypes.DeleteClustersRequest) error {
	delNodeReq := p.genDeleteNodesRequest(request)
	delNamespacesReq := p.genDeleteNamespacesRequest(request)

	// Delete clusters
	clusterRepo := RepoInfluxCluster.NewClusterRepository(p.InfluxDBConfig)
	if err := clusterRepo.DeleteClusters(request); err != nil {
		scope.Error(err.Error())
		return err
	}

	// Delete nodes
	nodeDAO := NewNodeWithConfig(p.InfluxDBConfig)
	if err := nodeDAO.DeleteNodes(delNodeReq); err != nil {
		scope.Error(err.Error())
		return err
	}

	// Delete namespaces
	namespaceDAO := NewNamespaceWithConfig(p.InfluxDBConfig)
	if err := namespaceDAO.DeleteNamespaces(delNamespacesReq); err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *Cluster) genDeleteNodesRequest(request *DaoClusterTypes.DeleteClustersRequest) *DaoClusterTypes.DeleteNodesRequest {
	delNodesReq := DaoClusterTypes.NewDeleteNodesRequest()

	for _, objectMeta := range request.ObjectMeta {
		metadata := &Metadata.ObjectMeta{}
		metadata.ClusterName = objectMeta.Name

		delNodesReq.ObjectMeta = append(delNodesReq.ObjectMeta, metadata)
	}

	return delNodesReq
}

func (p *Cluster) genDeleteNamespacesRequest(request *DaoClusterTypes.DeleteClustersRequest) *DaoClusterTypes.DeleteNamespacesRequest {
	delNamespacesReq := DaoClusterTypes.NewDeleteNamespacesRequest()

	for _, objectMeta := range request.ObjectMeta {
		metadata := &Metadata.ObjectMeta{}
		metadata.ClusterName = objectMeta.Name

		delNamespacesReq.ObjectMeta = append(delNamespacesReq.ObjectMeta, metadata)
	}

	return delNamespacesReq
}
