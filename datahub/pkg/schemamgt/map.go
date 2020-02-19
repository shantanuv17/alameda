package schemamgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
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
