package enumconv

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

var ScopeNameMap = map[ApiSchema.Scope]schemas.Scope{
	ApiSchema.Scope_SCOPE_UNDEFINED:      schemas.ScopeUndefined,
	ApiSchema.Scope_SCOPE_APPLICATION:    schemas.Application,
	ApiSchema.Scope_SCOPE_METRIC:         schemas.Metric,
	ApiSchema.Scope_SCOPE_MONITOR:        schemas.Monitor,
	ApiSchema.Scope_SCOPE_PLANNING:       schemas.Planning,
	ApiSchema.Scope_SCOPE_PREDICTION:     schemas.Prediction,
	ApiSchema.Scope_SCOPE_RECOMMENDATION: schemas.Recommendation,
	ApiSchema.Scope_SCOPE_RESOURCE:       schemas.Resource,
}
