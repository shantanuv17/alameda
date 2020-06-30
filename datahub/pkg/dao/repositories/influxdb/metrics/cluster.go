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

type Cluster struct {
	InfluxDBConfig InternalInflux.Config
}

func NewClusterWithConfig(config InternalInflux.Config) *Cluster {
	return &Cluster{InfluxDBConfig: config}
}

func (p *Cluster) GetMetricMap(metricType enumconv.MetricType, clusters []*DaoClusterTypes.Cluster, req DaoMetricTypes.ListClusterMetricsRequest) (DaoMetricTypes.ClusterMetricMap, error) {
	metricMap := DaoMetricTypes.NewClusterMetricMap()

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(InternalSchemas.Metric, "cluster_status", "node")[0]
	m := schema.GetMeasurement("", metricTypeMapTable[metricType], InternalSchemas.ResourceBoundaryUndefined, InternalSchemas.ResourceQuotaUndefined)
	measurement := InternalInflux.NewMeasurement(SchemaMgt.DatabaseNameMap[InternalSchemas.Metric], m, p.InfluxDBConfig)

	for _, cluster := range clusters {
		// List nodes which are belonged to this cluster
		nodes, err := ListNodesByCluster(p.InfluxDBConfig, cluster)
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.ClusterMetricMap{}, err
		}

		p.rebuildQueryCondition(nodes, req.QueryCondition.SubQuery)

		groups, err := measurement.Read(InternalInflux.NewQuery(&req.QueryCondition, measurement.Name))
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

func (p *Cluster) rebuildQueryCondition(nodes []*DaoClusterTypes.Node, queryCondition *common.QueryCondition) {
	queryCondition.WhereCondition = make([]*common.Condition, 0)

	for _, node := range nodes {
		condition := common.Condition{}
		condition.Keys = []string{"name", "cluster_name"}
		condition.Values = []string{node.ObjectMeta.Name, node.ObjectMeta.ClusterName}
		condition.Operators = []string{"=", "="}
		condition.Types = []common.DataType{common.String, common.String}
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, &condition)
	}
}

func (p *Cluster) genMetric(metricType enumconv.MetricType, cluster *DaoClusterTypes.Cluster, groups []*common.Group) *DaoMetricTypes.ClusterMetric {
	metric := DaoMetricTypes.NewClusterMetric()
	metric.ObjectMeta = *cluster.ObjectMeta
	for _, row := range groups[0].Rows {
		sample := types.Sample{Timestamp: *row.Time, Value: row.Values[0]}
		metric.AddSample(metricType, sample)
	}
	return metric
}
