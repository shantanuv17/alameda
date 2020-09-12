package notifier

import (
	Metrics "github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
)

const (
	DefaultNotifierEnabled = true
)

type Config struct {
	Enabled bool
	Keycode *Metrics.KeycodeConfig `mapstructure:"keycode"`
}

func NewDefaultConfig() *Config {
	var config = Config{
		Enabled: DefaultNotifierEnabled,
		Keycode: Metrics.NewDefaultKeycodeConfig(),
	}
	return &config
}

func (c *Config) Validate() error {
	return nil
}
