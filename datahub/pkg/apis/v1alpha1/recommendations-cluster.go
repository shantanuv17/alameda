package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoRecommendation "prophetstor.com/alameda/datahub/pkg/dao/interfaces/recommendations"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

func (s *ServiceV1alpha1) CreateClusterRecommendations(ctx context.Context, in *ApiRecommendations.CreateClusterRecommendationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateClusterRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	clusterRecommendationList := in.GetClusterRecommendations()
	clusterDAO := DaoRecommendation.NewClusterRecommendationsDAO(*s.Config)
	err := clusterDAO.CreateRecommendations(clusterRecommendationList)

	if err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, err
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListClusterRecommendations(ctx context.Context, in *ApiRecommendations.ListClusterRecommendationsRequest) (*ApiRecommendations.ListClusterRecommendationsResponse, error) {
	scope.Debug("Request received from ListClusterRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	clusterDAO := DaoRecommendation.NewClusterRecommendationsDAO(*s.Config)
	clusterRecommendations, err := clusterDAO.ListRecommendations(in)
	if err != nil {
		scope.Errorf("api ListClusterRecommendations failed: %v", err)
		response := &ApiRecommendations.ListClusterRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
			ClusterRecommendations: clusterRecommendations,
		}
		return response, nil
	}

	response := &ApiRecommendations.ListClusterRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		ClusterRecommendations: clusterRecommendations,
	}

	return response, nil
}
