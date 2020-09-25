package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
)

func (s *ServiceV1alpha1) CreatePredictions(ctx context.Context, in *ApiPredictions.CreatePredictionsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListPredictions(ctx context.Context, in *ApiPredictions.ListPredictionsRequest) (*ApiPredictions.ListPredictionsResponse, error) {
	response := ApiPredictions.ListPredictionsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Predictions: &ApiPredictions.Prediction{},
	}

	return &response, nil
}
