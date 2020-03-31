package datamappingmgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

type Map struct {
	DataMappings map[schemas.Scope][]*datamapping.DataMapping
}

func NewMap() *Map {
	m := Map{}
	m.DataMappings = make(map[schemas.Scope][]*datamapping.DataMapping)
	return &m
}

func (p *Map) AddDataMapping(dataMapping *datamapping.DataMapping) {
	if p.DataMappings == nil {
		p.DataMappings = make(map[schemas.Scope][]*datamapping.DataMapping)
	}

	scope := dataMapping.SchemaMeta.Scope
	if _, ok := p.DataMappings[scope]; !ok {
		p.DataMappings[scope] = make([]*datamapping.DataMapping, 0)
	}

	// Do a deep copy of schema
	d := datamapping.DataMapping{}
	utils.DeepCopy(&d, dataMapping)
	p.DataMappings[scope] = append(p.DataMappings[scope], &d)
}

// Initialize() must be called after Empty()
func (p *Map) Initialize() {
	p.DataMappings = make(map[schemas.Scope][]*datamapping.DataMapping)
	for scope := range MeasurementNameMap {
		p.DataMappings[scope] = make([]*datamapping.DataMapping, 0)
	}
}

func (p *Map) Empty() {
	if p.DataMappings != nil {
		for k := range p.DataMappings {
			delete(p.DataMappings, k)
		}
		p.DataMappings = nil
	}
}
