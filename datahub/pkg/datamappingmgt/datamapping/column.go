package datamapping

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

type Column struct {
	ColumnMeta     *schemas.Column
	Name           string
	SourceMappings []*SourceMapping
}

func NewColumn(name string) *Column {
	column := Column{}
	column.Name = name
	column.SourceMappings = make([]*SourceMapping, 0)
	return &column
}

func (p *Column) AddSourceMapping(source Source, mapping string) error {
	m := p.GetSourceMapping(source)
	if m == nil {
		m = NewSourceMapping(source, mapping)
		p.SourceMappings = append(p.SourceMappings, m)
	}
	m.Mapping = mapping
	return nil
}

func (p *Column) GetSourceMapping(source Source) *SourceMapping {
	for _, sourceMapping := range p.SourceMappings {
		if sourceMapping.Source == source {
			return sourceMapping
		}
	}
	return nil
}

func (p *Column) DeleteSourceMapping(source Source) error {
	found := false
	index := 0
	for index = range p.SourceMappings {
		if p.SourceMappings[index].Source == source {
			found = true
			break
		}
	}
	if found {
		p.SourceMappings[index] = p.SourceMappings[len(p.SourceMappings)-1] // Copy last element to index i
		p.SourceMappings[len(p.SourceMappings)-1] = nil                     // Erase last element (write zero value)
		p.SourceMappings = p.SourceMappings[:len(p.SourceMappings)-1]
	}
	return nil
}
