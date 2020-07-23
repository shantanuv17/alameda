package schemas

import (
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewSchemaMeta(schemaMeta *ApiSchema.SchemaMeta) *InfluxSchema.SchemaMeta {
	if schemaMeta != nil {
		meta := InfluxSchema.SchemaMeta{}
		meta.Scope = InfluxSchema.Scope(schemaMeta.GetScope())
		meta.Category = schemaMeta.GetCategory()
		meta.Type = schemaMeta.GetType()
		return &meta
	}
	return nil
}
