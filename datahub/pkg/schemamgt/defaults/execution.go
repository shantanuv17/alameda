package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaExecution() []interface{} {
	schemas := make([]interface{}, 0)

	// cluster-autoscaler machineset
	schemas = append(schemas, &entities.ExecutionClusterAutoscalerMachineset{})

	return schemas
}
