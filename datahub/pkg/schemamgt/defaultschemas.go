package schemamgt

import (
	"prophetstor.com/alameda/datahub/pkg/schemamgt/defaults"
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
)

var DefaultSchemas = map[string][]interface{}{
	"application":    defaults.DefaultSchemaApplication(),
	"config":         defaults.DefaultSchemaConfig(),
	"fedemeter":      defaults.DefaultSchemaFedemeter(),
	"metering":       defaults.DefaultSchemaMetering(),
	"metric":         defaults.DefaultSchemaMetric(),
	"planning":       defaults.DefaultSchemaPlanning(),
	"prediction":     defaults.DefaultSchemaPrediction(),
	"recommendation": defaults.DefaultSchemaRecommendation(),
	"resource":       defaults.DefaultSchemaResource(),
	"target":         defaults.DefaultSchemaTarget(),
}

func DefaultSchemasInit() {
	schemaMgt := NewSchemaManagement()
	schemaList := make([]*schemas.Schema, 0)

	for key, entities := range DefaultSchemas {
		scope.Infof("configuring %s schemas", key)
		for _, entity := range entities {
			schema := NewSchema(entity)
			schemaList = append(schemaList, schema)
		}
	}

	schemaMgt.AddSchemas(schemaList)
	schemaMgt.Flush() // TODO: ONLY DO one time !!!
}
