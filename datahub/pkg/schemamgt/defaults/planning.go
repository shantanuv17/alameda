package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaPlanningResourceApplication() *schemas.Schema {
	// cluster-status application
	schema := schemas.NewSchema(schemas.Planning, "cluster_status", "application")
	measurement := schemas.NewMeasurement("application", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)

	measurement.AddColumn("planning_id", true, schemas.Tag, common.String)
	measurement.AddColumn("planning_type", true, schemas.Tag, common.String)
	measurement.AddColumn("time", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("granularity", true, schemas.Tag, common.String)

	measurement.AddColumn("resource_request_cpu", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_request_memory", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_limit_cpu", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_limit_memory", false, schemas.Field, common.Int32)
	measurement.AddColumn("initial_resource_request_cpu", false, schemas.Field, common.Int32)
	measurement.AddColumn("initial_resource_request_memory", false, schemas.Field, common.Int32)
	measurement.AddColumn("initial_resource_limit_cpu", false, schemas.Field, common.Int32)
	measurement.AddColumn("initial_resource_limit_memory", false, schemas.Field, common.Int32)
	measurement.AddColumn("start_time", false, schemas.Field, common.Int32)
	measurement.AddColumn("end_time", false, schemas.Field, common.Int32)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Int32)
	measurement.AddColumn("apply_planning_now", false, schemas.Field, common.Int32)

	schema.Measurements = append(schema.Measurements, measurement)

	return schema
}

func SchemaPlanningResourceCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Planning, "cluster_status", "cluster")
	measurement := schemas.NewMeasurement("cluster", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)

	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaPlanningResourceContainer() *schemas.Schema {
	// cluster-status container
	schema := schemas.NewSchema(schemas.Planning, "cluster_status", "container")
	measurement := schemas.NewMeasurement("container", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)

	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaPlanningResourceController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Planning, "cluster_status", "controller")
	measurement := schemas.NewMeasurement("controller", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)

	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaPlanningResourceNamespace() *schemas.Schema {
	// cluster-status namespace
	schema := schemas.NewSchema(schemas.Planning, "cluster_status", "namespace")
	measurement := schemas.NewMeasurement("namespace", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)

	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaPlanningResourceNode() *schemas.Schema {
	// cluster-status node
	schema := schemas.NewSchema(schemas.Planning, "cluster_status", "node")
	measurement := schemas.NewMeasurement("node", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)

	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
