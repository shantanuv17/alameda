package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaMetricKafkaTopic() *schemas.Schema {
	// Kafka topic
	schema := schemas.NewSchema(schemas.Metric, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka_topic_partition_current_offset", schemas.CurrentOffset, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaMetricKafkaCG() *schemas.Schema {
	// Kafka consumer group
	schema := schemas.NewSchema(schemas.Metric, "kafka", "consumer_group")

	// Current offset
	currentOffset := schemas.NewMeasurement("kafka_consumer_group_current_offset", schemas.CurrentOffset, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	currentOffset.AddColumn("name", true, schemas.Tag, common.String)
	currentOffset.AddColumn("namespace", true, schemas.Tag, common.String)
	currentOffset.AddColumn("cluster_name", true, schemas.Tag, common.String)
	currentOffset.AddColumn("topic_name", true, schemas.Tag, common.String)
	currentOffset.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, currentOffset)

	// Lag
	lag := schemas.NewMeasurement("kafka_consumer_group_lag", schemas.Lag, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	lag.AddColumn("name", true, schemas.Tag, common.String)
	lag.AddColumn("namespace", true, schemas.Tag, common.String)
	lag.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lag.AddColumn("topic_name", true, schemas.Tag, common.String)
	lag.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lag)

	return schema
}

func SchemaMetricResourceApplication() *schemas.Schema {
	// cluster-status application
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "application")

	// CPU
	cpu := schemas.NewMeasurement("application_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory
	memory := schemas.NewMeasurement("application_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("namespace", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricResourceCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "cluster")

	// CPU
	cpu := schemas.NewMeasurement("cluster_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory
	memory := schemas.NewMeasurement("cluster_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricResourceContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "container")

	// CPU
	cpu := schemas.NewMeasurement("container_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("node_name", true, schemas.Tag, common.String)
	cpu.AddColumn("rate_range", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory
	memory := schemas.NewMeasurement("container_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("pod_name", true, schemas.Tag, common.String)
	memory.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	memory.AddColumn("node_name", true, schemas.Tag, common.String)
	memory.AddColumn("rate_range", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricResourceController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "controller")

	// CPU
	cpu := schemas.NewMeasurement("controller_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("kind", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory
	memory := schemas.NewMeasurement("controller_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("namespace", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("kind", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricResourceNamespace() *schemas.Schema {
	// cluster-status namespace
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "namespace")

	// CPU
	cpu := schemas.NewMeasurement("namespace_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory
	memory := schemas.NewMeasurement("namespace_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricResourceNode() *schemas.Schema {
	// cluster-status node
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "node")

	// CPU
	cpu := schemas.NewMeasurement("node_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory
	memory := schemas.NewMeasurement("node_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}
