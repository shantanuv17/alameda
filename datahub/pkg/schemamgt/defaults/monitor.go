package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaMonitorKafka() *schemas.Schema {
	schema := schemas.NewSchema(schemas.Monitor, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("category", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("username", true, schemas.Tag, common.String)
	measurement.AddColumn("metric_type", true, schemas.Tag, common.String)
	measurement.AddColumn("agent_version", false, schemas.Field, common.String)
	measurement.AddColumn("is_monitor", true, schemas.Field, common.Bool)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaMonitorKafkaCG() *schemas.Schema {
	schema := schemas.NewSchema(schemas.Monitor, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("category", true, schemas.Tag, common.String)
	measurement.AddColumn("type", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("username", true, schemas.Tag, common.String)
	measurement.AddColumn("metric_type", true, schemas.Tag, common.String)
	measurement.AddColumn("agent_version", true, schemas.Field, common.String)
	measurement.AddColumn("is_monitor", true, schemas.Field, common.Bool)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
