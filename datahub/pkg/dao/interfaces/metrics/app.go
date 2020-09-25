package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	DaoClusterStatus "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/prometheus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
)

func NewAppMetricsReaderDAO(config config.Config) types.AppMetricsDAO {
	switch config.Apis.Metrics.Source {
	case "influxdb":
		return influxdb.NewAppMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewAppMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewApplicationDAO(config), config.ClusterUID)
	default:
		return prometheus.NewAppMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewApplicationDAO(config), config.ClusterUID)
	}
}

func NewAppMetricsWriterDAO(config config.Config) types.AppMetricsDAO {
	switch config.Apis.Metrics.Target {
	case "influxdb":
		return influxdb.NewAppMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewAppMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewApplicationDAO(config), config.ClusterUID)
	default:
		return prometheus.NewAppMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewApplicationDAO(config), config.ClusterUID)
	}
}
