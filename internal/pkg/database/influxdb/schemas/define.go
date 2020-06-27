package schemas

type Scope int

// Table enumerator
const (
	ScopeUndefined Scope = iota
	Application
	Execution
	Metric
	Planning
	Prediction
	Recommendation
	Resource
)

type MetricType int

// Metric type enumerator
const (
	MetricTypeUndefined MetricType = iota
	CPUSecondsTotal
	CPUMilliCoresTotal
	CPUMilliCoresAvail
	CPUMilliCoresUsage // Change from CPUUsageSecondsPercentage
	CPUMilliCoresUsagePCT
	CPUMilliCoresAllocatable
	MemoryBytesTotal
	MemoryBytesAvail
	MemoryBytesUsage // Change from MemoryUsageBytes
	MemoryBytesUsagePCT
	MemoryBytesAllocatable
	FSBytesTotal
	FSBytesAvail
	FSBytesUsage
	FSBytesUsagePCT
	HttpRequestsCount
	HttpRequestsTotal
	HttpResponseCount
	HttpResponseTotal
	DiskIOSecondsTotal
	DiskIOUtilization
	RestartsTotal
	Unschedulable
	Health
	PowerUsageWatts
	TemperatureCelsius
	DutyCycle
	CurrentOffset
	Lag
	Latency
	Number
)

type ResourceBoundary int

// Resource boundary enumerator
const (
	ResourceBoundaryUndefined ResourceBoundary = iota
	ResourceRaw
	ResourceUpperBound
	ResourceLowerBound
)

type ResourceQuota int

// Resource quota enumerator
const (
	ResourceQuotaUndefined ResourceQuota = iota
	ResourceLimit
	ResourceRequest
	ResourceInitialLimit
	ResourceInitialRequest
)

type ColumnType int

// Influxdb column type enumerator
const (
	ColumnTypeUndefined ColumnType = iota
	Tag
	Field
)
