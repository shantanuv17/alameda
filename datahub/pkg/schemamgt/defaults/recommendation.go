package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaRecommendationKafkaCG() *schemas.Schema {
	// Kafka consumer group
	schema := schemas.NewSchema(schemas.Recommendation, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka_consumer_group", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Tag, common.String)
	measurement.AddColumn("create_time", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("execution_time", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaRecommendationResourceApplication() *schemas.Schema {
	// cluster-status application
	schema := schemas.NewSchema(schemas.Recommendation, "cluster_status", "application")
	measurement := schemas.NewMeasurement("application", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("create_time", true, schemas.Field, common.Int64)
	measurement.AddColumn("current_cpu_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaRecommendationResourceCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Recommendation, "cluster_status", "cluster")
	measurement := schemas.NewMeasurement("cluster", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("create_time", true, schemas.Field, common.Int64)
	measurement.AddColumn("current_cpu_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaRecommendationResourceContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Recommendation, "cluster_status", "container")

	// Resource limit
	limit := schemas.NewMeasurement("container", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceLimit)
	limit.AddColumn("name", true, schemas.Tag, common.String)
	limit.AddColumn("namespace", true, schemas.Tag, common.String)
	limit.AddColumn("cluster_name", true, schemas.Tag, common.String)
	limit.AddColumn("pod_name", true, schemas.Tag, common.String)
	limit.AddColumn("granularity", true, schemas.Tag, common.String)
	limit.AddColumn("top_controller_name", true, schemas.Field, common.String)
	limit.AddColumn("top_controller_kind", true, schemas.Field, common.String)
	limit.AddColumn("policy", true, schemas.Field, common.String)
	limit.AddColumn("policy_time", true, schemas.Field, common.Int64)
	limit.AddColumn("pod_total_cost", true, schemas.Field, common.Float64)
	limit.AddColumn("start_time", true, schemas.Field, common.Int64)
	limit.AddColumn("end_time", true, schemas.Field, common.Int64)
	limit.AddColumn("resource_limit_cpu", true, schemas.Field, common.Float64)
	limit.AddColumn("resource_limit_memory", true, schemas.Field, common.Float64)
	limit.AddColumn("initial_resource_limit_cpu", true, schemas.Field, common.Float64)
	limit.AddColumn("initial_resource_limit_memory", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, limit)

	// Resource request
	request := schemas.NewMeasurement("container", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceRequest)
	request.AddColumn("name", true, schemas.Tag, common.String)
	request.AddColumn("namespace", true, schemas.Tag, common.String)
	request.AddColumn("cluster_name", true, schemas.Tag, common.String)
	request.AddColumn("pod_name", true, schemas.Tag, common.String)
	request.AddColumn("granularity", true, schemas.Tag, common.String)
	request.AddColumn("top_controller_name", true, schemas.Field, common.String)
	request.AddColumn("top_controller_kind", true, schemas.Field, common.String)
	request.AddColumn("policy", true, schemas.Field, common.String)
	request.AddColumn("policy_time", true, schemas.Field, common.Int64)
	request.AddColumn("pod_total_cost", true, schemas.Field, common.Float64)
	request.AddColumn("start_time", true, schemas.Field, common.Int64)
	request.AddColumn("end_time", true, schemas.Field, common.Int64)
	request.AddColumn("resource_request_cpu", true, schemas.Field, common.Float64)
	request.AddColumn("resource_request_memory", true, schemas.Field, common.Float64)
	request.AddColumn("initial_resource_request_cpu", true, schemas.Field, common.Float64)
	request.AddColumn("initial_resource_request_memory", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, request)

	return schema
}

func SchemaRecommendationResourceController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Recommendation, "cluster_status", "controller")
	measurement := schemas.NewMeasurement("controller", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("create_time", true, schemas.Field, common.Int64)
	measurement.AddColumn("current_cpu_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaRecommendationResourceNamespace() *schemas.Schema {
	// cluster-status namespace
	schema := schemas.NewSchema(schemas.Recommendation, "cluster_status", "namespace")
	measurement := schemas.NewMeasurement("namespace", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("create_time", true, schemas.Field, common.Int64)
	measurement.AddColumn("current_cpu_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaRecommendationResourceNode() *schemas.Schema {
	// cluster-status node
	schema := schemas.NewSchema(schemas.Recommendation, "cluster_status", "node")
	measurement := schemas.NewMeasurement("node", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("create_time", true, schemas.Field, common.Int64)
	measurement.AddColumn("current_cpu_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_requests", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("current_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_cpu_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("desired_mem_limits", false, schemas.Field, common.Float64)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
