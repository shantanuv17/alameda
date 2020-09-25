package enumconv

import (
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiCommon "prophetstor.com/api/datahub/common"
)

var MetricTypeNameMap = map[ApiCommon.MetricType]schemas.MetricType{
	ApiCommon.MetricType_METRICS_TYPE_UNDEFINED:   schemas.MetricTypeUndefined,
	ApiCommon.MetricType_CPU_SECONDS_TOTAL:        schemas.CPUSecondsTotal,
	ApiCommon.MetricType_CPU_CORES_ALLOCATABLE:    schemas.CPUCoresAllocatable,
	ApiCommon.MetricType_CPU_MILLICORES_TOTAL:     schemas.CPUMilliCoresTotal,
	ApiCommon.MetricType_CPU_MILLICORES_AVAIL:     schemas.CPUMilliCoresAvail,
	ApiCommon.MetricType_CPU_MILLICORES_USAGE:     schemas.CPUMilliCoresUsage,
	ApiCommon.MetricType_CPU_MILLICORES_USAGE_PCT: schemas.CPUMilliCoresUsagePCT,
	ApiCommon.MetricType_MEMORY_BYTES_ALLOCATABLE: schemas.MemoryBytesAllocatable,
	ApiCommon.MetricType_MEMORY_BYTES_TOTAL:       schemas.MemoryBytesTotal,
	ApiCommon.MetricType_MEMORY_BYTES_AVAIL:       schemas.MemoryBytesAvail,
	ApiCommon.MetricType_MEMORY_BYTES_USAGE:       schemas.MemoryBytesUsage,
	ApiCommon.MetricType_MEMORY_BYTES_USAGE_PCT:   schemas.MemoryBytesUsagePCT,
	ApiCommon.MetricType_FS_BYTES_TOTAL:           schemas.FSBytesTotal,
	ApiCommon.MetricType_FS_BYTES_AVAIL:           schemas.FSBytesAvail,
	ApiCommon.MetricType_FS_BYTES_USAGE:           schemas.FSBytesUsage,
	ApiCommon.MetricType_FS_BYTES_USAGE_PCT:       schemas.FSBytesUsagePCT,
	ApiCommon.MetricType_HTTP_REQUESTS_COUNT:      schemas.HttpRequestsCount,
	ApiCommon.MetricType_HTTP_REQUESTS_TOTAL:      schemas.HttpRequestsTotal,
	ApiCommon.MetricType_HTTP_RESPONSE_COUNT:      schemas.HttpResponseCount,
	ApiCommon.MetricType_HTTP_RESPONSE_TOTAL:      schemas.HttpResponseTotal,
	ApiCommon.MetricType_DISK_IO_SECONDS_TOTAL:    schemas.DiskIOSecondsTotal,
	ApiCommon.MetricType_DISK_IO_UTILIZATION:      schemas.DiskIOUtilization,
	ApiCommon.MetricType_RESTARTS_TOTAL:           schemas.RestartsTotal,
	ApiCommon.MetricType_UNSCHEDULABLE:            schemas.Unschedulable,
	ApiCommon.MetricType_HEALTH:                   schemas.Health,
	ApiCommon.MetricType_POWER_USAGE_WATTS:        schemas.PowerUsageWatts,
	ApiCommon.MetricType_TEMPERATURE_CELSIUS:      schemas.TemperatureCelsius,
	ApiCommon.MetricType_DUTY_CYCLE:               schemas.DutyCycle,
	ApiCommon.MetricType_CURRENT_OFFSET:           schemas.CurrentOffset,
	ApiCommon.MetricType_LAG:                      schemas.Lag,
	ApiCommon.MetricType_LATENCY:                  schemas.Latency,
	ApiCommon.MetricType_NUMBER:                   schemas.Number,
}
