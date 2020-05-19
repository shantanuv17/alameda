package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaExecutionClusterAutoscalerMachineset() *schemas.Schema {
	// cluster-autoscaler machineset
	schema := schemas.NewSchema(schemas.Execution, "cluster_autoscaler", "machineset")
	measurement := schemas.NewMeasurement("machineset", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("machinegroup_name", true, schemas.Tag, common.String)
	measurement.AddColumn("execution_time", true, schemas.Field, common.String)
	measurement.AddColumn("replicas_from", true, schemas.Field, common.Int32)
	measurement.AddColumn("replicas_to", true, schemas.Field, common.Int32)
	measurement.AddColumn("delta_up_time​", true, schemas.Field, common.Int64)
	measurement.AddColumn("delta_down_time", true, schemas.Field, common.Int64)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
