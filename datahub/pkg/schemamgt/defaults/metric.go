package defaults

import (
	"prophetstor.com/alameda/datahub/pkg/entities"
)

func DefaultSchemaMetric() []interface{} {
	schemas := make([]interface{}, 0)

	// Kafka topic
	schemas = append(schemas, &entities.MetricKafkaTopicCurrentOffset{})

	// Kafka consumer group
	schemas = append(schemas, &entities.MetricKafkaConsumerGroupCurrentOffset{})
	schemas = append(schemas, &entities.MetricKafkaConsumerGroupLag{})

	// Cluster-status application
	schemas = append(schemas, &entities.MetricClusterStatusApplicationCPUUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusApplicationMemoryUsage{})

	// Cluster-status cluster
	schemas = append(schemas, &entities.MetricClusterStatusClusterCPUUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusClusterMemoryUsage{})

	// Cluster-status container
	schemas = append(schemas, &entities.MetricClusterStatusContainerCPUUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusContainerMemoryUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusContainerRscReqCPU{})
	schemas = append(schemas, &entities.MetricClusterStatusContainerRscLimitCPU{})
	schemas = append(schemas, &entities.MetricClusterStatusContainerRscReqMemory{})
	schemas = append(schemas, &entities.MetricClusterStatusContainerRscLimitMemory{})
	schemas = append(schemas, &entities.MetricClusterStatusContainerRestartsTotal{})

	// Cluster-status controller
	schemas = append(schemas, &entities.MetricClusterStatusControllerCPUUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusControllerMemoryUsage{})

	// Cluster-status namespace
	schemas = append(schemas, &entities.MetricClusterStatusNamespaceCPUUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusNamespaceMemoryUsage{})

	// Cluster-status node
	schemas = append(schemas, &entities.MetricClusterStatusNodeCPUAllocatable{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeCPUTotal{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeCPUUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeMemoryTotal{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeMemoryUsage{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeFSPCT{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeDiskIOUtil{})
	schemas = append(schemas, &entities.MetricClusterStatusNodePodPhaseCount{})
	schemas = append(schemas, &entities.MetricClusterStatusNodeUnschedulable{})

	// Cluster-status service
	schemas = append(schemas, &entities.MetricClusterStatusServiceHealth{})

	// Cluster-status top container
	schemas = append(schemas, &entities.MetricTopContainerCPUUsagePCT{})
	schemas = append(schemas, &entities.MetricTopContainerMemoryUsage{})

	return schemas
}
