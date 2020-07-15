package schemas

import (
	InternalSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewMeasurement(measurement *ApiSchema.Measurement) *InternalSchema.Measurement {
	if measurement != nil {
		m := InternalSchema.Measurement{
			Name:       measurement.GetName(),
			MetricType: InternalSchema.MetricType(measurement.GetMetricType()),
			Boundary:   InternalSchema.ResourceBoundary(measurement.GetResourceBoundary()),
			Quota:      InternalSchema.ResourceQuota(measurement.GetResourceQuota()),
		}
		for _, column := range measurement.Columns {
			m.Columns = append(m.Columns, NewColumn(column))
		}
		return &m
	}
	return nil
}
