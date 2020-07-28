package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaTarget() []interface{} {
	schemas := make([]interface{}, 0)

	// Cluster-status cluster
	schemas = append(schemas, &entities.TargetClusterStatusCluster{})

	// Cluster-status controller
	schemas = append(schemas, &entities.TargetClusterStatusController{})

	// Kafka topic
	schemas = append(schemas, &entities.TargetKafkaTopic{})

	// Kafka consumer group
	schemas = append(schemas, &entities.TargetKafkaConsumerGroup{})

	return schemas
}
