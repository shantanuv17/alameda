package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaPredictionKafkaTopic() *schemas.Schema {
	// Kafka topic
	schema := schemas.NewSchema(schemas.Prediction, "kafka", "topic")

	// Raw
	raw := schemas.NewMeasurement("kafka_topic_partition_current_offset", schemas.CurrentOffset, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	raw.AddColumn("name", true, schemas.Tag, common.String)
	raw.AddColumn("namespace", true, schemas.Tag, common.String)
	raw.AddColumn("cluster_name", true, schemas.Tag, common.String)
	raw.AddColumn("granularity", true, schemas.Tag, common.String)
	raw.AddColumn("model_id", true, schemas.Field, common.String)
	raw.AddColumn("prediction_id", true, schemas.Field, common.String)
	raw.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, raw)

	// Upper bound
	upperBound := schemas.NewMeasurement("kafka_topic_partition_current_offset_upper_bound", schemas.CurrentOffset, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperBound.AddColumn("name", true, schemas.Tag, common.String)
	upperBound.AddColumn("namespace", true, schemas.Tag, common.String)
	upperBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperBound.AddColumn("granularity", true, schemas.Tag, common.String)
	upperBound.AddColumn("model_id", true, schemas.Field, common.String)
	upperBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperBound.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperBound)

	// Lower bound
	lowerBound := schemas.NewMeasurement("kafka_topic_partition_current_offset_lower_bound", schemas.CurrentOffset, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerBound.AddColumn("name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerBound.AddColumn("model_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerBound)

	return schema
}

func SchemaPredictionKafkaCG() *schemas.Schema {
	// Kafka consumer group
	schema := schemas.NewSchema(schemas.Prediction, "kafka", "consumer_group")

	// Raw
	raw := schemas.NewMeasurement("kafka_consumer_group_current_offset", schemas.CurrentOffset, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	raw.AddColumn("name", true, schemas.Tag, common.String)
	raw.AddColumn("namespace", true, schemas.Tag, common.String)
	raw.AddColumn("cluster_name", true, schemas.Tag, common.String)
	raw.AddColumn("topic_name", true, schemas.Tag, common.String)
	raw.AddColumn("granularity", true, schemas.Tag, common.String)
	raw.AddColumn("model_id", true, schemas.Field, common.String)
	raw.AddColumn("prediction_id", true, schemas.Field, common.String)
	raw.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, raw)

	// Upper bound
	upperBound := schemas.NewMeasurement("kafka_consumer_group_current_offset_upper_bound", schemas.CurrentOffset, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperBound.AddColumn("name", true, schemas.Tag, common.String)
	upperBound.AddColumn("namespace", true, schemas.Tag, common.String)
	upperBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperBound.AddColumn("topic_name", true, schemas.Tag, common.String)
	upperBound.AddColumn("granularity", true, schemas.Tag, common.String)
	upperBound.AddColumn("model_id", true, schemas.Field, common.String)
	upperBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperBound.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperBound)

	// Lower bound
	lowerBound := schemas.NewMeasurement("kafka_consumer_group_current_offset_lower_bound", schemas.CurrentOffset, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerBound.AddColumn("name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("topic_name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerBound.AddColumn("model_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerBound)

	return schema
}

func SchemaPredictionNginx() *schemas.Schema {
	// Nginx
	schema := schemas.NewSchema(schemas.Prediction, "nginx", "nginx")

	// Raw response total
	rawTotal := schemas.NewMeasurement("nginx_http_response_total", schemas.Number, schemas.ResourceRaw, schemas.ResourceQuotaUndefined)
	rawTotal.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawTotal.AddColumn("resource_k8s_service_name", true, schemas.Tag, common.String)
	rawTotal.AddColumn("resource_k8s_service_namespace", true, schemas.Tag, common.String)
	rawTotal.AddColumn("resource_k8s_name", true, schemas.Tag, common.String)
	rawTotal.AddColumn("resource_k8s_namespace", true, schemas.Tag, common.String)
	rawTotal.AddColumn("resource_k8s_kind", true, schemas.Tag, common.String)
	rawTotal.AddColumn("granularity", true, schemas.Tag, common.String)
	rawTotal.AddColumn("model_id", true, schemas.Field, common.String)
	rawTotal.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawTotal.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawTotal)

	// Upper bound response total
	upperTotal := schemas.NewMeasurement("nginx_http_response_total_upper_bound", schemas.Number, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined)
	upperTotal.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperTotal.AddColumn("resource_k8s_service_name", true, schemas.Tag, common.String)
	upperTotal.AddColumn("resource_k8s_service_namespace", true, schemas.Tag, common.String)
	upperTotal.AddColumn("resource_k8s_name", true, schemas.Tag, common.String)
	upperTotal.AddColumn("resource_k8s_namespace", true, schemas.Tag, common.String)
	upperTotal.AddColumn("resource_k8s_kind", true, schemas.Tag, common.String)
	upperTotal.AddColumn("granularity", true, schemas.Tag, common.String)
	upperTotal.AddColumn("model_id", true, schemas.Field, common.String)
	upperTotal.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperTotal.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperTotal)

	// Lower bound response total
	lowerTotal := schemas.NewMeasurement("nginx_http_response_total_lower_bound", schemas.Number, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined)
	lowerTotal.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("resource_k8s_service_name", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("resource_k8s_service_namespace", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("resource_k8s_name", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("resource_k8s_namespace", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("resource_k8s_kind", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerTotal.AddColumn("model_id", true, schemas.Field, common.String)
	lowerTotal.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerTotal.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerTotal)

	return schema
}

func SchemaPredictionClusterAutoscalerMachinegroup() *schemas.Schema {
	// cluster-autoscaler machinegroup
	schema := schemas.NewSchema(schemas.Prediction, "cluster_autoscaler", "machinegroup")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("machinegroup_cpu", schemas.CPUUsageSecondsPercentage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	rawCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("machinegroup_memory", schemas.MemoryUsageBytes, schemas.ResourceRaw, schemas.ResourceQuotaUndefined)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("namespace", true, schemas.Tag, common.String)
	rawMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("machinegroup_cpu_upper_bound", schemas.CPUUsageSecondsPercentage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	upperCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound CPU
	upperMem := schemas.NewMeasurement("machinegroup_memory_upper_bound", schemas.MemoryUsageBytes, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("namespace", true, schemas.Tag, common.String)
	upperMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("machinegroup_cpu_lower_bound", schemas.CPUUsageSecondsPercentage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("machinegroup_memory_lower_bound", schemas.MemoryUsageBytes, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}

func SchemaPredictionClusterStatusApplication() *schemas.Schema {
	// cluster-status application
	schema := schemas.NewSchema(schemas.Prediction, "cluster_status", "application")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("application", schemas.CPUMilliCoresUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	rawCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("metric", false, schemas.Tag, common.String)
	rawCpu.AddColumn("kind", false, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("application", schemas.MemoryBytesUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("namespace", true, schemas.Tag, common.String)
	rawMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("metric", false, schemas.Tag, common.String)
	rawMem.AddColumn("kind", false, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("application", schemas.CPUMilliCoresUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	upperCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("metric", false, schemas.Tag, common.String)
	upperCpu.AddColumn("kind", false, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound memory
	upperMem := schemas.NewMeasurement("application", schemas.MemoryBytesUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("namespace", true, schemas.Tag, common.String)
	upperMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("metric", false, schemas.Tag, common.String)
	upperMem.AddColumn("kind", false, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("application", schemas.CPUMilliCoresUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("metric", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("kind", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("application", schemas.MemoryBytesUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("metric", false, schemas.Tag, common.String)
	lowerMem.AddColumn("kind", false, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}

func SchemaPredictionClusterStatusCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Prediction, "cluster_status", "cluster")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("cluster", schemas.CPUMilliCoresUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("metric", false, schemas.Tag, common.String)
	rawCpu.AddColumn("kind", false, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("cluster", schemas.MemoryBytesUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("metric", false, schemas.Tag, common.String)
	rawMem.AddColumn("kind", false, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("cluster", schemas.CPUMilliCoresUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("metric", false, schemas.Tag, common.String)
	upperCpu.AddColumn("kind", false, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound memory
	upperMem := schemas.NewMeasurement("cluster", schemas.MemoryBytesUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("metric", false, schemas.Tag, common.String)
	upperMem.AddColumn("kind", false, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("cluster", schemas.CPUMilliCoresUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("metric", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("kind", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("cluster", schemas.MemoryBytesUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("metric", false, schemas.Tag, common.String)
	lowerMem.AddColumn("kind", false, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}

func SchemaPredictionClusterStatusContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Prediction, "cluster_status", "container")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("container", schemas.CPUMilliCoresUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	rawCpu.AddColumn("node_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("metric", false, schemas.Tag, common.String)
	rawCpu.AddColumn("kind", false, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("container", schemas.MemoryBytesUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("pod_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("namespace", true, schemas.Tag, common.String)
	rawMem.AddColumn("node_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("metric", false, schemas.Tag, common.String)
	rawMem.AddColumn("kind", false, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("container", schemas.CPUMilliCoresUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	upperCpu.AddColumn("node_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("metric", false, schemas.Tag, common.String)
	upperCpu.AddColumn("kind", false, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound memory
	upperMem := schemas.NewMeasurement("container", schemas.MemoryBytesUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("pod_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("namespace", true, schemas.Tag, common.String)
	upperMem.AddColumn("node_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("metric", false, schemas.Tag, common.String)
	upperMem.AddColumn("kind", false, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("container", schemas.CPUMilliCoresUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("pod_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("node_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("metric", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("kind", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("container", schemas.MemoryBytesUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("pod_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerMem.AddColumn("node_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("metric", false, schemas.Tag, common.String)
	lowerMem.AddColumn("kind", false, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}

func SchemaPredictionClusterStatusController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Prediction, "cluster_status", "controller")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("controller", schemas.CPUMilliCoresUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	rawCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("metric", false, schemas.Tag, common.String)
	rawCpu.AddColumn("kind", false, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("controller_kind", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("controller", schemas.MemoryBytesUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("namespace", true, schemas.Tag, common.String)
	rawMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("metric", false, schemas.Tag, common.String)
	rawMem.AddColumn("kind", false, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("controller_kind", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("controller", schemas.CPUMilliCoresUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	upperCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("metric", false, schemas.Tag, common.String)
	upperCpu.AddColumn("kind", false, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("controller_kind", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound memory
	upperMem := schemas.NewMeasurement("controller", schemas.MemoryBytesUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("namespace", true, schemas.Tag, common.String)
	upperMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("metric", false, schemas.Tag, common.String)
	upperMem.AddColumn("kind", false, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("controller_kind", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("controller", schemas.CPUMilliCoresUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("metric", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("kind", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("controller_kind", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("controller", schemas.MemoryBytesUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("metric", false, schemas.Tag, common.String)
	lowerMem.AddColumn("kind", false, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("controller_kind", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}

func SchemaPredictionClusterStatusNamespace() *schemas.Schema {
	// cluster-status namespace
	schema := schemas.NewSchema(schemas.Prediction, "cluster_status", "namespace")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("namespace", schemas.CPUMilliCoresUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("metric", false, schemas.Tag, common.String)
	rawCpu.AddColumn("kind", false, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("namespace", schemas.MemoryBytesUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("metric", false, schemas.Tag, common.String)
	rawMem.AddColumn("kind", false, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("namespace", schemas.CPUMilliCoresUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("metric", false, schemas.Tag, common.String)
	upperCpu.AddColumn("kind", false, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound memory
	upperMem := schemas.NewMeasurement("namespace", schemas.MemoryBytesUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("metric", false, schemas.Tag, common.String)
	upperMem.AddColumn("kind", false, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("namespace", schemas.CPUMilliCoresUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("metric", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("kind", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("namespace", schemas.MemoryBytesUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("metric", false, schemas.Tag, common.String)
	lowerMem.AddColumn("kind", false, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}

func SchemaPredictionClusterStatusNode() *schemas.Schema {
	// cluster-status node
	schema := schemas.NewSchema(schemas.Prediction, "cluster_status", "node")

	// Raw CPU
	rawCpu := schemas.NewMeasurement("node", schemas.CPUMilliCoresUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawCpu.AddColumn("name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawCpu.AddColumn("metric", false, schemas.Tag, common.String)
	rawCpu.AddColumn("kind", false, schemas.Tag, common.String)
	rawCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	rawCpu.AddColumn("is_scheduled", true, schemas.Tag, common.String)
	rawCpu.AddColumn("model_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawCpu)

	// Raw memory
	rawMem := schemas.NewMeasurement("node", schemas.MemoryBytesUsage, schemas.ResourceRaw, schemas.ResourceQuotaUndefined, true)
	rawMem.AddColumn("name", true, schemas.Tag, common.String)
	rawMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	rawMem.AddColumn("metric", false, schemas.Tag, common.String)
	rawMem.AddColumn("kind", false, schemas.Tag, common.String)
	rawMem.AddColumn("granularity", true, schemas.Tag, common.String)
	rawMem.AddColumn("is_scheduled", true, schemas.Tag, common.String)
	rawMem.AddColumn("model_id", true, schemas.Field, common.String)
	rawMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	rawMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, rawMem)

	// Upper bound CPU
	upperCpu := schemas.NewMeasurement("node", schemas.CPUMilliCoresUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperCpu.AddColumn("name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperCpu.AddColumn("metric", false, schemas.Tag, common.String)
	upperCpu.AddColumn("kind", false, schemas.Tag, common.String)
	upperCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	upperCpu.AddColumn("is_scheduled", true, schemas.Tag, common.String)
	upperCpu.AddColumn("model_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperCpu)

	// Upper bound memory
	upperMem := schemas.NewMeasurement("node", schemas.MemoryBytesUsage, schemas.ResourceUpperBound, schemas.ResourceQuotaUndefined, true)
	upperMem.AddColumn("name", true, schemas.Tag, common.String)
	upperMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperMem.AddColumn("metric", false, schemas.Tag, common.String)
	upperMem.AddColumn("kind", false, schemas.Tag, common.String)
	upperMem.AddColumn("granularity", true, schemas.Tag, common.String)
	upperMem.AddColumn("is_scheduled", true, schemas.Tag, common.String)
	upperMem.AddColumn("model_id", true, schemas.Field, common.String)
	upperMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, upperMem)

	// Lower bound CPU
	lowerCpu := schemas.NewMeasurement("node", schemas.CPUMilliCoresUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerCpu.AddColumn("name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("metric", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("kind", false, schemas.Tag, common.String)
	lowerCpu.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("is_scheduled", true, schemas.Tag, common.String)
	lowerCpu.AddColumn("model_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerCpu.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerCpu)

	// Lower bound memory
	lowerMem := schemas.NewMeasurement("node", schemas.MemoryBytesUsage, schemas.ResourceLowerBound, schemas.ResourceQuotaUndefined, true)
	lowerMem.AddColumn("name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerMem.AddColumn("metric", false, schemas.Tag, common.String)
	lowerMem.AddColumn("kind", false, schemas.Tag, common.String)
	lowerMem.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerMem.AddColumn("is_scheduled", true, schemas.Tag, common.String)
	lowerMem.AddColumn("model_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerMem.AddColumn("value", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, lowerMem)

	return schema
}
