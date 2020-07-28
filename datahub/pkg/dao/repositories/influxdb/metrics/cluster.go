package metrics

import (
	DaoClusterTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	InfluxSchemas "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
)

type Cluster struct {
	InfluxDBConfig InfluxDB.Config
}

func NewClusterWithConfig(config InfluxDB.Config) *Cluster {
	return &Cluster{InfluxDBConfig: config}
}

func (p *Cluster) GetMetricMap(metricType enumconv.MetricType, clusters []*DaoClusterTypes.Cluster, req DaoMetricTypes.ListClusterMetricsRequest) (DaoMetricTypes.ClusterMetricMap, error) {
	metricMap := DaoMetricTypes.NewClusterMetricMap()

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(InfluxSchemas.Metric, "cluster_status", "node")[0]
	m := schema.GetMeasurement("", metricTypeMapTable[metricType], InfluxSchemas.ResourceBoundaryUndefined, InfluxSchemas.ResourceQuotaUndefined)
	measurement := InfluxDB.NewMeasurement(InfluxSchemas.DatabaseNameMap[InfluxSchemas.Metric], m, p.InfluxDBConfig)

	for _, cluster := range clusters {
		// List nodes which are belonged to this cluster
		nodes, err := ListNodesByCluster(p.InfluxDBConfig, cluster)
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.ClusterMetricMap{}, err
		}

		p.rebuildQueryCondition(nodes, req.QueryCondition.SubQuery)

		groups, err := measurement.Read(InfluxDB.NewQuery(&req.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.ClusterMetricMap{}, err
		}

		if len(groups) > 0 {
			metric := p.genMetric(metricType, cluster, groups)
			metricMap.AddClusterMetric(metric)
		}
	}

	return metricMap, nil
}

func (p *Cluster) rebuildQueryCondition(nodes []*DaoClusterTypes.Node, queryCondition *DBCommon.QueryCondition) {
	queryCondition.WhereCondition = make([]*DBCommon.Condition, 0)

	for _, node := range nodes {
		condition := DBCommon.Condition{}
		condition.Keys = []string{"name", "cluster_name"}
		condition.Values = []string{node.ObjectMeta.Name, node.ObjectMeta.ClusterName}
		condition.Operators = []string{"=", "="}
		condition.Types = []DBCommon.DataType{DBCommon.String, DBCommon.String}
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, &condition)
	}
}

func (p *Cluster) genMetric(metricType enumconv.MetricType, cluster *DaoClusterTypes.Cluster, groups []*DBCommon.Group) *DaoMetricTypes.ClusterMetric {
	metric := DaoMetricTypes.NewClusterMetric()
	metric.ObjectMeta = *cluster.ObjectMeta
	for _, row := range groups[0].Rows {
		sample := types.Sample{Timestamp: *row.Time, Value: row.Values[0]}
		metric.AddSample(metricType, sample)
	}
	return metric
}
