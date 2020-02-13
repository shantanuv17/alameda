package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricModelTimeGauges   = map[string]*prometheus.GaugeVec{}
	metricModelTimeCounters = map[string]*prometheus.CounterVec{}
	metricMAPEGauges        = map[string]*prometheus.GaugeVec{}
	metricRMSEGauges        = map[string]*prometheus.GaugeVec{}
	metricDriftCounters     = map[string]*prometheus.CounterVec{}
)
