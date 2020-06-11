package datahub

import (
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

// Order enumerator
type Order = int

// Function type enumerator
type FunctionType = int

// Sort order definition
const (
	NoneOrder Order = iota
	Asc  // Represent ascending order
	Desc // Represent descending order
)

const (
	NoneFunction FunctionType = iota

	// Aggregation Function
	FunctionCount  // Returns the number of non-null field values
	FuncDistinct   // Returns the list of unique field values
	FuncIntegral   // Returns the area under the curve for subsequent field values
    FunctionMean   // Returns the arithmetic mean (average) of field values
    FunctionMedian // Returns the middle value from a sorted list of field values
    FunctionMode   // Returns the most frequent value in a list of field values
    FunctionSpread // Returns the difference between the minimum and maximum field values
    FunctionStddev // Returns the standard deviation of field values
    FunctionSum    // Returns the sum of field values

    // Selector function
    FunctionBottom     // Returns the smallest N field values
    FunctionFirst      // Returns the field value with the oldest timestamp
    FunctionLast       // Returns the field value with the most recent timestamp
    FunctionMax        // Returns the greatest field value
    FunctionMin        // Returns the lowest field value
    FunctionPercentile // Returns the Nth percentile field value
    FunctionSample     // Returns a random sample of N field values. SAMPLE() uses reservoir sampling to generate the random points
    FunctionTop        // Returns the greatest N field values

    // Transformation function
    FuncDerivative // Returns the rate of change between subsequent field values
)

var ScopeValue = map[string]schemas.Scope{
	"undefined":      schemas.Scope_SCOPE_UNDEFINED,
	"application":    schemas.Scope_SCOPE_APPLICATION,
	"execution":      schemas.Scope_SCOPE_EXECUTION,
	"metric":         schemas.Scope_SCOPE_METRIC,
	"planning":       schemas.Scope_SCOPE_PLANNING,
	"prediction":     schemas.Scope_SCOPE_PREDICTION,
	"recommendation": schemas.Scope_SCOPE_RECOMMENDATION,
	"resource":       schemas.Scope_SCOPE_RESOURCE,
}

var MetricTypeValue = map[string]common.MetricType{
	"undefined":            common.MetricType_METRICS_TYPE_UNDEFINED,
	"cpu_usage_percentage": common.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
	"memory_usage_bytes":   common.MetricType_MEMORY_USAGE_BYTES,
	"power_usage_watts":    common.MetricType_POWER_USAGE_WATTS,
	"temperature_celsius":  common.MetricType_TEMPERATURE_CELSIUS,
	"duty_cycle":           common.MetricType_DUTY_CYCLE,
	"current_offset":       common.MetricType_CURRENT_OFFSET,
	"lag":                  common.MetricType_LAG,
	"latency":              common.MetricType_LATENCY,
	"number":               common.MetricType_NUMBER,
	"cpu_cores":            common.MetricType_CPU_CORES,
}

var ResourceBoundaryValue = map[string]common.ResourceBoundary{
	"undefined":   common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED,
	"raw":         common.ResourceBoundary_RESOURCE_RAW,
	"upper_bound": common.ResourceBoundary_RESOURCE_UPPER_BOUND,
	"lower_bound": common.ResourceBoundary_RESOURCE_LOWER_BOUND,
}

var ResourceQuotaValue = map[string]common.ResourceQuota{
	"undefined":       common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED,
	"limit":           common.ResourceQuota_RESOURCE_LIMIT,
	"request":         common.ResourceQuota_RESOURCE_REQUEST,
	"initial_limit":   common.ResourceQuota_RESOURCE_INITIAL_LIMIT,
	"initial_request": common.ResourceQuota_RESOURCE_INITIAL_REQUEST,
}
