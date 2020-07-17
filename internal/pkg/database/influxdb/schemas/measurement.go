package schemas

import (
	"errors"
	"fmt"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"strings"
)

type Measurement struct {
	Name       string
	MetricType MetricType
	Boundary   ResourceBoundary
	Quota      ResourceQuota
	IsTS       bool
	Columns    []*Column
}

func NewMeasurement(name string, metricType MetricType, boundary ResourceBoundary, quota ResourceQuota, isTS bool) *Measurement {
	measurement := Measurement{}
	measurement.Name = name
	measurement.MetricType = metricType
	measurement.Boundary = boundary
	measurement.Quota = quota
	measurement.IsTS = isTS
	measurement.Columns = make([]*Column, 0)
	return &measurement
}

func (p *Measurement) Copy(measurement *Measurement) error {
	p.Name = measurement.Name
	p.MetricType = measurement.MetricType
	p.Boundary = measurement.Boundary
	p.Quota = measurement.Quota
	p.IsTS = measurement.IsTS
	return p.Initialize(measurement.String())
}

func (p *Measurement) Initialize(columns string) error {
	if columns != "" {
		p.Columns = nil
		columnList := strings.Split(columns, ",")
		for _, columnStr := range columnList {
			column := NewColumn()
			if err := column.Parse(columnStr); err != nil {
				return err
			}
			p.Columns = append(p.Columns, column)
		}
	}
	return nil
}

func (p *Measurement) GetTags() []*Column {
	columns := make([]*Column, 0)
	for _, column := range p.Columns {
		if column.ColumnType == Tag {
			columns = append(columns, column)
		}
	}
	return columns
}

func (p *Measurement) GetFields() []*Column {
	columns := make([]*Column, 0)
	for _, column := range p.Columns {
		if column.ColumnType == Field {
			columns = append(columns, column)
		}
	}
	return columns
}

func (p *Measurement) AddColumn(name string, required bool, columnType ColumnType, dataType common.DataType) {
	column := NewColumn()
	column.Name = name
	column.Required = required
	column.ColumnType = columnType
	column.DataType = dataType
	p.Columns = append(p.Columns, column)
}

func (p *Measurement) ColumnRequired(columns []string) error {
	// Check if column required
	for _, c := range p.Columns {
		if c.Required {
			found := false
			for _, name := range columns {
				if c.Name == name {
					found = true
					break
				}
			}
			if !found {
				return errors.New(fmt.Sprintf("column(%s) is not given in measurement(%s)", c.Name, p.Name))
			}
		}
	}
	return nil
}

func (p *Measurement) ColumnSupported(columns []string) error {
	// Check if column supported
	for _, name := range columns {
		found := false
		// Since time field is not in schema, we have to skip its column supported check
		if name == "time" {
			continue
		}
		for _, column := range p.Columns {
			if name == column.Name {
				found = true
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("column(%s) is not supported in measurement(%s)", name, p.Name))
		}
	}
	return nil
}

func (p *Measurement) ColumnTag(columns []string) error {
	tags := p.GetTags()
	for _, column := range columns {
		found := false
		for _, tag := range tags {
			if column == tag.Name {
				found = true
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("column(%s) is not tag", column))
		}
	}
	return nil
}

func (p *Measurement) String() string {
	values := make([]string, 0)
	for _, column := range p.Columns {
		values = append(values, column.String())
	}
	return strings.Join(values, ",")
}
