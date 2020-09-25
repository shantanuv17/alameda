package metrics

import (
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	K8sMetadata "prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
)

type NodeMemoryBytesUsageEntity struct {
	NodeName string
	Samples  []FormatTypes.Sample
}

// NodeMetric Build NodeMetric base on entity properties
func (e *NodeMemoryBytesUsageEntity) NodeMetric() DaoMetricTypes.NodeMetric {

	var (
		nodeMetric DaoMetricTypes.NodeMetric
	)

	nodeMetric = DaoMetricTypes.NodeMetric{
		ObjectMeta: K8sMetadata.ObjectMeta{
			Name: e.NodeName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeMemoryUsageBytes: e.Samples,
		},
	}

	return nodeMetric
}

type NodeCPUUsageMillicoresEntity struct {
	NodeName string
	Samples  []FormatTypes.Sample
}

// NodeMetric Build NodeMetric base on entity properties
func (e *NodeCPUUsageMillicoresEntity) NodeMetric() DaoMetricTypes.NodeMetric {

	var (
		nodeMetric DaoMetricTypes.NodeMetric
	)

	nodeMetric = DaoMetricTypes.NodeMetric{
		ObjectMeta: K8sMetadata.ObjectMeta{
			Name: e.NodeName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeCPUUsageSecondsPercentage: e.Samples,
		},
	}

	return nodeMetric
}
