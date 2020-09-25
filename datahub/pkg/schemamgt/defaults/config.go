package defaults

import (
	"prophetstor.com/alameda/datahub/pkg/entities"
)

func DefaultSchemaConfig() []interface{} {
	schemas := make([]interface{}, 0)

	// Tenancy cluster
	schemas = append(schemas, &entities.ConfigTenancyTenant{})

	// Tenancy organization
	schemas = append(schemas, &entities.ConfigTenancyOrganization{})

	// Tenancy cluster
	schemas = append(schemas, &entities.ConfigTenancyCluster{})

	return schemas
}
