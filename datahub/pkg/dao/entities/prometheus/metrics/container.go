package metrics

import (
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/alameda/pkg/database/prometheus"
)

type ContainerCPUUsageMillicoresEntity struct {
	PrometheusEntity prometheus.Entity

	Namespace     string
	PodName       string
	ContainerName string
	Samples       []FormatTypes.Sample
}

// ContainerMetric Build ContainerMetric base on entity properties
func (e *ContainerCPUUsageMillicoresEntity) ContainerMetric() DaoMetricTypes.ContainerMetric {

	var (
		containerMetric DaoMetricTypes.ContainerMetric
	)

	containerMetric = DaoMetricTypes.ContainerMetric{
		ObjectMeta: DaoMetricTypes.ContainerMeta{
			ObjectMeta: metadata.ObjectMeta{
				Namespace: e.Namespace,
				Name:      e.ContainerName,
			},
			PodName: e.PodName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeCPUUsageSecondsPercentage: e.Samples,
		},
	}

	return containerMetric
}

type ContainerMemoryUsageBytesEntity struct {
	PrometheusEntity prometheus.Entity

	Namespace     string
	PodName       string
	ContainerName string
	Samples       []FormatTypes.Sample
}

// ContainerMetric Build ContainerMetric base on entity properties
func (e *ContainerMemoryUsageBytesEntity) ContainerMetric() DaoMetricTypes.ContainerMetric {

	var (
		containerMetric DaoMetricTypes.ContainerMetric
	)

	containerMetric = DaoMetricTypes.ContainerMetric{
		ObjectMeta: DaoMetricTypes.ContainerMeta{
			ObjectMeta: metadata.ObjectMeta{
				Namespace: e.Namespace,
				Name:      e.ContainerName,
			},
			PodName: e.PodName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeMemoryUsageBytes: e.Samples,
		},
	}

	return containerMetric
}
