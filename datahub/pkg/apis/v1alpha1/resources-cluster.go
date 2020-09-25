package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoCluster "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/resources"
	resources2 "prophetstor.com/alameda/datahub/pkg/formatconversion/responses/resources"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiResources "prophetstor.com/api/datahub/resources"
)

func (s *ServiceV1alpha1) CreateClusters(ctx context.Context, in *ApiResources.CreateClustersRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateClusters grpc function: " + AlamedaUtils.InterfaceToString(in))

	if in.GetClusters() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	requestExtended := resources.CreateClustersRequestExtended{CreateClustersRequest: in}
	if requestExtended.Validate() != nil {
		return &status.Status{
			Code: int32(code.Code_INVALID_ARGUMENT),
		}, nil
	}

	clusterDAO := DaoCluster.NewClusterDAO(*s.Config)
	if err := clusterDAO.CreateClusters(requestExtended.ProduceClusters()); err != nil {
		scope.Errorf("failed to create clusters: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListClusters(ctx context.Context, in *ApiResources.ListClustersRequest) (*ApiResources.ListClustersResponse, error) {
	scope.Debug("Request received from ListClusters grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := resources.ListClustersRequestExtended{ListClustersRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiResources.ListClustersResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	clusterDAO := DaoCluster.NewClusterDAO(*s.Config)
	clsts, err := clusterDAO.ListClusters(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ListClusters failed: %+v", err)
		return &ApiResources.ListClustersResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	clusters := make([]*ApiResources.Cluster, 0)
	for _, clst := range clsts {
		clusterExtended := resources2.ClusterExtended{Cluster: clst}
		cluster := clusterExtended.ProduceCluster()
		clusters = append(clusters, cluster)
	}

	return &ApiResources.ListClustersResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Clusters: clusters,
	}, nil
}

func (s *ServiceV1alpha1) DeleteClusters(ctx context.Context, in *ApiResources.DeleteClustersRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteClusters grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := resources.DeleteClustersRequestExtended{DeleteClustersRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	namespaceDAO := DaoCluster.NewClusterDAO(*s.Config)
	if err := namespaceDAO.DeleteClusters(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete clusters: %+v", err)
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
