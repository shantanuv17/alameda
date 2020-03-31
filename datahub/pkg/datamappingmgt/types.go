package datamappingmgt

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

const (
	DatabaseName = "alameda_data_mapping"
)

var MeasurementNameMap = map[schemas.Scope]string{
	schemas.Application:    "application",
	schemas.Metric:         "metric",
	schemas.Planning:       "planning",
	schemas.Prediction:     "prediction",
	schemas.Recommendation: "recommendation",
	schemas.Resource:       "resource",
}
