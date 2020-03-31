package datamapping

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	"github.com/containers-ai/alameda/pkg/utils/log"
)

var (
	scope = log.RegisterScope("data-mapping", "data mapping library", 0)
)

type DataMapping struct {
	SchemaMeta *schemas.SchemaMeta
	MetricType schemas.MetricType
	Columns    []*Column
}

func NewDataMapping(scope schemas.Scope, category, schemaType string, metricType schemas.MetricType) *DataMapping {
	dataMapping := DataMapping{}
	dataMapping.SchemaMeta = schemas.NewSchemaMeta(scope, category, schemaType)
	dataMapping.MetricType = metricType
	dataMapping.Columns = make([]*Column, 0)
	return &dataMapping
}

func (p *DataMapping) AddColumn(name string, source Source, mapping string) {
	column := p.GetColumn(name)
	if column == nil {
		column = NewColumn(name)
		p.Columns = append(p.Columns, column)
	}
	column.AddSourceMapping(source, mapping)
}

func (p *DataMapping) GetColumn(name string) *Column {
	for _, column := range p.Columns {
		if column.Name == name {
			return column
		}
	}
	return nil
}

func (p *DataMapping) DeleteColumn(name string) error {
	found := false
	index := 0
	for index = range p.Columns {
		if p.Columns[index].Name == name {
			found = true
			break
		}
	}
	if found {
		p.Columns[index] = p.Columns[len(p.Columns)-1] // Copy last element to index i
		p.Columns[len(p.Columns)-1] = nil              // Erase last element (write zero value)
		p.Columns = p.Columns[:len(p.Columns)-1]
	}
	return nil
}
