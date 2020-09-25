package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	ApiRec "prophetstor.com/api/datahub/recommendations"
)

func (s *ServiceV1alpha1) CreateRecommendations(ctx context.Context, in *ApiRec.CreateRecommendationsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListRecommendations(ctx context.Context, in *ApiRec.ListRecommendationsRequest) (*ApiRec.ListRecommendationsResponse, error) {
	response := ApiRec.ListRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Recommendations: &ApiRec.Recommendation{},
	}

	return &response, nil
}
