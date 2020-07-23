package schemas

import (
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
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
