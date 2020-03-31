package datamapping

type Source int

// Table enumerator
const (
	SourceUndefined Source = iota
	Datadog
	Dynatrace
	K8s
	Prometheus
	Sysdig
)
