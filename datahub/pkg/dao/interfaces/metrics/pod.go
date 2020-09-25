package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	DaoClusterStatus "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/prometheus"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
)

func NewPodMetricsReaderDAO(config config.Config) types.PodMetricsDAO {
	switch config.Apis.Metrics.Source {
	case "influxdb":
		return influxdb.NewPodMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewPodMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewPodDAO(config), config.ClusterUID)
	default:
		return prometheus.NewPodMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewPodDAO(config), config.ClusterUID)
	}
}

func NewPodMetricsWriterDAO(config config.Config) types.PodMetricsDAO {
	switch config.Apis.Metrics.Target {
	case "influxdb":
		return influxdb.NewPodMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewPodMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewPodDAO(config), config.ClusterUID)
	default:
		return prometheus.NewPodMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewPodDAO(config), config.ClusterUID)
	}
}
