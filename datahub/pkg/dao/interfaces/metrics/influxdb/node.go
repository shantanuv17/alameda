package influxdb

import (
	"context"

	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	RepoInfluxMetric "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/metrics"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	Utils "prophetstor.com/alameda/pkg/utils"
)

type NodeMetrics struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNodeMetricsWithConfig(config InfluxDB.Config) DaoMetricTypes.NodeMetricsDAO {
	return &NodeMetrics{InfluxDBConfig: config}
}

func (p *NodeMetrics) CreateMetrics(ctx context.Context, metrics DaoMetricTypes.NodeMetricMap) error {
	// Write node cpu metrics
	nodeCpuRepo := RepoInfluxMetric.NewNodeCpuRepositoryWithConfig(p.InfluxDBConfig)
	err := nodeCpuRepo.CreateMetrics(metrics.GetSamples(FormatEnum.MetricTypeCPUUsageSecondsPercentage))
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	// Write node memory metrics
	nodeMemoryRepo := RepoInfluxMetric.NewNodeMemoryRepositoryWithConfig(p.InfluxDBConfig)
	err = nodeMemoryRepo.CreateMetrics(metrics.GetSamples(FormatEnum.MetricTypeMemoryUsageBytes))
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *NodeMetrics) ListMetrics(ctx context.Context, req DaoMetricTypes.ListNodeMetricsRequest) (DaoMetricTypes.NodeMetricMap, error) {
	nodeMetricMap := DaoMetricTypes.NewNodeMetricMap()

	// Read node cpu metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeCPUUsageSecondsPercentage) {
		nodeCpuRepo := RepoInfluxMetric.NewNodeCpuRepositoryWithConfig(p.InfluxDBConfig)
		cpuMetrics, err := nodeCpuRepo.ListMetrics(req)
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.NewNodeMetricMap(), err
		}
		for _, nodeMetric := range cpuMetrics {
			nodeMetricMap.AddNodeMetric(nodeMetric)
		}
	}

	// Read node memory metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeMemoryUsageBytes) {
		nodeMemoryRepo := RepoInfluxMetric.NewNodeMemoryRepositoryWithConfig(p.InfluxDBConfig)
		memoryMetrics, err := nodeMemoryRepo.ListMetrics(req)
		if err != nil {
			scope.Error(err.Error())
			return DaoMetricTypes.NewNodeMetricMap(), err
		}
		for _, nodeMetric := range memoryMetrics {
			nodeMetricMap.AddNodeMetric(nodeMetric)
		}
	}

	return nodeMetricMap, nil
}
