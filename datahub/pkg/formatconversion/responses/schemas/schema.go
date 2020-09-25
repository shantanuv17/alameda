package schemas

import (
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiSchema "prophetstor.com/api/datahub/schemas"
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
