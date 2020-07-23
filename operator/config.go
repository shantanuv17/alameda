package main

import (
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	"github.com/containers-ai/alameda/operator/datahub"
	"github.com/containers-ai/alameda/pkg/database/prometheus"
	"github.com/containers-ai/alameda/pkg/utils/log"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Config defines configurations
type Config struct {
	Log     *log.Config     `mapstructure:"log"`
	Datahub *datahub.Config `mapstructure:"datahub"`
	Manager manager.Manager

	Kafka      *kafka.Config `mapstructure:"kafka"`
	Prometheus *Prometheus   `mapstructure:"prometheus"`
}

// NewConfig returns Config objecdt
func NewConfig(manager manager.Manager) Config {

	c := Config{
		Manager: manager,
	}
	c.init()

	return c
}

func NewConfigWithoutMgr() Config {

	c := Config{}
	c.init()

	return c
}

func (c *Config) init() {

	defaultLogConfig := log.NewDefaultConfig()

	c.Log = &defaultLogConfig
	c.Datahub = datahub.NewConfig()

	// TODO: If need to new kafka config
	prometheus := NewDefaultPrometheus()
	c.Prometheus = &prometheus
}

func (c Config) Validate() error {

	return nil
}

// Prometheus wraps Prometheus config and add addition configurations for Alameda Operator
type Prometheus struct {
	prometheus.Config `mapstructure:",squash"`
	RequiredMetrics   []string `mapstructure:"requiredMetrics"`
}

func NewDefaultPrometheus() Prometheus {
	c := prometheus.NewDefaultConfig()
	return Prometheus{
		Config:          *c,
		RequiredMetrics: []string{},
	}
}
