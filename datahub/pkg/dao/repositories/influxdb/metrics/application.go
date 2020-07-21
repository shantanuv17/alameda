package metrics

import (
	DaoClusterTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InternalSchemas "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

type Application struct {
	InfluxDBConfig InternalInflux.Config
}

func NewApplicationWithConfig(config InternalInflux.Config) *Application {
	return &Application{InfluxDBConfig: config}
}

func (p *Application) GetMetricMap(metricType enumconv.MetricType, applications []*DaoClusterTypes.Application, req DaoMetricTypes.ListAppMetricsRequest) (DaoMetricTypes.AppMetricMap, error) {
	metricMap := DaoMetricTypes.NewAppMetricMap()

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(InternalSchemas.Metric, "cluster_status", "container")[0]
	m := schema.GetMeasurement("", metricTypeMapTable[metricType], InternalSchemas.ResourceBoundaryUndefined, InternalSchemas.ResourceQuotaUndefined)
	measurement := InternalInflux.NewMeasurement(SchemaMgt.DatabaseNameMap[InternalSchemas.Metric], m, p.InfluxDBConfig)

	for _, application := range applications {
		// List pods which are belonged to this application
		pods, err := ListPodsByApplication(p.InfluxDBConfig, application)
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.AppMetricMap{}, err
		}

		p.rebuildQueryCondition(pods, req.QueryCondition.SubQuery)

		groups, err := measurement.Read(InternalInflux.NewQuery(&req.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.AppMetricMap{}, err
		}

		if len(groups) > 0 {
			metric := p.genMetric(metricType, application, groups)
			metricMap.AddAppMetric(metric)
		}
	}

	return metricMap, nil
}

func (p *Application) rebuildQueryCondition(pods []*DaoClusterTypes.Pod, queryCondition *common.QueryCondition) {
	queryCondition.WhereCondition = make([]*common.Condition, 0)

	for _, pod := range pods {
		condition := common.Condition{}
		condition.Keys = []string{"pod_name", "pod_namespace", "cluster_name"}
		condition.Values = []string{pod.ObjectMeta.Name, pod.ObjectMeta.Namespace, pod.ObjectMeta.ClusterName}
		condition.Operators = []string{"=", "=", "="}
		condition.Types = []common.DataType{common.String, common.String, common.String}
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, &condition)
	}
}

func (p *Application) genMetric(metricType enumconv.MetricType, application *DaoClusterTypes.Application, groups []*common.Group) *DaoMetricTypes.AppMetric {
	metric := DaoMetricTypes.NewAppMetric()
	metric.ObjectMeta = *application.ObjectMeta
	for _, row := range groups[0].Rows {
		sample := types.Sample{Timestamp: *row.Time, Value: row.Values[0]}
		metric.AddSample(metricType, sample)
	}
	return metric
}
