package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func MetricKafkaTopicCurrentOffset() *datamapping.DataMapping {
	dataMapping := datamapping.NewDataMapping(schemas.Metric, "kafka", "topic", schemas.CurrentOffset)
	dataMapping.AddColumn("name", datamapping.Prometheus, "")
	dataMapping.AddColumn("namespace", datamapping.Prometheus, "")
	dataMapping.AddColumn("cluster_name", datamapping.Prometheus, "")
	dataMapping.AddColumn("value", datamapping.Prometheus, "")
	return dataMapping
}

func MetricKafkaCGCurrentOffset() *datamapping.DataMapping {
	dataMapping := datamapping.NewDataMapping(schemas.Metric, "kafka", "consumer_group", schemas.CurrentOffset)
	dataMapping.AddColumn("name", datamapping.Prometheus, "")
	dataMapping.AddColumn("namespace", datamapping.Prometheus, "")
	dataMapping.AddColumn("cluster_name", datamapping.Prometheus, "")
	dataMapping.AddColumn("topic_name", datamapping.Prometheus, "")
	dataMapping.AddColumn("value", datamapping.Prometheus, "")
	return dataMapping
}

func MetricKafkaCGLag() *datamapping.DataMapping {
	dataMapping := datamapping.NewDataMapping(schemas.Metric, "kafka", "consumer_group", schemas.Lag)
	dataMapping.AddColumn("name", datamapping.Prometheus, "")
	dataMapping.AddColumn("namespace", datamapping.Prometheus, "")
	dataMapping.AddColumn("cluster_name", datamapping.Prometheus, "")
	dataMapping.AddColumn("topic_name", datamapping.Prometheus, "")
	dataMapping.AddColumn("value", datamapping.Prometheus, "")
	return dataMapping
}
