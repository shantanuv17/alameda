package defaults

import (
	"github.com/containers-ai/alameda/pkg/database/common"
	"github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
)

func SchemaTargetCluster() *schemas.Schema {
	// cluster-status cluster
	schema := schemas.NewSchema(schemas.Target, "cluster_status", "cluster")
	measurement := schemas.NewMeasurement("cluster", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, false)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("raw_spec", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaTargetController() *schemas.Schema {
	// cluster-status controller
	schema := schemas.NewSchema(schemas.Target, "cluster_status", "controller")
	measurement := schemas.NewMeasurement("controller", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, false)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("kind", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_scaling_tool", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_min_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_k8s_max_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("policy", false, schemas.Field, common.String)
	measurement.AddColumn("enable_execution", false, schemas.Field, common.Bool)
	measurement.AddColumn("raw_spec", false, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaTargetKafkaTopic() *schemas.Schema {
	// Kafka topic
	schema := schemas.NewSchema(schemas.Target, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka_topic", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, false)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("exporter_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("raw_spec", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaTargetKafkaConsumerGroup() *schemas.Schema {
	// Kafka consumer group
	schema := schemas.NewSchema(schemas.Target, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka_consumer_group", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined, false)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("exporter_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("topic_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_scaling_tool", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_name", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_kind", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_min_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_k8s_max_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("policy", false, schemas.Field, common.String)
	measurement.AddColumn("enable_execution", false, schemas.Field, common.Bool)
	measurement.AddColumn("raw_spec", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
