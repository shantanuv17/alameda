package config

import (
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

type Config struct {
	Units []Unit `mapstructure:"units"`

	// enum table definition
	MetricType  map[string]datahub_common.MetricType                  `mapstructure:"metricType"`
	Scope       map[string]datahub_schemas.Scope                      `mapstructure:"scope"`
	Aggregation map[string]datahub_common.TimeRange_AggregateFunction `mapstructure:"aggregation"`
}

// api proto enum mapping
func (cfg *Config) Init() {
	for jobIdx := range cfg.Units {
		// unit scope enum
		if val, ok := cfg.Scope[cfg.Units[jobIdx].ScopeStr]; ok {
			cfg.Units[jobIdx].Scope = val
		}

		// model metric type enum
		for _, metric := range cfg.Units[jobIdx].MetricTypeStrs {
			if val, ok := cfg.MetricType[metric]; ok {
				cfg.Units[jobIdx].MetricTypes = append(cfg.Units[jobIdx].MetricTypes, val)
			}
		}

		// metric related
		if val, ok := cfg.Scope[cfg.Units[jobIdx].Metric.ScopeStr]; ok {
			cfg.Units[jobIdx].Metric.Scope = val
		}
		if val, ok := cfg.Aggregation[cfg.Units[jobIdx].Metric.AggregationStr]; ok {
			cfg.Units[jobIdx].Metric.Aggregation = val
		}

		// prediction related
		if val, ok := cfg.Scope[cfg.Units[jobIdx].Prediction.ScopeStr]; ok {
			cfg.Units[jobIdx].Prediction.Scope = val
		}
	}
}

func (cfg *Config) GetUnits() []Unit {
	return cfg.Units
}

type Unit struct {
	Enabled        bool                        `mapstructure:"enabled"`
	Predictor      string                      `mapstructure:"predictor"`
	ScopeStr       string                      `mapstructure:"scope"`
	Scope          datahub_schemas.Scope       `mapstructure:"-"`
	Category       string                      `mapstructure:"category"`
	Type           string                      `mapstructure:"type"`
	Granularities  []string                    `mapstructure:"granularities"`
	MetricTypeStrs []string                    `mapstructure:"metricTypes"`
	Metric         metricSeriesMeta            `mapstructure:"metric"`
	Prediction     predictSeriesMeta           `mapstructure:"prediction"`
	IDKeys         []string                    `mapstructure:"idKeys"`
	Measurement    string                      `mapstructure:"measurement"`
	UnitValueKeys  *unitValueKeys              `mapstructure:"valueKeys"`
	UnitParameters *unitParameters             `mapstructure:"parameters"`
	MetricTypes    []datahub_common.MetricType `mapstructure:"-"`
}

type unitParameters struct {
	// CA
	MachineSetQueryKeys   []string `mapstructure:"machineSetQueryKeys"`
	NodeQueryKeys         []string `mapstructure:"nodeQueryKeys"`
	ClusterStatusCategory string   `mapstructure:"clusterStatusCategory"`
	MachineSetType        string   `mapstructure:"machineSetType"`
	NodeType              string   `mapstructure:"nodeType"`
}

type metricSeriesMeta struct {
	ScopeStr        string                                     `mapstructure:"scope"`
	Scope           datahub_schemas.Scope                      `mapstructure:"-"`
	Category        string                                     `mapstructure:"category"`
	Type            string                                     `mapstructure:"type"`
	AggregationStr  string                                     `mapstructure:"aggregation"`
	Aggregation     datahub_common.TimeRange_AggregateFunction `mapstructure:"-"`
	MetricValueKeys metricValueKeys                            `mapstructure:"valueKeys"`
}

type predictSeriesMeta struct {
	ScopeStr         string                `mapstructure:"scope"`
	Scope            datahub_schemas.Scope `mapstructure:"-"`
	Category         string                `mapstructure:"category"`
	Type             string                `mapstructure:"type"`
	PredictValueKeys predictValueKeys      `mapstructure:"valueKeys"`
}

type unitValueKeys struct {
	ClusterName          string  `mapstructure:"clusterName"`
	Namespace            string  `mapstructure:"namespace"`
	Name                 string  `mapstructure:"name"`
	ScaleNamespace       string  `mapstructure:"scalerNamespace"`
	ScaleName            string  `mapstructure:"scalerName"`
	ResourceK8SNamespace *string `mapstructure:"resourceK8SNamespace"`
	ResourceK8SName      *string `mapstructure:"resourceK8SName"`
}

type metricValueKeys struct {
	Value string `mapstructure:"value"`
}

type predictValueKeys struct {
	ModelID      string `mapstructure:"modelID"`
	PredictionID string `mapstructure:"predictID"`
	Granularity  string `mapstructure:"granularity"`
	Value        string `mapstructure:"value"`
}
