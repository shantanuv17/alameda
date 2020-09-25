package schemas

import (
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiSchema "prophetstor.com/api/datahub/schemas"
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
