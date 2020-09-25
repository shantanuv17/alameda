package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiMetrics "prophetstor.com/api/datahub/metrics"
)

func (s *ServiceV1alpha1) CreateMetrics(ctx context.Context, in *ApiMetrics.CreateMetricsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateMetrics grpc function: " + AlamedaUtils.InterfaceToString(in))

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListMetrics(ctx context.Context, in *ApiMetrics.ListMetricsRequest) (*ApiMetrics.ListMetricsResponse, error) {
	response := ApiMetrics.ListMetricsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Metrics: &ApiMetrics.Metric{},
	}

	return &response, nil
}
