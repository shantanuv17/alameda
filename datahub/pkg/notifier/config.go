package notifier

import (
	"github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
)

const (
	DefaultNotifierEnabled = true
)

type Config struct {
	Enabled  bool
	Keycode  *metrics.KeycodeConfig  `mapstructure:"keycode"`
	Metering *metrics.MeteringConfig `mapstructure:"metering"`
}

func NewDefaultConfig() *Config {
	var config = Config{
		Enabled:  DefaultNotifierEnabled,
		Keycode:  metrics.NewDefaultKeycodeConfig(),
		Metering: metrics.NewDefaultMeteringConfig(),
	}
	return &config
}

func (c *Config) Validate() error {
	return nil
}
