package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaMetricKafkaTopic() *schemas.Schema {
	// Kafka topic
	schema := schemas.NewSchema(schemas.Metric, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka_topic_partition_current_offset", schemas.CurrentOffset, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
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
	currentOffset := schemas.NewMeasurement("kafka_consumer_group_current_offset", schemas.CurrentOffset, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	currentOffset.AddColumn("name", true, schemas.Tag, common.String)
	currentOffset.AddColumn("namespace", true, schemas.Tag, common.String)
	currentOffset.AddColumn("cluster_name", true, schemas.Tag, common.String)
	currentOffset.AddColumn("topic_name", true, schemas.Tag, common.String)
	currentOffset.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, currentOffset)

	// Lag
	lag := schemas.NewMeasurement("kafka_consumer_group_lag", schemas.Lag, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
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

	// CPU millicores usage
	cpu := schemas.NewMeasurement("application_cpu", schemas.CPUMilliCoresUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes usage
	memory := schemas.NewMeasurement("application_memory", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
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

	// CPU millicores usage
	cpu := schemas.NewMeasurement("cluster_cpu", schemas.CPUMilliCoresUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes usage
	memory := schemas.NewMeasurement("cluster_memory", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}

func SchemaMetricClusterStatusContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "container")

	// CPU millicores usage
	cpu := schemas.NewMeasurement("container_cpu", schemas.CPUMilliCoresUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("node_name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("rate_range", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes usage
	memory := schemas.NewMeasurement("container_memory", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("node_name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("pod_name", true, schemas.Tag, common.String)
	memory.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	memory.AddColumn("rate_range", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	// Resource requests cpu millicores
	requestCpu := schemas.NewMeasurement("container_resource_requests_cpu_millicores", schemas.CPUMilliCoresTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceRequest, true)
	requestCpu.AddColumn("name", true, schemas.Tag, common.String)
	requestCpu.AddColumn("node_name", true, schemas.Tag, common.String)
	requestCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	requestCpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	requestCpu.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	requestCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, requestCpu)

	// Resource limits cpu millicores
	limitCpu := schemas.NewMeasurement("container_resource_limits_cpu_millicores", schemas.CPUMilliCoresTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceLimit, true)
	limitCpu.AddColumn("name", true, schemas.Tag, common.String)
	limitCpu.AddColumn("node_name", true, schemas.Tag, common.String)
	limitCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	limitCpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	limitCpu.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	limitCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, limitCpu)

	// Resource requests memory bytes
	requestMemory := schemas.NewMeasurement("container_resource_requests_memory_bytes", schemas.MemoryBytesTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceRequest, true)
	requestMemory.AddColumn("name", true, schemas.Tag, common.String)
	requestMemory.AddColumn("node_name", true, schemas.Tag, common.String)
	requestMemory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	requestMemory.AddColumn("pod_name", true, schemas.Tag, common.String)
	requestMemory.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	requestMemory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, requestMemory)

	// Resource limits memory bytes
	limitMemory := schemas.NewMeasurement("container_resource_limits_memory_bytes", schemas.MemoryBytesTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceLimit, true)
	limitMemory.AddColumn("name", true, schemas.Tag, common.String)
	limitMemory.AddColumn("node_name", true, schemas.Tag, common.String)
	limitMemory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	limitMemory.AddColumn("pod_name", true, schemas.Tag, common.String)
	limitMemory.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	limitMemory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, limitMemory)

	// Restarts total
	restarts := schemas.NewMeasurement("container_restarts_total", schemas.RestartsTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	restarts.AddColumn("name", true, schemas.Tag, common.String)
	restarts.AddColumn("node_name", true, schemas.Tag, common.String)
	restarts.AddColumn("cluster_name", true, schemas.Tag, common.String)
	restarts.AddColumn("pod_name", true, schemas.Tag, common.String)
	restarts.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	restarts.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, restarts)

	return schema
}

func SchemaMetricClusterStatusController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "controller")

	// CPU millicores usage
	cpu := schemas.NewMeasurement("controller_cpu", schemas.CPUMilliCoresUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("kind", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes usage
	memory := schemas.NewMeasurement("controller_memory", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
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

	// CPU millicores usage
	cpu := schemas.NewMeasurement("namespace_cpu", schemas.CPUMilliCoresUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes usage
	memory := schemas.NewMeasurement("namespace_memory", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
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

	// CPU cores allocatable
	cores := schemas.NewMeasurement("node_cpu_cores_allocatable", schemas.CPUCoresAllocatable, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cores.AddColumn("name", true, schemas.Tag, common.String)
	cores.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cores.AddColumn("uid", true, schemas.Tag, common.String)
	cores.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cores)

	// CPU millicores total
	cpuTotal := schemas.NewMeasurement("node_cpu_millicores_total", schemas.CPUMilliCoresTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpuTotal.AddColumn("name", true, schemas.Tag, common.String)
	cpuTotal.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpuTotal.AddColumn("uid", true, schemas.Tag, common.String)
	cpuTotal.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpuTotal)

	// CPU millicores usage
	cpu := schemas.NewMeasurement("node_cpu", schemas.CPUMilliCoresUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("uid", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes total
	memoryTotal := schemas.NewMeasurement("node_memory_bytes_total", schemas.MemoryBytesTotal, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	memoryTotal.AddColumn("name", true, schemas.Tag, common.String)
	memoryTotal.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memoryTotal.AddColumn("uid", true, schemas.Tag, common.String)
	memoryTotal.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memoryTotal)

	// Memory bytes usage
	memory := schemas.NewMeasurement("node_memory", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("uid", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	// Filesystem bytes usage percentage
	fsPCT := schemas.NewMeasurement("node_fs_bytes_usage_pct", schemas.FSBytesUsagePCT, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	fsPCT.AddColumn("name", true, schemas.Tag, common.String)
	fsPCT.AddColumn("cluster_name", true, schemas.Tag, common.String)
	fsPCT.AddColumn("uid", true, schemas.Tag, common.String)
	fsPCT.AddColumn("device", false, schemas.Field, common.String)
	fsPCT.AddColumn("fs_type", false, schemas.Field, common.String)
	fsPCT.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, fsPCT)

	// Disk I/O utilization
	diskIO := schemas.NewMeasurement("node_disk_io_util", schemas.DiskIOUtilization, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	diskIO.AddColumn("name", true, schemas.Tag, common.String)
	diskIO.AddColumn("cluster_name", true, schemas.Tag, common.String)
	diskIO.AddColumn("uid", true, schemas.Tag, common.String)
	diskIO.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, diskIO)

	// Pod phase count
	podPhase := schemas.NewMeasurement("node_pod_phase_count", schemas.Number, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	podPhase.AddColumn("name", true, schemas.Tag, common.String)
	podPhase.AddColumn("cluster_name", true, schemas.Tag, common.String)
	podPhase.AddColumn("uid", true, schemas.Tag, common.String)
	podPhase.AddColumn("pending", true, schemas.Field, common.Int64)
	podPhase.AddColumn("running", true, schemas.Field, common.Int64)
	podPhase.AddColumn("succeeded", true, schemas.Field, common.Int64)
	podPhase.AddColumn("failed", true, schemas.Field, common.Int64)
	podPhase.AddColumn("unknown", true, schemas.Field, common.Int64)
	schema.Measurements = append(schema.Measurements, podPhase)

	// Unschedulable
	unschedulable := schemas.NewMeasurement("node_unschedulable", schemas.Unschedulable, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	unschedulable.AddColumn("name", true, schemas.Tag, common.String)
	unschedulable.AddColumn("cluster_name", true, schemas.Tag, common.String)
	unschedulable.AddColumn("uid", true, schemas.Tag, common.String)
	unschedulable.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, unschedulable)

	return schema
}

func SchemaMetricClusterStatusService() *schemas.Schema {
	// cluster-status service
	schema := schemas.NewSchema(schemas.Metric, "cluster_status", "service")
	measurement := schemas.NewMeasurement("service_health", schemas.Health, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("node_name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("pod_name", true, schemas.Tag, common.String)
	measurement.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("job", true, schemas.Tag, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.Int32) // 1: healthy, 0: failed
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaMetricTopContainer() *schemas.Schema {
	// top container
	schema := schemas.NewSchema(schemas.Metric, "top", "container")

	// CPU millicores usage percentage
	cpu := schemas.NewMeasurement("top_container_cpu_millicores_usage_pct", schemas.CPUMilliCoresUsagePCT, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	cpu.AddColumn("name", true, schemas.Tag, common.String)
	cpu.AddColumn("node_name", true, schemas.Tag, common.String)
	cpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	cpu.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	cpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, cpu)

	// Memory bytes usage
	memory := schemas.NewMeasurement("top_container_memory_bytes_usage", schemas.MemoryBytesUsage, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, true)
	memory.AddColumn("name", true, schemas.Tag, common.String)
	memory.AddColumn("node_name", true, schemas.Tag, common.String)
	memory.AddColumn("cluster_name", true, schemas.Tag, common.String)
	memory.AddColumn("pod_name", true, schemas.Tag, common.String)
	memory.AddColumn("pod_namespace", true, schemas.Tag, common.String)
	memory.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, memory)

	return schema
}
