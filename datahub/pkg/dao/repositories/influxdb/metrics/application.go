package metrics

import (
	"fmt"
	DaoClusterTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	SchemaMgt "prophetstor.com/alameda/datahub/pkg/schemamgt"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	InfluxSchemas "prophetstor.com/alameda/pkg/database/influxdb/schemas"
)

type Application struct {
	InfluxDBConfig InfluxDB.Config
}

func NewApplicationWithConfig(config InfluxDB.Config) *Application {
	return &Application{InfluxDBConfig: config}
}

func (p *Application) GetMetricMap(metricType enumconv.MetricType, applications []*DaoClusterTypes.Application, req DaoMetricTypes.ListAppMetricsRequest) (DaoMetricTypes.AppMetricMap, error) {
	metricMap := DaoMetricTypes.NewAppMetricMap()

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(InfluxSchemas.Metric, "cluster_status", "container")[0]
	m := schema.GetMeasurement("", metricTypeMapTable[metricType], InfluxSchemas.ResourceBoundaryUndefined, InfluxSchemas.ResourceQuotaUndefined)
	measurement := InfluxDB.NewMeasurement(InfluxSchemas.DatabaseNameMap[InfluxSchemas.Metric], m, p.InfluxDBConfig)

	for _, application := range applications {
		// If no controller is found, return empty metrics data
		if len(application.Controllers) == 0 {
			metric := p.genMetric(metricType, application, nil)
			metricMap.AddAppMetric(metric)
			continue
		}

		p.rebuildQueryCondition(application.Controllers, req.QueryCondition.SubQuery)

		groups, err := measurement.Read(InfluxDB.NewQuery(&req.QueryCondition, measurement.Name))
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

func (p *Application) rebuildQueryCondition(controllers []*DaoClusterTypes.Controller, queryCondition *DBCommon.QueryCondition) {
	queryCondition.WhereCondition = make([]*DBCommon.Condition, 0)

	for _, controller := range controllers {
		condition := DBCommon.Condition{}
		condition.Keys = []string{"pod_name", "pod_namespace", "cluster_name"}
		condition.Values = append(condition.Values, fmt.Sprintf(PodNameRegularExpression[controller.Kind], controller.ObjectMeta.Name))
		condition.Values = append(condition.Values, controller.ObjectMeta.Namespace)
		condition.Values = append(condition.Values, controller.ObjectMeta.ClusterName)
		condition.Operators = []string{"=~", "=", "="}
		condition.Types = []DBCommon.DataType{DBCommon.String, DBCommon.String, DBCommon.String}
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, &condition)
	}
}

func (p *Application) genMetric(metricType enumconv.MetricType, application *DaoClusterTypes.Application, groups []*DBCommon.Group) *DaoMetricTypes.AppMetric {
	metric := DaoMetricTypes.NewAppMetric()
	metric.ObjectMeta = *application.ObjectMeta

	if groups == nil {
		metric.Metrics = make(map[enumconv.MetricType][]types.Sample)
		metric.Metrics[metricType] = make([]types.Sample, 0)
		return metric
	}

	for _, row := range groups[0].Rows {
		sample := types.Sample{Timestamp: *row.Time, Value: row.Values[0]}
		metric.AddSample(metricType, sample)
	}

	return metric
}
