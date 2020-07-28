package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaPrediction() []interface{} {
	schemas := make([]interface{}, 0)

	// Kafka topic
	schemas = append(schemas, &entities.PredictionKafkaTopicCurrentOffset{})
	schemas = append(schemas, &entities.PredictionKafkaTopicCurrentOffsetUpperBound{})
	schemas = append(schemas, &entities.PredictionKafkaTopicCurrentOffsetLowerBound{})

	// Kafka consumer group
	schemas = append(schemas, &entities.PredictionKafkaConsumerGroupCurrentOffset{})
	schemas = append(schemas, &entities.PredictionKafkaConsumerGroupCurrentOffsetUpperBound{})
	schemas = append(schemas, &entities.PredictionKafkaConsumerGroupCurrentOffsetLowerBound{})

	// Nginx
	schemas = append(schemas, &entities.PredictionNginxHttpResponseTotal{})
	schemas = append(schemas, &entities.PredictionNginxHttpResponseTotalUpperBound{})
	schemas = append(schemas, &entities.PredictionNginxHttpResponseTotalLowerBound{})

	// Cluster-autoscaler machinegroup
	schemas = append(schemas, &entities.PredictionClusterAutoscalerMachinegroupCPU{})
	schemas = append(schemas, &entities.PredictionClusterAutoscalerMachinegroupCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterAutoscalerMachinegroupCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterAutoscalerMachinegroupMemory{})
	schemas = append(schemas, &entities.PredictionClusterAutoscalerMachinegroupMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterAutoscalerMachinegroupMemoryLowerBound{})

	// Cluster-status application
	schemas = append(schemas, &entities.PredictionClusterStatusApplicationCPU{})
	schemas = append(schemas, &entities.PredictionClusterStatusApplicationCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusApplicationCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusApplicationMemory{})
	schemas = append(schemas, &entities.PredictionClusterStatusApplicationMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusApplicationMemoryLowerBound{})

	// Cluster-status cluster
	schemas = append(schemas, &entities.PredictionClusterStatusClusterCPU{})
	schemas = append(schemas, &entities.PredictionClusterStatusClusterCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusClusterCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusClusterMemory{})
	schemas = append(schemas, &entities.PredictionClusterStatusClusterMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusClusterMemoryLowerBound{})

	// Cluster-status container
	schemas = append(schemas, &entities.PredictionClusterStatusContainerCPU{})
	schemas = append(schemas, &entities.PredictionClusterStatusContainerCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusContainerCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusContainerMemory{})
	schemas = append(schemas, &entities.PredictionClusterStatusContainerMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusContainerMemoryLowerBound{})

	// Cluster-status controller
	schemas = append(schemas, &entities.PredictionClusterStatusControllerCPU{})
	schemas = append(schemas, &entities.PredictionClusterStatusControllerCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusControllerCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusControllerMemory{})
	schemas = append(schemas, &entities.PredictionClusterStatusControllerMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusControllerMemoryLowerBound{})

	// Cluster-status namespace
	schemas = append(schemas, &entities.PredictionClusterStatusNamespaceCPU{})
	schemas = append(schemas, &entities.PredictionClusterStatusNamespaceCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusNamespaceCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusNamespaceMemory{})
	schemas = append(schemas, &entities.PredictionClusterStatusNamespaceMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusNamespaceMemoryLowerBound{})

	// Cluster-status node
	schemas = append(schemas, &entities.PredictionClusterStatusNodeCPU{})
	schemas = append(schemas, &entities.PredictionClusterStatusNodeCPUUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusNodeCPULowerBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusNodeMemory{})
	schemas = append(schemas, &entities.PredictionClusterStatusNodeMemoryUpperBound{})
	schemas = append(schemas, &entities.PredictionClusterStatusNodeMemoryLowerBound{})

	return schemas
}
