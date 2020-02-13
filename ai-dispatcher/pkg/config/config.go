package config

import (
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

type Config struct {
	Units      []Unit                               `mapstructure:"units"`
	MetricType map[string]datahub_common.MetricType `mapstructure:"metricType"`
	Scope      map[string]datahub_schemas.Scope     `mapstructure:"scope"`
}

func (cfg *Config) Init() {
	for jobIdx := range cfg.Units {
		if val, ok := cfg.Scope[cfg.Units[jobIdx].ScopeStr]; ok {
			cfg.Units[jobIdx].Scope = val
		}

		for _, metric := range cfg.Units[jobIdx].MetricTypeStrs {
			if val, ok := cfg.MetricType[metric]; ok {
				cfg.Units[jobIdx].MetricTypes = append(cfg.Units[jobIdx].MetricTypes, val)
			}
		}

		if val, ok := cfg.Scope[cfg.Units[jobIdx].Metric.ScopeStr]; ok {
			cfg.Units[jobIdx].Metric.Scope = val
		}
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
	MetricTypes    []datahub_common.MetricType `mapstructure:"-"`
}

type metricSeriesMeta struct {
	ScopeStr        string                `mapstructure:"scope"`
	Scope           datahub_schemas.Scope `mapstructure:"-"`
	Category        string                `mapstructure:"category"`
	Type            string                `mapstructure:"type"`
	MetricValueKeys metricValueKeys       `mapstructure:"valueKeys"`
}

type predictSeriesMeta struct {
	ScopeStr         string                `mapstructure:"scope"`
	Scope            datahub_schemas.Scope `mapstructure:"-"`
	Category         string                `mapstructure:"category"`
	Type             string                `mapstructure:"type"`
	PredictValueKeys predictValueKeys      `mapstructure:"valueKeys"`
}

type metricValueKeys struct {
	Value string `mapstructure:"value"`
}

type predictValueKeys struct {
	ModelID      string `mapstructure:"modelID"`
	PredictionID string `mapstructure:"predictID"`
	Value        string `mapstructure:"value"`
}
