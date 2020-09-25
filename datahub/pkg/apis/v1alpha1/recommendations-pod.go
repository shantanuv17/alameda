package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoRecommendation "prophetstor.com/alameda/datahub/pkg/dao/interfaces/recommendations"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

// CreatePodRecommendations add pod recommendations information to database
func (s *ServiceV1alpha1) CreatePodRecommendations(ctx context.Context, in *ApiRecommendations.CreatePodRecommendationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreatePodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))
	containerDAO := DaoRecommendation.NewContainerRecommendationsDAO(*s.Config)
	if err := containerDAO.CreatePodRecommendations(in); err != nil {
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

// ListPodRecommendations list pod recommendations
func (s *ServiceV1alpha1) ListPodRecommendations(ctx context.Context, in *ApiRecommendations.ListPodRecommendationsRequest) (*ApiRecommendations.ListPodRecommendationsResponse, error) {
	scope.Debug("Request received from ListPodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	containerDAO := DaoRecommendation.NewContainerRecommendationsDAO(*s.Config)
	podRecommendations, err := containerDAO.ListPodRecommendations(in)
	if err != nil {
		scope.Error(err.Error())
		return &ApiRecommendations.ListPodRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	res := &ApiRecommendations.ListPodRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		PodRecommendations: podRecommendations,
	}
	scope.Debug("Response sent from ListPodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(res))
	return res, nil
}

// ListAvailablePodRecommendations list pod recommendations
func (s *ServiceV1alpha1) ListAvailablePodRecommendations(ctx context.Context, in *ApiRecommendations.ListPodRecommendationsRequest) (*ApiRecommendations.ListPodRecommendationsResponse, error) {
	scope.Debug("Request received from ListAvailablePodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	containerDAO := DaoRecommendation.NewContainerRecommendationsDAO(*s.Config)
	podRecommendations, err := containerDAO.ListAvailablePodRecommendations(in)
	if err != nil {
		scope.Error(err.Error())
		return &ApiRecommendations.ListPodRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	res := &ApiRecommendations.ListPodRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		PodRecommendations: podRecommendations,
	}
	scope.Debug("Response sent from ListPodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(res))
	return res, nil
}
