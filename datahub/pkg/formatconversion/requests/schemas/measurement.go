package schemas

import (
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewMeasurement(measurement *ApiSchema.Measurement) *InfluxSchema.Measurement {
	if measurement != nil {
		m := InfluxSchema.Measurement{
			Name:       measurement.GetName(),
			MetricType: InfluxSchema.MetricType(measurement.GetMetricType()),
			Boundary:   InfluxSchema.ResourceBoundary(measurement.GetResourceBoundary()),
			Quota:      InfluxSchema.ResourceQuota(measurement.GetResourceQuota()),
			IsTS:       measurement.GetIsTs(),
		}
		for _, column := range measurement.Columns {
			m.Columns = append(m.Columns, NewColumn(column))
		}
		return &m
	}
	return nil
}
