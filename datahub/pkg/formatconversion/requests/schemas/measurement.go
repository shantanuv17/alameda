package schemas

import (
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiSchema "prophetstor.com/api/datahub/schemas"
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
