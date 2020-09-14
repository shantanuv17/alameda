package schemas

type Scope int

// Table enumerator
const (
	ScopeUndefined Scope = iota
	Application
	Config
	Execution
	Fedemeter
	Metering
	Metric
	Planning
	Prediction
	Recommendation
	Resource
	Target
)

type MetricType int

// Metric type enumerator
const (
	MetricTypeUndefined MetricType = iota
	CPUSecondsTotal
	CPUCoresAllocatable
	CPUMilliCoresTotal
	CPUMilliCoresAvail
	CPUMilliCoresUsage
	CPUMilliCoresUsagePCT
	MemoryBytesAllocatable
	MemoryBytesTotal
	MemoryBytesAvail
	MemoryBytesUsage
	MemoryBytesUsagePCT
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
