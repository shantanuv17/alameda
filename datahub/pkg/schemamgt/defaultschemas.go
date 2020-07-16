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

	// Fedemeter
	schemaList = append(schemaList, defaults.SchemaFedemeterCalculation())
	schemaList = append(schemaList, defaults.SchemaFedemeterRecommendationJERI())
	schemaList = append(schemaList, defaults.SchemaFedemeterResourceHistoryCost())
	schemaList = append(schemaList, defaults.SchemaFedemeterResourcePredictionCost())

	// Metric
	schemaList = append(schemaList, defaults.SchemaMetricKafkaTopic())
	schemaList = append(schemaList, defaults.SchemaMetricKafkaCG())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusApplication())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusCluster())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusContainer())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusController())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusNamespace())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusNode())
	schemaList = append(schemaList, defaults.SchemaMetricClusterStatusService())
	schemaList = append(schemaList, defaults.SchemaMetricTopContainer())

	// Planning
	schemaList = append(schemaList, defaults.SchemaPlanningClusterStatusApplication())
	schemaList = append(schemaList, defaults.SchemaPlanningClusterStatusCluster())
	schemaList = append(schemaList, defaults.SchemaPlanningClusterStatusContainer())
	schemaList = append(schemaList, defaults.SchemaPlanningClusterStatusController())
	schemaList = append(schemaList, defaults.SchemaPlanningClusterStatusNamespace())
	schemaList = append(schemaList, defaults.SchemaPlanningClusterStatusNode())

	// Prediction
	schemaList = append(schemaList, defaults.SchemaPredictionKafkaTopic())
	schemaList = append(schemaList, defaults.SchemaPredictionKafkaCG())
	schemaList = append(schemaList, defaults.SchemaPredictionClusterStatusApplication())
	schemaList = append(schemaList, defaults.SchemaPredictionClusterStatusCluster())
	schemaList = append(schemaList, defaults.SchemaPredictionClusterStatusContainer())
	schemaList = append(schemaList, defaults.SchemaPredictionClusterStatusController())
	schemaList = append(schemaList, defaults.SchemaPredictionClusterStatusNamespace())
	schemaList = append(schemaList, defaults.SchemaPredictionClusterStatusNode())

	// Recommendation
	schemaList = append(schemaList, defaults.SchemaRecommendationKafkaCG())
	schemaList = append(schemaList, defaults.SchemaRecommendationClusterStatusApplication())
	schemaList = append(schemaList, defaults.SchemaRecommendationClusterStatusCluster())
	schemaList = append(schemaList, defaults.SchemaRecommendationClusterStatusContainer())
	schemaList = append(schemaList, defaults.SchemaRecommendationClusterStatusController())
	schemaList = append(schemaList, defaults.SchemaRecommendationClusterStatusNamespace())
	schemaList = append(schemaList, defaults.SchemaRecommendationClusterStatusNode())

	// Resource
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusApplication())
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusCluster())
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusContainer())
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusController())
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusNamespace())
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusNode())
	schemaList = append(schemaList, defaults.SchemaResourceClusterStatusPod())

	// Target
	schemaList = append(schemaList, defaults.SchemaTargetController())
	schemaList = append(schemaList, defaults.SchemaTargetKafkaTopic())
	schemaList = append(schemaList, defaults.SchemaTargetKafkaConsumerGroup())

	schemaMgt.AddSchemas(schemaList)
	schemaMgt.Flush() // TODO: ONLY DO one time !!!
}
