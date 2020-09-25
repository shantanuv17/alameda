package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	DaoClusterStatus "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/prometheus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
)

func NewControllerMetricsReaderDAO(config config.Config) types.ControllerMetricsDAO {
	switch config.Apis.Metrics.Source {
	case "influxdb":
		return influxdb.NewControllerMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewControllerMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewControllerDAO(config), config.ClusterUID)
	default:
		return prometheus.NewControllerMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewControllerDAO(config), config.ClusterUID)
	}
}

func NewControllerMetricsWriterDAO(config config.Config) types.ControllerMetricsDAO {
	switch config.Apis.Metrics.Target {
	case "influxdb":
		return influxdb.NewControllerMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewControllerMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewControllerDAO(config), config.ClusterUID)
	default:
		return prometheus.NewControllerMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewControllerDAO(config), config.ClusterUID)
	}
}
