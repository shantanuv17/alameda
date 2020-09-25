package schemamgt

import (
	"prophetstor.com/alameda/datahub/pkg/utils"
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
)

type SchemaMap struct {
	Schemas map[schemas.Scope][]*schemas.Schema
}

func NewSchemaMap() *SchemaMap {
	schemaMap := SchemaMap{}
	schemaMap.Schemas = make(map[schemas.Scope][]*schemas.Schema)
	return &schemaMap
}

func (p *SchemaMap) AddSchema(schema *schemas.Schema) {
	if p.Schemas == nil {
		p.Schemas = make(map[schemas.Scope][]*schemas.Schema)
	}

	scope := schema.SchemaMeta.Scope
	if _, ok := p.Schemas[scope]; !ok {
		p.Schemas[scope] = make([]*schemas.Schema, 0)
	}

	// Check if schema already exist
	for _, s := range p.Schemas[scope] {
		if schemas.CompareSchemaMeta(s.SchemaMeta, schema.SchemaMeta) {
			for _, m := range schema.Measurements {
				s.AddMeasurement(m.Name, m.MetricType, m.Boundary, m.Quota, m.IsTS, m.String())
			}
			return
		}
	}

	// Do a deep copy of schema
	s := schemas.Schema{}
	utils.DeepCopy(&s, schema)
	p.Schemas[scope] = append(p.Schemas[scope], &s)
}

func (p *SchemaMap) Empty() {
	if p.Schemas != nil {
		for k := range p.Schemas {
			delete(p.Schemas, k)
		}
		p.Schemas = nil
	}
}
