package schemas

import (
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiSchema "prophetstor.com/api/datahub/schemas"
)

func NewSchemaMeta(schemaMeta *InfluxSchema.SchemaMeta) *ApiSchema.SchemaMeta {
	if schemaMeta != nil {
		meta := ApiSchema.SchemaMeta{}
		meta.Scope = ApiSchema.Scope(schemaMeta.Scope)
		meta.Category = schemaMeta.Category
		meta.Type = schemaMeta.Type
		return &meta
	}
	return nil
}
