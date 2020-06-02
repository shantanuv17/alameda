package enumconv

import (
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var MetricTypeNameMap = map[schemas.MetricType]common.MetricType{
	schemas.MetricTypeUndefined:       common.MetricType_METRICS_TYPE_UNDEFINED,
	schemas.CPUUsageSecondsPercentage: common.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
	schemas.MemoryUsageBytes:          common.MetricType_MEMORY_USAGE_BYTES,
	schemas.PowerUsageWatts:           common.MetricType_POWER_USAGE_WATTS,
	schemas.TemperatureCelsius:        common.MetricType_TEMPERATURE_CELSIUS,
	schemas.DutyCycle:                 common.MetricType_DUTY_CYCLE,
	schemas.CurrentOffset:             common.MetricType_CURRENT_OFFSET,
	schemas.Lag:                       common.MetricType_LAG,
	schemas.Latency:                   common.MetricType_LATENCY,
	schemas.Number:                    common.MetricType_NUMBER,
}

var ResourceBoundaryNameMap = map[schemas.ResourceBoundary]common.ResourceBoundary{
	schemas.ResourceBoundaryUndefined: common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED,
	schemas.ResourceRaw:               common.ResourceBoundary_RESOURCE_RAW,
	schemas.ResourceUpperBound:        common.ResourceBoundary_RESOURCE_UPPER_BOUND,
	schemas.ResourceLowerBound:        common.ResourceBoundary_RESOURCE_LOWER_BOUND,
}

var ResourceQuotaNameMap = map[schemas.ResourceQuota]common.ResourceQuota{
	schemas.ResourceQuotaUndefined: common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED,
	schemas.ResourceLimit:          common.ResourceQuota_RESOURCE_LIMIT,
	schemas.ResourceRequest:        common.ResourceQuota_RESOURCE_REQUEST,
	schemas.ResourceInitialLimit:   common.ResourceQuota_RESOURCE_INITIAL_LIMIT,
	schemas.ResourceInitialRequest: common.ResourceQuota_RESOURCE_INITIAL_REQUEST,
}

var QueryConditionFunctionNameMap = map[DBCommon.FunctionType]common.FunctionType{
	DBCommon.NoneFunction:   common.FunctionType_FUNCTIONTYPE_UNDEFINED,
	DBCommon.FunctionCount:  common.FunctionType_FUNCTIONTYPE_COUNT,
	DBCommon.FunctionMean:   common.FunctionType_FUNCTIONTYPE_MEAN,
	DBCommon.FunctionMedian: common.FunctionType_FUNCTIONTYPE_MEDIAN,
	DBCommon.FunctionStddev: common.FunctionType_FUNCTIONTYPE_STDDEV,
	DBCommon.FunctionSum:    common.FunctionType_FUNCTIONTYPE_SUM,
	DBCommon.FunctionBottom: common.FunctionType_FUNCTIONTYPE_BOTTOM,
	DBCommon.FunctionFirst:  common.FunctionType_FUNCTIONTYPE_FIRST,
	DBCommon.FunctionLast:   common.FunctionType_FUNCTIONTYPE_LAST,
	DBCommon.FunctionMax:    common.FunctionType_FUNCTIONTYPE_MAX,
	DBCommon.FunctionMin:    common.FunctionType_FUNCTIONTYPE_MIN,
	DBCommon.FunctionTop:    common.FunctionType_FUNCTIONTYPE_TOP,
}

var QueryConditionOrderNameMap = map[DBCommon.Order]common.QueryCondition_Order{
	DBCommon.NoneOrder: common.QueryCondition_NONE,
	DBCommon.Asc:       common.QueryCondition_ASC,
	DBCommon.Desc:      common.QueryCondition_DESC,
}
