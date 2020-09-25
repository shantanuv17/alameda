package metrics

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"prophetstor.com/alameda/datahub/pkg/apis"
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/resources"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	ApiCommon "prophetstor.com/api/datahub/common"
	ApiMetrics "prophetstor.com/api/datahub/metrics"
)

type CreateApplicationMetricsRequestExtended struct {
	*ApiMetrics.CreateApplicationMetricsRequest
}

func (r *CreateApplicationMetricsRequestExtended) Validate() error {
	for _, m := range r.GetApplicationMetrics() {
		if m == nil || m.ObjectMeta == nil || m.ObjectMeta.Name == "" || m.ObjectMeta.Namespace == "" || m.ObjectMeta.ClusterName == "" {
			return errors.Errorf(`must provide "Name", "Namespace" and "ClusterName" in ObjectMeta`)
		}
	}
	return nil
}

func (r *CreateApplicationMetricsRequestExtended) ProduceMetrics() DaoMetricTypes.AppMetricMap {
	metricMap := DaoMetricTypes.NewAppMetricMap()

	for _, applicationMetric := range r.GetApplicationMetrics() {
		if applicationMetric == nil {
			continue
		}
		m := DaoMetricTypes.NewAppMetric()
		m.ObjectMeta = resources.NewObjectMeta(applicationMetric.GetObjectMeta())

		for _, data := range applicationMetric.GetMetricData() {
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
				m.AddSample(metricType, sample)
			}
		}

		metricMap.AddAppMetric(m)
	}

	return metricMap
}

type ListAppMetricsRequestExtended struct {
	Request *ApiMetrics.ListApplicationMetricsRequest
}

func (r *ListAppMetricsRequestExtended) Validate() error {
	return nil
}

func (r *ListAppMetricsRequestExtended) SetDefaultWithMetricsDBType(metricsConfig *apis.MetricsConfig) {
	q := normalizeListMetricsRequestQueryConditionWthMetricsDBType(r.Request.QueryCondition, metricsConfig.Source)
	q.TimeRange.AggregateFunction = ApiCommon.TimeRange_NONE
	r.Request.QueryCondition = q
}

func (r *ListAppMetricsRequestExtended) SetRollupFunction(metricsConfig *apis.MetricsConfig) {
	r.Request.QueryCondition.Function = newFunction(metricsConfig)
}

func (r *ListAppMetricsRequestExtended) ProduceRequest() DaoMetricTypes.ListAppMetricsRequest {
	request := DaoMetricTypes.ListAppMetricsRequest{}
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	subQuery := common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	request.SubQuery = &subQuery

	objectMetas := make([]metadata.ObjectMeta, len(r.Request.GetObjectMeta()))
	for i, objectMeta := range r.Request.GetObjectMeta() {
		copyObjectMeta := objectMeta
		o := resources.NewObjectMeta(copyObjectMeta)
		if o.IsEmpty() {
			objectMetas = nil
			break
		}
		objectMetas[i] = o
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
	return request
}
