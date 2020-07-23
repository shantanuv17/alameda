package influxdb

import (
	"context"

	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	DaoClusterStatusTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	RepoInfluxMetric "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/metrics"
	FormatEnum "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	Utils "github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/pkg/errors"
	"time"
)

type ClusterMetrics struct {
	InfluxDBConfig InfluxDB.Config
	ClusterDAO     DaoClusterStatusTypes.ClusterDAO
	ClusterRepo    *RepoInfluxMetric.Cluster
}

func NewClusterMetricsWithConfig(config InfluxDB.Config) DaoMetricTypes.ClusterMetricsDAO {
	clusterMetrics := ClusterMetrics{
		InfluxDBConfig: config,
		ClusterDAO:     influxdb.NewClusterWithConfig(config),
		ClusterRepo:    RepoInfluxMetric.NewClusterWithConfig(config),
	}
	return clusterMetrics
}

func (n ClusterMetrics) CreateMetrics(ctx context.Context, m DaoMetricTypes.ClusterMetricMap) error {
	// Write cluster cpu metrics
	cpuRepo := RepoInfluxMetric.NewClusterCPURepositoryWithConfig(n.InfluxDBConfig)
	err := cpuRepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeCPUUsageSecondsPercentage))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create cluster cpu metrics failed")
	}

	// Write cluster memory metrics
	memoryRepo := RepoInfluxMetric.NewClusterMemoryRepositoryWithConfig(n.InfluxDBConfig)
	err = memoryRepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeMemoryUsageBytes))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create cluster memory metrics failed")
	}
	return nil
}

func (n ClusterMetrics) ListMetrics(ctx context.Context, req DaoMetricTypes.ListClusterMetricsRequest) (DaoMetricTypes.ClusterMetricMap, error) {
	metricMap := DaoMetricTypes.NewClusterMetricMap()

	n.rebuildRequest(&req)

	clusters, err := n.listClusters(req)
	if err != nil {
		return DaoMetricTypes.ClusterMetricMap{}, errors.Wrap(err, "list cluster metadatas from request failed")
	}

	// Read cluster cpu metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeCPUUsageSecondsPercentage) {
		cpuMetricMap, err := n.ClusterRepo.GetMetricMap(FormatEnum.MetricTypeCPUUsageSecondsPercentage, clusters, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get cluster cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddClusterMetric(copyM)
		}

		/*cpuRepo := RepoInfluxMetric.NewClusterCPURepositoryWithConfig(n.InfluxDBConfig)
		cpuMetricMap, err := cpuRepo.GetClusterMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get cluster cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddClusterMetric(copyM)
		}*/
	}

	// Read cluster memory metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeMemoryUsageBytes) {
		memoryMetricMap, err := n.ClusterRepo.GetMetricMap(FormatEnum.MetricTypeMemoryUsageBytes, clusters, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get cluster memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddClusterMetric(copyM)
		}

		/*memoryRepo := RepoInfluxMetric.NewClusterMemoryRepositoryWithConfig(n.InfluxDBConfig)
		memoryMetricMap, err := memoryRepo.GetClusterMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get cluster memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddClusterMetric(copyM)
		}*/
	}

	return metricMap, nil
}

func (n ClusterMetrics) listClusters(req DaoMetricTypes.ListClusterMetricsRequest) ([]*DaoClusterStatusTypes.Cluster, error) {
	// Generate list resource cluster request
	listClustersReq := DaoClusterStatusTypes.NewListClustersRequest()
	for index := range req.ObjectMetas {
		listClustersReq.ObjectMeta = append(listClustersReq.ObjectMeta, &req.ObjectMetas[index])
	}

	clusters, err := n.ClusterDAO.ListClusters(listClustersReq)
	if err != nil {
		return nil, errors.Wrap(err, "list cluster metadatas failed")
	}

	return clusters, nil
}

func (n ClusterMetrics) rebuildRequest(req *DaoMetricTypes.ListClusterMetricsRequest) {
	req.Selects = append(req.Selects, "value")
	if req.AggregateOverTimeFunction != common.None {
		req.Function = NewFunction(common.FunctionTypeMap[req.AggregateOverTimeFunction])
		req.AggregateOverTimeFunction = common.None
	}

	step := time.Duration(15) * time.Second
	req.QueryCondition.SubQuery.StepTime = &step
	req.QueryCondition.SubQuery.Selects = append(req.QueryCondition.SubQuery.Selects, "value")
	req.QueryCondition.SubQuery.Groups = append(req.QueryCondition.SubQuery.Groups, "cluster_name")
	req.QueryCondition.SubQuery.AggregateOverTimeFunction = common.None
	req.QueryCondition.SubQuery.Function = NewFunction(common.FunctionSum)
}
