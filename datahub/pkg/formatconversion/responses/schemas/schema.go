package schemas

import (
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewSchema(schema *InfluxSchema.Schema) *ApiSchema.Schema {
	if schema != nil {
		s := ApiSchema.Schema{}
		s.SchemaMeta = NewSchemaMeta(schema.SchemaMeta)
		for _, measurement := range schema.Measurements {
			s.Measurements = append(s.Measurements, NewMeasurement(measurement))
		}
		return &s
	}
	return nil
}
