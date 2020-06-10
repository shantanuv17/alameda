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

func SchemaMetricNginx() *schemas.Schema {
	// Nginx
	schema := schemas.NewSchema(schemas.Metric, "nginx", "nginx")

	// Response total
	responseTotal := schemas.NewMeasurement("nginx_http_response_total", schemas.Number, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	responseTotal.AddColumn("cluster_name", true, schemas.Tag, common.String)
	responseTotal.AddColumn("resource_k8s_service_name", true, schemas.Tag, common.String)
	responseTotal.AddColumn("resource_k8s_service_namespace", true, schemas.Tag, common.String)
	responseTotal.AddColumn("resource_k8s_name", true, schemas.Tag, common.String)
	responseTotal.AddColumn("resource_k8s_namespace", true, schemas.Tag, common.String)
	responseTotal.AddColumn("resource_k8s_kind", true, schemas.Tag, common.String)
	responseTotal.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, responseTotal)

	// Response latency
	latency := schemas.NewMeasurement("nginx_http_response_latency_ms", schemas.Latency, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	latency.AddColumn("cluster_name", true, schemas.Tag, common.String)
	latency.AddColumn("resource_k8s_service_name", true, schemas.Tag, common.String)
	latency.AddColumn("resource_k8s_service_namespace", true, schemas.Tag, common.String)
	latency.AddColumn("resource_k8s_name", true, schemas.Tag, common.String)
	latency.AddColumn("resource_k8s_namespace", true, schemas.Tag, common.String)
	latency.AddColumn("resource_k8s_kind", true, schemas.Tag, common.String)
	latency.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, latency)

	return schema
}

func SchemaMetricClusterStatusApplication() *schemas.Schema {
	// cluster-status application
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "application")

	// CPU usage percentage
	cpu := schemas.NewMeasurement("application_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory usage bytes
	memory := schemas.NewMeasurement("application_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("namespace", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricClusterStatusCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "cluster")

	// CPU usage percentage
	cpu := schemas.NewMeasurement("cluster_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory usage bytes
	memory := schemas.NewMeasurement("cluster_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricClusterStatusContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "container")

	// CPU usage percentage
	cpu := schemas.NewMeasurement("container_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("node_name", true, schemas.Tag, common.String)
	cpu.AddColumn("rate_range", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory usage bytes
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

func SchemaMetricClusterStatusController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "controller")

	// CPU usage percentage
	cpu := schemas.NewMeasurement("controller_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("kind", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory usage bytes
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

func SchemaMetricClusterStatusNamespace() *schemas.Schema {
	// cluster-status namespace
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "namespace")

	// CPU usage percentage
	cpu := schemas.NewMeasurement("namespace_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory usage bytes
	memory := schemas.NewMeasurement("namespace_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricClusterStatusNode() *schemas.Schema {
	// cluster-status node
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "node")

	// CPU usage percentage
	cpu := schemas.NewMeasurement("node_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory usage bytes
	memory := schemas.NewMeasurement("node_memory", schemas.MemoryUsageBytes, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	// CPU cores
	cores := schemas.NewMeasurement("node_cpu_cores", schemas.CPUCores, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	cores.AddColumn("name", true, schemas.Tag, common.String)
	cores.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cores.AddColumn("uid", true, schemas.Tag, common.String)
	cores.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cores)

	return schema
}
