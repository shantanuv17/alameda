package defaults

import (
	"github.com/containers-ai/alameda/datahub/pkg/entities"
)

func DefaultSchemaMetering() []interface{} {
	schemas := make([]interface{}, 0)

	// Federatorai
	schemas = append(schemas, &entities.MeteringFederatoraiCapacity{})

	return schemas
}
