package metrics

import (
	"fmt"
	DaoClusterTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	InfluxSchemas "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
)

type Controller struct {
	InfluxDBConfig InfluxDB.Config
}

func NewControllerWithConfig(config InfluxDB.Config) *Controller {
	return &Controller{InfluxDBConfig: config}
}

func (p *Controller) GetMetricMap(metricType enumconv.MetricType, controllers []*DaoClusterTypes.Controller, req DaoMetricTypes.ListControllerMetricsRequest) (DaoMetricTypes.ControllerMetricMap, error) {
	metricMap := DaoMetricTypes.NewControllerMetricMap()

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(InfluxSchemas.Metric, "cluster_status", "container")[0]
	m := schema.GetMeasurement("", metricTypeMapTable[metricType], InfluxSchemas.ResourceBoundaryUndefined, InfluxSchemas.ResourceQuotaUndefined)
	measurement := InfluxDB.NewMeasurement(InfluxSchemas.DatabaseNameMap[InfluxSchemas.Metric], m, p.InfluxDBConfig)

	for _, controller := range controllers {
		p.rebuildQueryCondition(controller, req.QueryCondition.SubQuery)

		groups, err := measurement.Read(InfluxDB.NewQuery(&req.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.ControllerMetricMap{}, err
		}

		if len(groups) > 0 {
			metric := p.genMetric(metricType, controller, groups)
			metricMap.AddControllerMetric(metric)
		}
	}

	return metricMap, nil
}

func (p *Controller) rebuildQueryCondition(controller *DaoClusterTypes.Controller, queryCondition *DBCommon.QueryCondition) {
	condition := DBCommon.Condition{}
	condition.Keys = []string{"pod_name", "pod_namespace", "cluster_name"}
	condition.Values = append(condition.Values, fmt.Sprintf(PodNameRegularExpression[controller.Kind], controller.ObjectMeta.Name))
	condition.Values = append(condition.Values, controller.ObjectMeta.Namespace)
	condition.Values = append(condition.Values, controller.ObjectMeta.ClusterName)
	condition.Operators = []string{"=~", "=", "="}
	condition.Types = []DBCommon.DataType{DBCommon.String, DBCommon.String, DBCommon.String}
	queryCondition.WhereCondition = append(queryCondition.WhereCondition, &condition)
}

func (p *Controller) genMetric(metricType enumconv.MetricType, controller *DaoClusterTypes.Controller, groups []*DBCommon.Group) *DaoMetricTypes.ControllerMetric {
	metric := DaoMetricTypes.NewControllerMetric()
	metric.ObjectMeta = DaoMetricTypes.ControllerObjectMeta{
		ObjectMeta: *controller.ObjectMeta,
		Kind:       controller.Kind,
	}
	for _, row := range groups[0].Rows {
		sample := types.Sample{Timestamp: *row.Time, Value: row.Values[0]}
		metric.AddSample(metricType, sample)
	}
	return metric
}
