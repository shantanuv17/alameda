package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaResourceApplication() *schemas.Schema {
	// cluster-status application
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "application")
	measurement := schemas.NewMeasurement("application", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("scaling_tool", true, schemas.Tag, common.String)
	measurement.AddColumn("machineset_names", false, schemas.Field, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "cluster")
	measurement := schemas.NewMeasurement("cluster", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceClusterScaler() *schemas.Schema {
	// cluster-status clusterscaler
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "clusterscaler")
	measurement := schemas.NewMeasurement("clusterscaler", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("foresee_time", true, schemas.Field, common.Int32)
	measurement.AddColumn("scale_down_idle_time", true, schemas.Field, common.Int32)
	measurement.AddColumn("max_nodes", true, schemas.Field, common.Int32)
	measurement.AddColumn("max_cpu", true, schemas.Field, common.Float64)
	measurement.AddColumn("max_mem", true, schemas.Field, common.Float64)
	measurement.AddColumn("enable_execution", true, schemas.Field, common.Bool)
	measurement.AddColumn("cpu_utilization_target", true, schemas.Field, common.Float64)
	measurement.AddColumn("cpu_scale_up_gap", true, schemas.Field, common.Float64)
	measurement.AddColumn("cpu_scale_down_gap", true, schemas.Field, common.Float64)
	measurement.AddColumn("memory_utilization_target", true, schemas.Field, common.Float64)
	measurement.AddColumn("memory_scale_up_gap", true, schemas.Field, common.Float64)
	measurement.AddColumn("memory_scale_down_gap", true, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "container")
	measurement := schemas.NewMeasurement("container", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("node_name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("pod_name", true, schemas.Tag, common.String)
	measurement.AddColumn("top_controller_name", true, schemas.Tag, common.String)
	measurement.AddColumn("top_controller_kind", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_scaling_tool", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_request_cpu", false, schemas.Field, common.String)
	measurement.AddColumn("resource_request_memory", false, schemas.Field, common.String)
	measurement.AddColumn("resource_limit_cpu", false, schemas.Field, common.String)
	measurement.AddColumn("resource_limit_memory", false, schemas.Field, common.String)
	measurement.AddColumn("status_waiting_reason", false, schemas.Field, common.String)
	measurement.AddColumn("status_waiting_message", false, schemas.Field, common.String)
	measurement.AddColumn("status_running_start_at", false, schemas.Field, common.Int64)
	measurement.AddColumn("status_terminated_exit_code", false, schemas.Field, common.Int32)
	measurement.AddColumn("status_terminated_reason", false, schemas.Field, common.String)
	measurement.AddColumn("status_terminated_message", false, schemas.Field, common.String)
	measurement.AddColumn("status_terminated_started_at", false, schemas.Field, common.Int64)
	measurement.AddColumn("status_terminated_finished_at", false, schemas.Field, common.Int64)
	measurement.AddColumn("last_termination_status_waiting_reason", false, schemas.Field, common.String)
	measurement.AddColumn("last_termination_status_waiting_message", false, schemas.Field, common.String)
	measurement.AddColumn("last_termination_status_running_start_at", false, schemas.Field, common.Int64)
	measurement.AddColumn("last_termination_status_terminated_exit_code", false, schemas.Field, common.Int32)
	measurement.AddColumn("last_termination_status_terminated_reason", false, schemas.Field, common.String)
	measurement.AddColumn("last_termination_status_terminated_message", false, schemas.Field, common.String)
	measurement.AddColumn("last_termination_status_terminated_started_at", false, schemas.Field, common.Int64)
	measurement.AddColumn("last_termination_status_terminated_finished_at", false, schemas.Field, common.Int64)
	measurement.AddColumn("restart_count", false, schemas.Field, common.Int32)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "controller")
	measurement := schemas.NewMeasurement("controller", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_scaling_tool", true, schemas.Tag, common.String)
	measurement.AddColumn("replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("spec_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("policy", false, schemas.Field, common.String)
	measurement.AddColumn("enable_execution", false, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceMachineScaler() *schemas.Schema {
	// cluster-status machinescaler
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "machinescaler")
	measurement := schemas.NewMeasurement("machinescaler", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("clusterscaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("min_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("max_replicas", true, schemas.Field, common.Int32)
	measurement.AddColumn("enable_execution", true, schemas.Field, common.Bool)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaMachineset() *schemas.Schema {
	// cluster-status machineset
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "machineset")
	measurement := schemas.NewMeasurement("machineset", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("machinescaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_k8s_spec_replicas", false, schemas.Field, common.Int32)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceNamespace() *schemas.Schema {
	// cluster-status namespace
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "namespace")
	measurement := schemas.NewMeasurement("namespace", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourceNode() *schemas.Schema {
	// Resource: cluster-status node
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "node")
	measurement := schemas.NewMeasurement("node", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("machineset_name", true, schemas.Tag, common.String)
	measurement.AddColumn("machineset_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("role_master", false, schemas.Field, common.Bool)
	measurement.AddColumn("role_worker", false, schemas.Field, common.Bool)
	measurement.AddColumn("role_infra", false, schemas.Field, common.Bool)
	measurement.AddColumn("create_time", false, schemas.Field, common.Int64)
	measurement.AddColumn("node_cpu_cores", false, schemas.Field, common.Int64)    // NodeCPUCores is the amount of cores in node
	measurement.AddColumn("node_memory_bytes", false, schemas.Field, common.Int64) // NodeMemoryBytes is the amount of memory bytes in node
	measurement.AddColumn("node_network_mbps", false, schemas.Field, common.Int64) // NodeNetworkMbps is mega bits per second
	measurement.AddColumn("io_provider", false, schemas.Field, common.String)      // Cloud service provider
	measurement.AddColumn("io_instance_type", false, schemas.Field, common.String)
	measurement.AddColumn("io_region", false, schemas.Field, common.String)
	measurement.AddColumn("io_zone", false, schemas.Field, common.String)
	measurement.AddColumn("io_os", false, schemas.Field, common.String)
	measurement.AddColumn("io_role", false, schemas.Field, common.String)
	measurement.AddColumn("io_instance_id", false, schemas.Field, common.String)
	measurement.AddColumn("io_storage_size", false, schemas.Field, common.Int64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaResourcePod() *schemas.Schema {
	// cluster-status pod
	schema := schemas.NewSchema(schemas.Resource, "cluster_status", "pod")
	measurement := schemas.NewMeasurement("pod", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("node_name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("uid", true, schemas.Tag, common.String)
	measurement.AddColumn("top_controller_name", true, schemas.Tag, common.String)
	measurement.AddColumn("top_controller_kind", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_scaling_tool", true, schemas.Tag, common.String)
	measurement.AddColumn("app_name", true, schemas.Tag, common.String)
	measurement.AddColumn("app_part_of", true, schemas.Tag, common.String)
	measurement.AddColumn("pod_create_time", false, schemas.Field, common.Int64)
	measurement.AddColumn("resource_link", false, schemas.Field, common.String)
	measurement.AddColumn("top_controller_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("pod_phase", false, schemas.Field, common.String)
	measurement.AddColumn("pod_message", false, schemas.Field, common.String)
	measurement.AddColumn("pod_reason", false, schemas.Field, common.String)
	measurement.AddColumn("policy", false, schemas.Field, common.String)
	measurement.AddColumn("used_recommendation_id", false, schemas.Field, common.String)
	measurement.AddColumn("alameda_scaler_resource_limit_cpu", false, schemas.Field, common.String)
	measurement.AddColumn("alameda_scaler_resource_limit_memory", false, schemas.Field, common.String)
	measurement.AddColumn("alameda_scaler_resource_request_cpu", false, schemas.Field, common.String)
	measurement.AddColumn("alameda_scaler_resource_request_memory", false, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
