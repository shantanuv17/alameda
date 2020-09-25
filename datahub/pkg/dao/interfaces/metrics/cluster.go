package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	DaoClusterStatus "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/prometheus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
)

func NewClusterMetricsReaderDAO(config config.Config) types.ClusterMetricsDAO {
	switch config.Apis.Metrics.Source {
	case "influxdb":
		return influxdb.NewClusterMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewClusterMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewClusterDAO(config), DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	default:
		return prometheus.NewClusterMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewClusterDAO(config), DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	}
}

func NewClusterMetricsWriterDAO(config config.Config) types.ClusterMetricsDAO {
	switch config.Apis.Metrics.Target {
	case "influxdb":
		return influxdb.NewClusterMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewClusterMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewClusterDAO(config), DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	default:
		return prometheus.NewClusterMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewClusterDAO(config), DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	}
}
