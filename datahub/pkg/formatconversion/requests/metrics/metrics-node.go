package metrics

import (
	"github.com/containers-ai/alameda/datahub/pkg/apis"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/common"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/resources"
	FormatTypes "github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	"github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	Database "github.com/containers-ai/alameda/pkg/database/common"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	ApiMetrics "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/metrics"
	"github.com/golang/protobuf/ptypes"
)

type CreateNodeMetricsRequestExtended struct {
	ApiMetrics.CreateNodeMetricsRequest
}

func (r *CreateNodeMetricsRequestExtended) Validate() error {
	return nil
}

func (r *CreateNodeMetricsRequestExtended) ProduceMetrics() DaoMetricTypes.NodeMetricMap {
	nodeMetricMap := DaoMetricTypes.NewNodeMetricMap()

	for _, node := range r.GetNodeMetrics() {
		nodeMetric := DaoMetricTypes.NewNodeMetric()
		nodeMetric.ObjectMeta = resources.NewObjectMeta(node.GetObjectMeta())

		for _, data := range node.GetMetricData() {
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
				nodeMetric.AddSample(metricType, sample)
			}
		}

		nodeMetricMap.AddNodeMetric(nodeMetric)
	}

	return nodeMetricMap
}

type ListNodeMetricsRequestExtended struct {
	Request *ApiMetrics.ListNodeMetricsRequest
}

func (r *ListNodeMetricsRequestExtended) Validate() error {
	return nil
}

func (r *ListNodeMetricsRequestExtended) SetDefaultWithMetricsDBType(dbType MetricsDBType) {
	q := normalizeListMetricsRequestQueryConditionWthMetricsDBType(*r.Request.QueryCondition, dbType)
	q.TimeRange.AggregateFunction = ApiCommon.TimeRange_MAX
	r.Request.QueryCondition = &q
}

func (r *ListNodeMetricsRequestExtended) SetRollupFunction(metricsConfig *apis.MetricsConfig) {
	r.Request.QueryCondition.Function = newFunction(metricsConfig)
}

func (r *ListNodeMetricsRequestExtended) ProduceRequest() DaoMetricTypes.ListNodeMetricsRequest {
	request := DaoMetricTypes.NewListNodeMetricsRequest()
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
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
