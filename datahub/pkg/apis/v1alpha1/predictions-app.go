package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoPrediction "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/predictions"
	predictions2 "prophetstor.com/alameda/datahub/pkg/formatconversion/responses/predictions"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
)

func (s *ServiceV1alpha1) CreateApplicationPredictions(ctx context.Context, in *ApiPredictions.CreateApplicationPredictionsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateApplicationPredictions grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExtended := predictions.CreateApplicationPredictionsRequestExtended{CreateApplicationPredictionsRequest: in}
	if requestExtended.Validate() != nil {
		return &status.Status{
			Code: int32(code.Code_INVALID_ARGUMENT),
		}, nil
	}

	predictionDAO := DaoPrediction.NewApplicationPredictionsDAO(*s.Config)
	err := predictionDAO.CreatePredictions(requestExtended.ProducePredictions())
	if err != nil {
		scope.Errorf("failed to create application predictions: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListApplicationPredictions(ctx context.Context, in *ApiPredictions.ListApplicationPredictionsRequest) (*ApiPredictions.ListApplicationPredictionsResponse, error) {
	scope.Debug("Request received from ListApplicationPredictions grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := predictions.ListApplicationPredictionsRequestExtended{Request: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiPredictions.ListApplicationPredictionsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	predictionDAO := DaoPrediction.NewApplicationPredictionsDAO(*s.Config)
	applicationsPredictionMap, err := predictionDAO.ListPredictions(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ListApplicationPredictions failed: %+v", err)
		return &ApiPredictions.ListApplicationPredictionsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	datahubApplicationPredictions := make([]*ApiPredictions.ApplicationPrediction, 0)
	for _, applicationPrediction := range applicationsPredictionMap.MetricMap {
		applicationPredictionExtended := predictions2.ApplicationPredictionExtended{ApplicationPrediction: applicationPrediction}
		datahubApplicationPrediction := applicationPredictionExtended.ProducePredictions()
		datahubApplicationPredictions = append(datahubApplicationPredictions, datahubApplicationPrediction)
	}

	return &ApiPredictions.ListApplicationPredictionsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		ApplicationPredictions: datahubApplicationPredictions,
	}, nil
}
