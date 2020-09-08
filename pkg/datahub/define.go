package datahub

import (
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

var ScopeValue = map[string]schemas.Scope{
	"undefined":      schemas.Scope_SCOPE_UNDEFINED,
	"application":    schemas.Scope_SCOPE_APPLICATION,
	"config":         schemas.Scope_SCOPE_CONFIG,
	"fedemeter":      schemas.Scope_SCOPE_FEDEMETER,
	"metric":         schemas.Scope_SCOPE_METRIC,
	"planning":       schemas.Scope_SCOPE_PLANNING,
	"prediction":     schemas.Scope_SCOPE_PREDICTION,
	"recommendation": schemas.Scope_SCOPE_RECOMMENDATION,
	"resource":       schemas.Scope_SCOPE_RESOURCE,
	"target":         schemas.Scope_SCOPE_TARGET,
}

var MetricTypeValue = map[string]common.MetricType{
	"undefined":                common.MetricType_METRICS_TYPE_UNDEFINED,
	"cpu_seconds_total":        common.MetricType_CPU_SECONDS_TOTAL,
	"cpu_cores_alloc":          common.MetricType_CPU_CORES_ALLOCATABLE,
	"cpu_millicores_total":     common.MetricType_CPU_MILLICORES_TOTAL,
	"cpu_millicores_avail":     common.MetricType_CPU_MILLICORES_AVAIL,
	"cpu_millicores_usage":     common.MetricType_CPU_MILLICORES_USAGE,
	"cpu_millicores_usage_pct": common.MetricType_CPU_MILLICORES_USAGE_PCT,
	"memory_bytes_alloc":       common.MetricType_MEMORY_BYTES_ALLOCATABLE,
	"memory_bytes_total":       common.MetricType_MEMORY_BYTES_TOTAL,
	"memory_bytes_avail":       common.MetricType_MEMORY_BYTES_AVAIL,
	"memory_bytes_usage":       common.MetricType_MEMORY_BYTES_USAGE,
	"memory_bytes_usage_pct":   common.MetricType_MEMORY_BYTES_USAGE_PCT,
	"fs_bytes_total":           common.MetricType_FS_BYTES_TOTAL,
	"fs_bytes_avail":           common.MetricType_FS_BYTES_AVAIL,
	"fs_bytes_usage":           common.MetricType_FS_BYTES_USAGE,
	"fs_bytes_usage_pct":       common.MetricType_FS_BYTES_USAGE_PCT,
	"http_req_count":           common.MetricType_HTTP_REQUESTS_COUNT,
	"http_req_total":           common.MetricType_HTTP_REQUESTS_TOTAL,
	"http_response_count":      common.MetricType_HTTP_RESPONSE_COUNT,
	"http_response_total":      common.MetricType_HTTP_RESPONSE_TOTAL,
	"disk_io_seconds_total":    common.MetricType_DISK_IO_SECONDS_TOTAL,
	"disk_io_util":             common.MetricType_DISK_IO_UTILIZATION,
	"restarts_total":           common.MetricType_RESTARTS_TOTAL,
	"unschedulable":            common.MetricType_UNSCHEDULABLE,
	"health":                   common.MetricType_HEALTH,
	"power_usage_watts":        common.MetricType_POWER_USAGE_WATTS,
	"temperature_celsius":      common.MetricType_TEMPERATURE_CELSIUS,
	"duty_cycle":               common.MetricType_DUTY_CYCLE,
	"current_offset":           common.MetricType_CURRENT_OFFSET,
	"lag":                      common.MetricType_LAG,
	"latency":                  common.MetricType_LATENCY,
	"number":                   common.MetricType_NUMBER,
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
