package schemamgt

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

var MeasurementNameMap = map[schemas.Scope]string{
	schemas.Application:    "application",
	schemas.Execution:      "execution",
	schemas.Fedemeter:      "fedemeter",
	schemas.Metric:         "metric",
	schemas.Planning:       "planning",
	schemas.Prediction:     "prediction",
	schemas.Recommendation: "recommendation",
	schemas.Resource:       "resource",
}

var MeasurementSchemaNameMap = map[schemas.Scope]string{
	schemas.Application:    "application_schema",
	schemas.Execution:      "execution_schema",
	schemas.Fedemeter:      "fedemeter_schema",
	schemas.Metric:         "metric_schema",
	schemas.Planning:       "planning_schema",
	schemas.Prediction:     "prediction_schema",
	schemas.Recommendation: "recommendation_schema",
	schemas.Resource:       "resource_schema",
}

var DatabaseNameMap = map[schemas.Scope]string{
	schemas.Application:    "alameda_application",
	schemas.Execution:      "alameda_execution",
	schemas.Fedemeter:      "alameda_fedemeter",
	schemas.Metric:         "alameda_metric",
	schemas.Planning:       "alameda_planning",
	schemas.Prediction:     "alameda_prediction",
	schemas.Recommendation: "alameda_recommendation",
	schemas.Resource:       "alameda_cluster_status",
}
