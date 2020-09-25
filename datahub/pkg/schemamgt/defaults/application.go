package defaults

import (
	"prophetstor.com/alameda/datahub/pkg/entities"
)

func DefaultSchemaApplication() []interface{} {
	schemas := make([]interface{}, 0)

	// Kafka topic
	schemas = append(schemas, &entities.ApplicationKafkaTopic{})

	// Kafka consumer group
	schemas = append(schemas, &entities.ApplicationKafkaConsumerGroup{})

	return schemas
}
