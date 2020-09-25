package influxdb

import (
	DaoClusterTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	RepoInfluxCluster "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	Metadata "prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
)

type Namespace struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNamespaceWithConfig(config InfluxDB.Config) DaoClusterTypes.NamespaceDAO {
	return &Namespace{InfluxDBConfig: config}
}

func (p *Namespace) CreateNamespaces(namespaces []*DaoClusterTypes.Namespace) error {
	namespaceRepo := RepoInfluxCluster.NewNamespaceRepository(p.InfluxDBConfig)
	err := namespaceRepo.CreateNamespaces(namespaces)
	if err != nil {
		scope.Error(err.Error())
		return err
	}
	return nil
}

func (p *Namespace) ListNamespaces(request *DaoClusterTypes.ListNamespacesRequest) ([]*DaoClusterTypes.Namespace, error) {
	namespaceRepo := RepoInfluxCluster.NewNamespaceRepository(p.InfluxDBConfig)
	namespaces, err := namespaceRepo.ListNamespaces(request)
	if err != nil {
		scope.Error(err.Error())
		return make([]*DaoClusterTypes.Namespace, 0), err
	}
	return namespaces, nil
}

func (p *Namespace) DeleteNamespaces(request *DaoClusterTypes.DeleteNamespacesRequest) error {
	delApplicationsReq := p.genDeleteApplicationsRequest(request)

	// Delete namespaces
	namespaceRepo := RepoInfluxCluster.NewNamespaceRepository(p.InfluxDBConfig)
	if err := namespaceRepo.DeleteNamespaces(request); err != nil {
		scope.Error(err.Error())
		return err
	}

	// Delete applications
	applicationDAO := NewApplicationWithConfig(p.InfluxDBConfig)
	if err := applicationDAO.DeleteApplications(delApplicationsReq); err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *Namespace) genDeleteApplicationsRequest(request *DaoClusterTypes.DeleteNamespacesRequest) *DaoClusterTypes.DeleteApplicationsRequest {
	delApplicationsReq := DaoClusterTypes.NewDeleteApplicationsRequest()

	for _, objectMeta := range request.ObjectMeta {
		metadata := &Metadata.ObjectMeta{}
		metadata.Namespace = objectMeta.Name
		metadata.ClusterName = objectMeta.ClusterName

		applicationObjectMeta := DaoClusterTypes.NewApplicationObjectMeta(metadata)
		delApplicationsReq.ApplicationObjectMeta = append(delApplicationsReq.ApplicationObjectMeta, applicationObjectMeta)
	}

	return delApplicationsReq
}
