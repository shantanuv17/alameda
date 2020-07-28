package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaApplication() []interface{} {
	schemas := make([]interface{}, 0)

	// Kafka topic
	schemas = append(schemas, &entities.ApplicationKafkaTopic{})

	// Kafka consumer group
	schemas = append(schemas, &entities.ApplicationKafkaConsumerGroup{})

	// Nginx
	schemas = append(schemas, &entities.ApplicationNginx{})

	return schemas
}
