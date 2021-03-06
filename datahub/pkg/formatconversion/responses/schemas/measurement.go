package schemas

import (
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewMeasurement(measurement *InfluxSchema.Measurement) *ApiSchema.Measurement {
	if measurement != nil {
		m := ApiSchema.Measurement{
			Name:             measurement.Name,
			MetricType:       ApiCommon.MetricType(measurement.MetricType),
			ResourceBoundary: ApiCommon.ResourceBoundary(measurement.Boundary),
			ResourceQuota:    ApiCommon.ResourceQuota(measurement.Quota),
			IsTs:             measurement.IsTS,
		}
		for _, column := range measurement.Columns {
			m.Columns = append(m.Columns, NewColumn(column))
		}
		return &m
	}
	return nil
}
