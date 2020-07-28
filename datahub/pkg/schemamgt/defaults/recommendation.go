package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaRecommendation() []interface{} {
	schemas := make([]interface{}, 0)

	// Kafka consumer group
	schemas = append(schemas, &entities.RecommendationKafkaConsumerGroup{})

	// Cluster-status application
	schemas = append(schemas, &entities.RecommendationClusterStatusApplication{})

	// Cluster-status cluster
	schemas = append(schemas, &entities.RecommendationClusterStatusCluster{})

	// Cluster-status container
	schemas = append(schemas, &entities.RecommendationClusterStatusContainerLimit{})
	schemas = append(schemas, &entities.RecommendationClusterStatusContainerRequest{})

	// Cluster-status controller
	schemas = append(schemas, &entities.RecommendationClusterStatusController{})

	// Cluster-status namespace
	schemas = append(schemas, &entities.RecommendationClusterStatusNamespace{})

	// Cluster-status node
	schemas = append(schemas, &entities.RecommendationClusterStatusNode{})

	return schemas
}
