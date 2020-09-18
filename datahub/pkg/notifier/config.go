package notifier

import (
	"github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
)

type Config struct {
	Enabled  bool
	Keycode  *metrics.KeycodeConfig  `mapstructure:"keycode"`
	License  *metrics.LicenseConfig  `mapstructure:"license"`
	Metering *metrics.MeteringConfig `mapstructure:"metering"`
}

func NewDefaultConfig() *Config {
	var config = Config{
		Enabled:  true,
		Keycode:  metrics.NewDefaultKeycodeConfig(),
		License:  metrics.NewDefaultLicenseConfig(),
		Metering: metrics.NewDefaultMeteringConfig(),
	}
	return &config
}

func (c *Config) Validate() error {
	return nil
}
