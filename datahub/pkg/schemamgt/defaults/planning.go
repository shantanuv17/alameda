package defaults

import (
	"prophetstor.com/alameda/datahub/pkg/entities"
)

/*
  Planning type: "PT_UNDEFINED", "PT_RECOMMENDATION", "PT_PLANNING"
  Kind: "KIND_UNDEFINED", "DEPLOYMENT", "DEPLOYMENTCONFIG", "STATEFULSET", "ALAMEDASCALER"
*/

func DefaultSchemaPlanning() []interface{} {
	schemas := make([]interface{}, 0)

	// Cluster-status application
	schemas = append(schemas, &entities.PlanningClusterStatusApplication{})

	// Cluster-status cluster
	schemas = append(schemas, &entities.PlanningClusterStatusCluster{})

	// Cluster-status container
	schemas = append(schemas, &entities.PlanningClusterStatusContainer{})

	// Cluster-status controller
	schemas = append(schemas, &entities.PlanningClusterStatusController{})

	// Cluster-status namespace
	schemas = append(schemas, &entities.PlanningClusterStatusNamespace{})

	// Cluster-status node
	schemas = append(schemas, &entities.PlanningClusterStatusNode{})

	return schemas
}
