package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaResource() []interface{} {
	schemas := make([]interface{}, 0)

	// Cluster-autoscaler machinegroup
	schemas = append(schemas, &entities.ResourceClusterAutoscalerMachinegroup{})

	// Cluster-autoscaler machineset
	schemas = append(schemas, &entities.ResourceClusterAutoscalerMachineset{})

	// Cluster-status application
	schemas = append(schemas, &entities.ResourceClusterStatusApplication{})

	// Cluster-status cluster
	schemas = append(schemas, &entities.ResourceClusterStatusCluster{})

	// Cluster-status container
	schemas = append(schemas, &entities.ResourceClusterStatusContainer{})

	// Cluster-status controller
	schemas = append(schemas, &entities.ResourceClusterStatusController{})

	// Cluster-status namespace
	schemas = append(schemas, &entities.ResourceClusterStatusNamespace{})

	// Cluster-status node
	schemas = append(schemas, &entities.ResourceClusterStatusNode{})

	// Cluster-status pod
	schemas = append(schemas, &entities.ResourceClusterStatusPod{})

	return schemas
}
