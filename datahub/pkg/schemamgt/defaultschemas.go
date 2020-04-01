package schemamgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt/defaults"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func DefaultSchemasInit() {
	schemaMgt := NewSchemaManagement()
	schemaList := make([]*schemas.Schema, 0)

	// Application
	schemaList = append(schemaList, defaults.SchemaApplicationKafkaTopic())
	schemaList = append(schemaList, defaults.SchemaApplicationKafkaCG())

	// Metric
	schemaList = append(schemaList, defaults.SchemaMetricKafkaTopic())
	schemaList = append(schemaList, defaults.SchemaMetricKafkaCG())
	schemaList = append(schemaList, defaults.SchemaMetricResourceApplication())
	schemaList = append(schemaList, defaults.SchemaMetricResourceCluster())
	schemaList = append(schemaList, defaults.SchemaMetricResourceContainer())
	schemaList = append(schemaList, defaults.SchemaMetricResourceController())
	schemaList = append(schemaList, defaults.SchemaMetricResourceNamespace())
	schemaList = append(schemaList, defaults.SchemaMetricResourceNode())

	// Monitor
	schemaList = append(schemaList, defaults.SchemaMonitorKafka())
	schemaList = append(schemaList, defaults.SchemaMonitorKafkaCG())

	// Planning
	schemaList = append(schemaList, defaults.SchemaPlanningResourceApplication())
	schemaList = append(schemaList, defaults.SchemaPlanningResourceCluster())
	schemaList = append(schemaList, defaults.SchemaPlanningResourceContainer())
	schemaList = append(schemaList, defaults.SchemaPlanningResourceController())
	schemaList = append(schemaList, defaults.SchemaPlanningResourceNamespace())
	schemaList = append(schemaList, defaults.SchemaPlanningResourceNode())

	// Prediction
	schemaList = append(schemaList, defaults.SchemaPredictionKafkaTopic())
	schemaList = append(schemaList, defaults.SchemaPredictionKafkaCG())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceApplication())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceCluster())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceContainer())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceController())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceNamespace())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceNode())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceNodeGroup())
	schemaList = append(schemaList, defaults.SchemaPredictionResourceMachineset())

	// Recommendation
	schemaList = append(schemaList, defaults.SchemaRecommendationKafkaCG())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceApplication())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceCluster())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceContainer())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceController())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceNamespace())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceNode())
	schemaList = append(schemaList, defaults.SchemaRecommendationResourceMachineset())

	// Resource
	schemaList = append(schemaList, defaults.SchemaResourceApplication())
	schemaList = append(schemaList, defaults.SchemaResourceCluster())
	schemaList = append(schemaList, defaults.SchemaResourceClusterScaler())
	schemaList = append(schemaList, defaults.SchemaResourceContainer())
	schemaList = append(schemaList, defaults.SchemaResourceController())
	schemaList = append(schemaList, defaults.SchemaResourceMachineScaler())
	schemaList = append(schemaList, defaults.SchemaMachineset())
	schemaList = append(schemaList, defaults.SchemaResourceNamespace())
	schemaList = append(schemaList, defaults.SchemaResourceNode())
	schemaList = append(schemaList, defaults.SchemaResourcePod())

	schemaMgt.AddSchemas(schemaList)
	schemaMgt.Flush() // TODO: ONLY DO one time !!!
}
