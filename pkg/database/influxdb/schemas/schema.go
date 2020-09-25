package schemas

import (
	"prophetstor.com/alameda/pkg/utils/log"
	"strings"
)

var (
	scope = log.RegisterScope("Database", "influxdb-schemas", 0)
)

type Schema struct {
	SchemaMeta   *SchemaMeta
	Measurements []*Measurement
}

func NewSchema(scope Scope, category, schemaType string) *Schema {
	schema := Schema{}
	schema.SchemaMeta = NewSchemaMeta(scope, category, schemaType)
	schema.Measurements = make([]*Measurement, 0)
	return &schema
}

func (p *Schema) AddMeasurement(name string, metricType MetricType, boundary ResourceBoundary, quota ResourceQuota, isTS bool, columns string) error {
	measurement := NewMeasurement(name, metricType, boundary, quota, isTS)
	if err := measurement.Initialize(columns); err != nil {
		return err
	}

	// Check if measurement already exists
	for _, m := range p.Measurements {
		if CompareMeasurement(m, measurement) {
			return m.Copy(measurement)
		}
	}

	// If not found, just append this measurement
	p.Measurements = append(p.Measurements, measurement)

	return nil
}

func (p *Schema) GetMeasurement(name string, metricType MetricType, boundary ResourceBoundary, quota ResourceQuota) *Measurement {
	if name != "" {
		for _, m := range p.Measurements {
			if m.Name == name && m.MetricType == metricType && m.Boundary == boundary && m.Quota == quota {
				return m
			}
		}
		return nil
	}
	for _, m := range p.Measurements {
		if m.MetricType == metricType && m.Boundary == boundary && m.Quota == quota {
			return m
		}
	}
	return nil
}

func (p *Schema) Validate() error {
	return nil
}

func (p *Schema) String() string {
	values := make([]string, 0)
	for _, measurement := range p.Measurements {
		values = append(values, measurement.Name)
	}
	return strings.Join(values, ",")
}
