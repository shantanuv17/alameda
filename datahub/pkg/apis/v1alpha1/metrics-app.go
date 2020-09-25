package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoMetrics "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/metrics"
	metrics2 "prophetstor.com/alameda/datahub/pkg/formatconversion/responses/metrics"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiMetrics "prophetstor.com/api/datahub/metrics"
)

func (s *ServiceV1alpha1) CreateApplicationMetrics(ctx context.Context, in *ApiMetrics.CreateApplicationMetricsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateApplicationMetrics grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExtended := metrics.CreateApplicationMetricsRequestExtended{CreateApplicationMetricsRequest: in}
	if err := requestExtended.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	metricDAO := DaoMetrics.NewAppMetricsWriterDAO(*s.Config)
	err := metricDAO.CreateMetrics(ctx, requestExtended.ProduceMetrics())
	if err != nil {
		scope.Errorf("failed to create application metrics: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListApplicationMetrics(ctx context.Context, in *ApiMetrics.ListApplicationMetricsRequest) (*ApiMetrics.ListApplicationMetricsResponse, error) {
	scope.Debug("Request received from ListApplicationMetrics grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExtended := metrics.ListAppMetricsRequestExtended{Request: in}
	if err := requestExtended.Validate(); err != nil {
		return &ApiMetrics.ListApplicationMetricsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}
	requestExtended.SetDefaultWithMetricsDBType(s.Config.Apis.Metrics)
	requestExtended.SetRollupFunction(s.Config.Apis.Metrics)

	metricsDao := DaoMetrics.NewAppMetricsReaderDAO(*s.Config)
	metricMap, err := metricsDao.ListMetrics(ctx, requestExtended.ProduceRequest())
	if err != nil {
		return &ApiMetrics.ListApplicationMetricsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}
	i := 0
	datahubAppMetrics := make([]*ApiMetrics.ApplicationMetric, len(metricMap.MetricMap))
	for _, metric := range metricMap.MetricMap {
		m := metrics2.AppMetricExtended{AppMetric: *metric}.ProduceMetrics()
		datahubAppMetrics[i] = m
		i++
	}

	return &ApiMetrics.ListApplicationMetricsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		ApplicationMetrics: datahubAppMetrics,
	}, nil
}
