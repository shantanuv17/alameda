package metrics

import (
	"github.com/containers-ai/alameda/pkg/database/influxdb"
)

type KeycodeConfig struct {
	Enabled    bool      `mapstructure:"enabled"`
	Expiration *Notifier `mapstructure:"expiration"`
}

func NewDefaultKeycodeConfig() *KeycodeConfig {
	var config = KeycodeConfig{
		Enabled: true,
		Expiration: &Notifier{
			Enabled:       true,
			Specs:         DefaultKeycodeExpirationSpecs,
			EventInterval: DefaultKeycodeExpirationEventInterval,
			EventLevel:    DefaultKeycodeExpirationEventLevel,
		},
	}
	return &config
}

func NewKeycodeMetrics(keycodeCfg *KeycodeConfig, influxCfg *influxdb.Config) []AlertInterface {
	notifiers := make([]AlertInterface, 0)
	if keycodeCfg.Enabled {
		notifiers = append(notifiers, NewKeycodeExpiration(keycodeCfg.Expiration, influxCfg))
	}
	return notifiers
}
