package metrics

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricModelTimeGauges   = map[string]*prometheus.GaugeVec{}
	metricModelTimeCounters = map[string]*prometheus.CounterVec{}
	metricMAPEGauges        = map[string]*prometheus.GaugeVec{}
	metricRMSEGauges        = map[string]*prometheus.GaugeVec{}
	metricDriftCounters     = map[string]*prometheus.CounterVec{}
)

func InitMetric(unitScope, category, unitType string, idKeys []string) {
	labels := []string{}
	labels = append(labels, "scope", "category", "type")
	labels = append(labels, idKeys...)
	labels = append(labels, "data_granularity", "metric_type", "export_timestamp")

	metricModelTimeGauges[fmt.Sprintf("%s/%s/%s", unitScope, category, unitType)] = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      fmt.Sprintf("%s_%s_model_seconds", category, unitType),
		Help:      fmt.Sprintf("Target modeling time of %s %s metric", category, unitType),
	}, labels)
	metricModelTimeCounters[fmt.Sprintf("%s/%s/%s", unitScope, category, unitType)] = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      fmt.Sprintf("%s_%s_model_seconds_total", category, unitType),
		Help:      fmt.Sprintf("Total target modeling time of %s %s metric", category, unitType),
	}, labels)
	metricMAPEGauges[fmt.Sprintf("%s/%s/%s", unitScope, category, unitType)] = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      fmt.Sprintf("%s_%s_metric_mape", category, unitType),
		Help:      fmt.Sprintf("MAPE of %s %s metric", category, unitType),
	}, labels)
	metricRMSEGauges[fmt.Sprintf("%s/%s/%s", unitScope, category, unitType)] = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      fmt.Sprintf("%s_%s_metric_rmse", category, unitType),
		Help:      fmt.Sprintf("RMSE of %s %s metric", category, unitType),
	}, labels)
	metricDriftCounters[fmt.Sprintf("%s/%s/%s", unitScope, category, unitType)] = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      fmt.Sprintf("%s_%s_metric_drift_total", category, unitType),
		Help:      fmt.Sprintf("Total number of %s %s metric drift", category, unitType),
	}, labels)
}

func SetMetricModelTime(jobID string, val float64) {
	for k, v := range metricModelTimeGauges {
		if strings.HasPrefix(jobID, k) {
			labels := strings.Split(jobID, "/")
			exportTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
			labels = append(labels, exportTimestamp)
			v.WithLabelValues(labels...).Set(val)
			break
		}
	}
}
func AddMetricModelTime(jobID string, val float64) {
	for k, v := range metricModelTimeCounters {
		if strings.HasPrefix(jobID, k) {
			labels := strings.Split(jobID, "/")
			exportTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
			labels = append(labels, exportTimestamp)
			v.WithLabelValues(labels...).Add(val)
			break
		}
	}
}
func SetMetricMAPE(jobID string, val float64) {
	for k, v := range metricMAPEGauges {
		if strings.HasPrefix(jobID, k) {
			labels := strings.Split(jobID, "/")
			exportTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
			labels = append(labels, exportTimestamp)
			v.WithLabelValues(labels...).Set(val)
			break
		}
	}
}
func SetMetricRMSE(jobID string, val float64) {
	for k, v := range metricRMSEGauges {
		if strings.HasPrefix(jobID, k) {
			labels := strings.Split(jobID, "/")
			exportTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
			labels = append(labels, exportTimestamp)
			v.WithLabelValues(labels...).Set(val)
			break
		}
	}
}
func AddMetricDrift(jobID string, val float64) {
	for k, v := range metricDriftCounters {
		if strings.HasPrefix(jobID, k) {
			labels := strings.Split(jobID, "/")
			exportTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
			labels = append(labels, exportTimestamp)
			v.WithLabelValues(labels...).Add(val)
			break
		}
	}
}
