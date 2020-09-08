package schemas

var ScopeValue = map[string]Scope{
	"undefined":      ScopeUndefined,
	"application":    Application,
	"config":         Config,
	"fedemeter":      Fedemeter,
	"metric":         Metric,
	"planning":       Planning,
	"prediction":     Prediction,
	"recommendation": Recommendation,
	"resource":       Resource,
	"target":         Target,
}

var MetricTypeValue = map[string]MetricType{
	"undefined":                MetricTypeUndefined,
	"cpu_seconds_total":        CPUSecondsTotal,
	"cpu_cores_alloc":          CPUCoresAllocatable,
	"cpu_millicores_total":     CPUMilliCoresTotal,
	"cpu_millicores_avail":     CPUMilliCoresAvail,
	"cpu_millicores_usage":     CPUMilliCoresUsage,
	"cpu_millicores_usage_pct": CPUMilliCoresUsagePCT,
	"memory_bytes_alloc":       MemoryBytesAllocatable,
	"memory_bytes_total":       MemoryBytesTotal,
	"memory_bytes_avail":       MemoryBytesAvail,
	"memory_bytes_usage":       MemoryBytesUsage,
	"memory_bytes_usage_pct":   MemoryBytesUsagePCT,
	"fs_bytes_total":           FSBytesTotal,
	"fs_bytes_avail":           FSBytesAvail,
	"fs_bytes_usage":           FSBytesUsage,
	"fs_bytes_usage_pct":       FSBytesUsagePCT,
	"http_req_count":           HttpRequestsCount,
	"http_req_total":           HttpRequestsTotal,
	"http_response_count":      HttpResponseCount,
	"http_response_total":      HttpResponseTotal,
	"disk_io_seconds_total":    DiskIOSecondsTotal,
	"disk_io_util":             DiskIOUtilization,
	"restarts_total":           RestartsTotal,
	"unschedulable":            Unschedulable,
	"health":                   Health,
	"power_usage_watts":        PowerUsageWatts,
	"temperature_celsius":      TemperatureCelsius,
	"duty_cycle":               DutyCycle,
	"current_offset":           CurrentOffset,
	"lag":                      Lag,
	"latency":                  Latency,
	"number":                   Number,
}

var ResourceBoundaryValue = map[string]ResourceBoundary{
	"undefined":   ResourceBoundaryUndefined,
	"raw":         ResourceRaw,
	"upper_bound": ResourceUpperBound,
	"lower_bound": ResourceLowerBound,
}

var ResourceQuotaValue = map[string]ResourceQuota{
	"undefined":       ResourceQuotaUndefined,
	"limit":           ResourceLimit,
	"request":         ResourceRequest,
	"initial_limit":   ResourceInitialLimit,
	"initial_request": ResourceInitialRequest,
}

var ColumnTypeValue = map[string]ColumnType{
	"undefined": ColumnTypeUndefined,
	"tag":       Tag,
	"field":     Field,
}
