package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricModelTimeGauges   = []*prometheus.GaugeVec{}
	metricModelTimeCounters = []*prometheus.CounterVec{}
	metricMAPEGauges        = []*prometheus.GaugeVec{}
	metricRMSEGauges        = []*prometheus.GaugeVec{}
	metricDriftCounters     = []*prometheus.CounterVec{}
)
