package metrics

import (
	"github.com/golang/protobuf/ptypes"
	"prophetstor.com/alameda/datahub/pkg/apis"
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/resources"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	Database "prophetstor.com/alameda/pkg/database/common"
	ApiCommon "prophetstor.com/api/datahub/common"
	ApiMetrics "prophetstor.com/api/datahub/metrics"
)

type CreatePodMetricsRequestExtended struct {
	*ApiMetrics.CreatePodMetricsRequest
}

func (r *CreatePodMetricsRequestExtended) Validate() error {
	return nil
}

func (r *CreatePodMetricsRequestExtended) ProduceMetrics() DaoMetricTypes.PodMetricMap {
	podMetricMap := DaoMetricTypes.NewPodMetricMap()

	rateRange := int64(5)
	if r.GetRateRange() != 0 {
		rateRange = int64(r.GetRateRange())
	}

	for _, pod := range r.GetPodMetrics() {
		podMetric := DaoMetricTypes.NewPodMetric()
		podMetric.ObjectMeta = resources.NewObjectMeta(pod.GetObjectMeta())
		podMetric.RateRange = rateRange

		for _, container := range pod.GetContainerMetrics() {
			containerMetric := DaoMetricTypes.NewContainerMetric()
			containerMetric.ObjectMeta.ObjectMeta = resources.NewObjectMeta(pod.GetObjectMeta())
			containerMetric.ObjectMeta.PodName = podMetric.ObjectMeta.Name
			containerMetric.ObjectMeta.Name = container.GetName()
			containerMetric.RateRange = rateRange

			for _, data := range container.GetMetricData() {
				metricType := MetricTypeNameMap[data.GetMetricType()]
				for _, sample := range data.GetData() {
					timestamp, err := ptypes.Timestamp(sample.GetTime())
					if err != nil {
						scope.Error(" failed: " + err.Error())
					}
					sample := FormatTypes.Sample{
						Timestamp: timestamp,
						Value:     sample.GetNumValue(),
					}
					containerMetric.AddSample(metricType, sample)
				}
			}

			podMetric.ContainerMetricMap.AddContainerMetric(containerMetric)
		}

		podMetricMap.AddPodMetric(podMetric)
	}

	return podMetricMap
}

type ListPodMetricsRequestExtended struct {
	Request *ApiMetrics.ListPodMetricsRequest
}

func (r *ListPodMetricsRequestExtended) Validate() error {
	return nil
}

func (r *ListPodMetricsRequestExtended) SetDefaultWithMetricsDBType(dbType MetricsDBType) {
	q := normalizeListMetricsRequestQueryConditionWthMetricsDBType(r.Request.QueryCondition, dbType)
	q.TimeRange.AggregateFunction = ApiCommon.TimeRange_MAX
	r.Request.QueryCondition = q
}

func (r *ListPodMetricsRequestExtended) SetRollupFunction(metricsConfig *apis.MetricsConfig) {
	r.Request.QueryCondition.Function = newFunction(metricsConfig)
}

func (r *ListPodMetricsRequestExtended) ProduceRequest() DaoMetricTypes.ListPodMetricsRequest {
	request := DaoMetricTypes.NewListPodMetricsRequest()
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	request.RateRange = 5
	if r.Request.GetRateRange() != 0 {
		request.RateRange = int64(r.Request.GetRateRange())
	}
	objectMetas := make([]*metadata.ObjectMeta, len(r.Request.GetObjectMeta()))
	for i, objectMeta := range r.Request.GetObjectMeta() {
		copyObjectMeta := objectMeta
		o := resources.NewObjectMeta(copyObjectMeta)
		if o.IsEmpty() {
			objectMetas = nil
			break
		}
		objectMetas[i] = &o
	}
	request.ObjectMetas = objectMetas
	metricTypes := make([]FormatEnum.MetricType, 0)
	for _, metricType := range r.Request.GetMetricTypes() {
		metricTypes = append(metricTypes, MetricTypeNameMap[metricType])
	}
	if len(metricTypes) == 0 {
		metricTypes = append(metricTypes, MetricTypeNameMap[ApiCommon.MetricType_CPU_MILLICORES_USAGE])
		metricTypes = append(metricTypes, MetricTypeNameMap[ApiCommon.MetricType_MEMORY_BYTES_USAGE])
	}
	request.MetricTypes = metricTypes
	if r.Request.QueryCondition.Function != nil {
		switch r.Request.QueryCondition.Function.Type {
		case ApiCommon.FunctionType_FUNCTIONTYPE_MEAN:
			request.AggregateOverTimeFunction = Database.AvgOverTime
		case ApiCommon.FunctionType_FUNCTIONTYPE_MAX:
			request.AggregateOverTimeFunction = Database.MaxOverTime
		case ApiCommon.FunctionType_FUNCTIONTYPE_PERCENTILE:
			request.AggregateOverTimeFunction = Database.PercentileOverTime
		default:
			request.AggregateOverTimeFunction = Database.None
		}
	}
	return request
}
