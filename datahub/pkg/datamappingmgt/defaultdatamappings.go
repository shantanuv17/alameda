package datamappingmgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/defaults"
)

func DefaultDataMappingsInit() {
	dataSourceMgt := NewDataMappingManagement()
	dataMappings := make([]*datamapping.DataMapping, 0)

	// Metric
	dataMappings = append(dataMappings, defaults.MetricKafkaTopicCurrentOffset())
	dataMappings = append(dataMappings, defaults.MetricKafkaCGCurrentOffset())
	dataMappings = append(dataMappings, defaults.MetricKafkaCGLag())

	dataSourceMgt.AddDataMappings(dataMappings)
	dataSourceMgt.Flush()
}
