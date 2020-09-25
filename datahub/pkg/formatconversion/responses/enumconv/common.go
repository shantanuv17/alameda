package enumconv

import (
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
	"prophetstor.com/api/datahub/common"
)

var MetricTypeNameMap = map[schemas.MetricType]common.MetricType{
	schemas.MetricTypeUndefined:    common.MetricType_METRICS_TYPE_UNDEFINED,
	schemas.CPUSecondsTotal:        common.MetricType_CPU_SECONDS_TOTAL,
	schemas.CPUCoresAllocatable:    common.MetricType_CPU_CORES_ALLOCATABLE,
	schemas.CPUMilliCoresTotal:     common.MetricType_CPU_MILLICORES_TOTAL,
	schemas.CPUMilliCoresAvail:     common.MetricType_CPU_MILLICORES_AVAIL,
	schemas.CPUMilliCoresUsage:     common.MetricType_CPU_MILLICORES_USAGE,
	schemas.CPUMilliCoresUsagePCT:  common.MetricType_CPU_MILLICORES_USAGE_PCT,
	schemas.MemoryBytesAllocatable: common.MetricType_MEMORY_BYTES_ALLOCATABLE,
	schemas.MemoryBytesTotal:       common.MetricType_MEMORY_BYTES_TOTAL,
	schemas.MemoryBytesAvail:       common.MetricType_MEMORY_BYTES_AVAIL,
	schemas.MemoryBytesUsage:       common.MetricType_MEMORY_BYTES_USAGE,
	schemas.MemoryBytesUsagePCT:    common.MetricType_MEMORY_BYTES_USAGE_PCT,
	schemas.FSBytesTotal:           common.MetricType_FS_BYTES_TOTAL,
	schemas.FSBytesAvail:           common.MetricType_FS_BYTES_AVAIL,
	schemas.FSBytesUsage:           common.MetricType_FS_BYTES_USAGE,
	schemas.FSBytesUsagePCT:        common.MetricType_FS_BYTES_USAGE_PCT,
	schemas.HttpRequestsCount:      common.MetricType_HTTP_REQUESTS_COUNT,
	schemas.HttpRequestsTotal:      common.MetricType_HTTP_REQUESTS_TOTAL,
	schemas.HttpResponseCount:      common.MetricType_HTTP_RESPONSE_COUNT,
	schemas.HttpResponseTotal:      common.MetricType_HTTP_RESPONSE_TOTAL,
	schemas.DiskIOSecondsTotal:     common.MetricType_DISK_IO_SECONDS_TOTAL,
	schemas.DiskIOUtilization:      common.MetricType_DISK_IO_UTILIZATION,
	schemas.RestartsTotal:          common.MetricType_RESTARTS_TOTAL,
	schemas.Unschedulable:          common.MetricType_UNSCHEDULABLE,
	schemas.Health:                 common.MetricType_HEALTH,
	schemas.PowerUsageWatts:        common.MetricType_POWER_USAGE_WATTS,
	schemas.TemperatureCelsius:     common.MetricType_TEMPERATURE_CELSIUS,
	schemas.DutyCycle:              common.MetricType_DUTY_CYCLE,
	schemas.CurrentOffset:          common.MetricType_CURRENT_OFFSET,
	schemas.Lag:                    common.MetricType_LAG,
	schemas.Latency:                common.MetricType_LATENCY,
	schemas.Number:                 common.MetricType_NUMBER,
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
	DBCommon.NoneFunction:       common.FunctionType_FUNCTIONTYPE_UNDEFINED,
	DBCommon.FunctionCount:      common.FunctionType_FUNCTIONTYPE_COUNT,
	DBCommon.FuncDistinct:       common.FunctionType_FUNCTIONTYPE_DISTINCT,
	DBCommon.FuncIntegral:       common.FunctionType_FUNCTIONTYPE_INTEGRAL,
	DBCommon.FunctionMean:       common.FunctionType_FUNCTIONTYPE_MEAN,
	DBCommon.FunctionMedian:     common.FunctionType_FUNCTIONTYPE_MEDIAN,
	DBCommon.FunctionMode:       common.FunctionType_FUNCTIONTYPE_MODE,
	DBCommon.FunctionSpread:     common.FunctionType_FUNCTIONTYPE_SPREAD,
	DBCommon.FunctionStddev:     common.FunctionType_FUNCTIONTYPE_STDDEV,
	DBCommon.FunctionSum:        common.FunctionType_FUNCTIONTYPE_SUM,
	DBCommon.FunctionBottom:     common.FunctionType_FUNCTIONTYPE_BOTTOM,
	DBCommon.FunctionFirst:      common.FunctionType_FUNCTIONTYPE_FIRST,
	DBCommon.FunctionLast:       common.FunctionType_FUNCTIONTYPE_LAST,
	DBCommon.FunctionMax:        common.FunctionType_FUNCTIONTYPE_MAX,
	DBCommon.FunctionMin:        common.FunctionType_FUNCTIONTYPE_MIN,
	DBCommon.FunctionPercentile: common.FunctionType_FUNCTIONTYPE_PERCENTILE,
	DBCommon.FunctionSample:     common.FunctionType_FUNCTIONTYPE_SAMPLE,
	DBCommon.FunctionTop:        common.FunctionType_FUNCTIONTYPE_TOP,
	DBCommon.FuncDerivative:     common.FunctionType_FUNCTIONTYPE_DERIVATIVE,
}

var QueryConditionOrderNameMap = map[DBCommon.Order]common.QueryCondition_Order{
	DBCommon.NoneOrder: common.QueryCondition_NONE,
	DBCommon.Asc:       common.QueryCondition_ASC,
	DBCommon.Desc:      common.QueryCondition_DESC,
}
