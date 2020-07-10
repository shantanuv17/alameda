package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaFedemeterCalculation() *schemas.Schema {
	// Fedemeter node cost
	schema := schemas.NewSchema(schemas.Fedemeter, "calculation", "price")

	// Instance cost
	instance := schemas.NewMeasurement("calculation_price_instance", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	instance.AddColumn("nodename", true, schemas.Tag, common.String)
	instance.AddColumn("clustername", true, schemas.Tag, common.String)
	instance.AddColumn("provider", true, schemas.Tag, common.String)
	instance.AddColumn("region", true, schemas.Tag, common.String)
	instance.AddColumn("unit", true, schemas.Tag, common.String)
	instance.AddColumn("granularity", true, schemas.Tag, common.String)
	instance.AddColumn("cpu", false, schemas.Field, common.Float64)
	instance.AddColumn("memory", false, schemas.Field, common.Float64)
	instance.AddColumn("description", false, schemas.Field, common.String)
	instance.AddColumn("displayname", false, schemas.Field, common.String)
	instance.AddColumn("instancenum", false, schemas.Field, common.Int32)
	instance.AddColumn("instancetype", false, schemas.Field, common.String)
	instance.AddColumn("nodetype", false, schemas.Field, common.String)
	instance.AddColumn("operatingsystem", false, schemas.Field, common.String)
	instance.AddColumn("period", false, schemas.Field, common.Int32)
	instance.AddColumn("preinstalledsw", false, schemas.Field, common.String)
	instance.AddColumn("starttime", false, schemas.Field, common.Int64)
	instance.AddColumn("cost", false, schemas.Field, common.Float64)
	instance.AddColumn("totalcost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, instance)

	// Storage cost
	storage := schemas.NewMeasurement("calculation_price_storage", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	storage.AddColumn("nodename", true, schemas.Tag, common.String)
	storage.AddColumn("clustername", true, schemas.Tag, common.String)
	storage.AddColumn("provider", true, schemas.Tag, common.String)
	storage.AddColumn("unit", true, schemas.Tag, common.String)
	storage.AddColumn("granularity", true, schemas.Tag, common.String)
	storage.AddColumn("description", false, schemas.Field, common.String)
	storage.AddColumn("displayname", false, schemas.Field, common.String)
	storage.AddColumn("period", false, schemas.Field, common.Int32)
	storage.AddColumn("storagenum", false, schemas.Field, common.Int32)
	storage.AddColumn("storagesize", false, schemas.Field, common.Int64)
	storage.AddColumn("volumetype", false, schemas.Field, common.String)
	storage.AddColumn("starttime", false, schemas.Field, common.Int64)
	storage.AddColumn("cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, storage)

	return schema
}

func SchemaFedemeterRecommendationJERI() *schemas.Schema {
	schema := schemas.NewSchema(schemas.Fedemeter, "recommendation", "jeri")
	measurement := schemas.NewMeasurement("recommendation_jeri", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("clustername", true, schemas.Tag, common.String)
	measurement.AddColumn("country", true, schemas.Tag, common.String)
	measurement.AddColumn("instancetype", true, schemas.Tag, common.String)
	measurement.AddColumn("provider", true, schemas.Tag, common.String)
	measurement.AddColumn("rank", true, schemas.Tag, common.String)
	measurement.AddColumn("region", true, schemas.Tag, common.String)
	measurement.AddColumn("reservedinstances", true, schemas.Tag, common.String)
	measurement.AddColumn("granularity", true, schemas.Tag, common.String)
	measurement.AddColumn("master_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("worker_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("master_storage_size", false, schemas.Field, common.Float64)
	measurement.AddColumn("worker_storage_size", false, schemas.Field, common.Float64)
	measurement.AddColumn("acc_cost", false, schemas.Field, common.Float64)
	measurement.AddColumn("display_name", false, schemas.Field, common.String)
	measurement.AddColumn("resource_name", false, schemas.Field, common.String)
	measurement.AddColumn("start_time", false, schemas.Field, common.Int64)
	measurement.AddColumn("ondemand_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("master_ri_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("worker_ri_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("master_spot_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("worker_spot_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("master_ondemand_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("worker_ondemand_num", false, schemas.Field, common.Int32)
	measurement.AddColumn("cost", false, schemas.Field, common.Float64)
	measurement.AddColumn("total_cost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaFedemeterResourceHistoryCost() *schemas.Schema {
	// Fedemeter resource cost allocation
	schema := schemas.NewSchema(schemas.Fedemeter, "resource_history", "cost")

	// Application cost allocation
	app := schemas.NewMeasurement("resource_history_cost_app", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	app.AddColumn("appname", true, schemas.Tag, common.String)
	app.AddColumn("namespacename", true, schemas.Tag, common.String)
	app.AddColumn("clustername", true, schemas.Tag, common.String)
	app.AddColumn("provider", true, schemas.Tag, common.String)
	app.AddColumn("type", true, schemas.Tag, common.String)
	app.AddColumn("granularity", true, schemas.Tag, common.String)
	app.AddColumn("createtime", false, schemas.Field, common.Int64)
	app.AddColumn("starttime", false, schemas.Field, common.Int64)
	app.AddColumn("costpercentage", false, schemas.Field, common.Float64)
	app.AddColumn("workloadcost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, app)

	// Namespace cost allocation
	namespace := schemas.NewMeasurement("resource_history_cost_namespace", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	namespace.AddColumn("namespacename", true, schemas.Tag, common.String)
	namespace.AddColumn("clustername", true, schemas.Tag, common.String)
	namespace.AddColumn("provider", true, schemas.Tag, common.String)
	namespace.AddColumn("type", true, schemas.Tag, common.String)
	namespace.AddColumn("granularity", true, schemas.Tag, common.String)
	namespace.AddColumn("createtime", false, schemas.Field, common.Int64)
	namespace.AddColumn("starttime", false, schemas.Field, common.Int64)
	namespace.AddColumn("costpercentage", false, schemas.Field, common.Float64)
	namespace.AddColumn("workloadcost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, namespace)

	return schema
}

func SchemaFedemeterResourcePredictionCost() *schemas.Schema {
	// Fedemeter resource cost allocation
	schema := schemas.NewSchema(schemas.Fedemeter, "resource_prediction", "cost")

	// Application cost allocation
	app := schemas.NewMeasurement("resource_prediction_cost_app", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	app.AddColumn("appname", true, schemas.Tag, common.String)
	app.AddColumn("namespacename", true, schemas.Tag, common.String)
	app.AddColumn("clustername", true, schemas.Tag, common.String)
	app.AddColumn("provider", true, schemas.Tag, common.String)
	app.AddColumn("type", true, schemas.Tag, common.String)
	app.AddColumn("granularity", true, schemas.Tag, common.String)
	app.AddColumn("createtime", false, schemas.Field, common.Int64)
	app.AddColumn("starttime", false, schemas.Field, common.Int64)
	app.AddColumn("costpercentage", false, schemas.Field, common.Float64)
	app.AddColumn("workloadcost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, app)

	// Namespace cost allocation
	namespace := schemas.NewMeasurement("resource_prediction_cost_namespace", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	namespace.AddColumn("namespacename", true, schemas.Tag, common.String)
	namespace.AddColumn("clustername", true, schemas.Tag, common.String)
	namespace.AddColumn("provider", true, schemas.Tag, common.String)
	namespace.AddColumn("type", true, schemas.Tag, common.String)
	namespace.AddColumn("granularity", true, schemas.Tag, common.String)
	namespace.AddColumn("createtime", false, schemas.Field, common.Int64)
	namespace.AddColumn("starttime", false, schemas.Field, common.Int64)
	namespace.AddColumn("costpercentage", false, schemas.Field, common.Float64)
	namespace.AddColumn("workloadcost", false, schemas.Field, common.Float64)
	schema.Measurements = append(schema.Measurements, namespace)

	return schema
}
