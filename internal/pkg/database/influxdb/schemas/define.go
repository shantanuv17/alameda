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
	CPUUsageSecondsPercentage
	MemoryUsageBytes
	PowerUsageWatts
	TemperatureCelsius
	DutyCycle
	CurrentOffset
	Lag
	Latency
	Number
	CPUCores
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
