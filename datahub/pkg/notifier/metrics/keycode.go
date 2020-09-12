package metrics

import (
	"github.com/containers-ai/alameda/pkg/database/influxdb"
)

type KeycodeCapacity struct {
	CPU *Notifier `mapstructure:"cpu"`
}

type KeycodeConfig struct {
	Enabled    bool             `mapstructure:"enabled"`
	Expiration *Notifier        `mapstructure:"expiration"`
	Capacity   *KeycodeCapacity `mapstructure:"capacity"`
}

func NewDefaultKeycodeConfig() *KeycodeConfig {
	var config = KeycodeConfig{
		Enabled: true,
		Expiration: &Notifier{
			Enabled:       DefaultKeycodeExpirationEnabled,
			Specs:         DefaultKeycodeExpirationSpecs,
			EventInterval: DefaultKeycodeExpirationEventInterval,
			EventLevel:    DefaultKeycodeExpirationEventLevel,
		},
		Capacity: &KeycodeCapacity{
			CPU: &Notifier{
				Enabled:       DefaultKeycodeCapacityCPUEnabled,
				Specs:         DefaultKeycodeCapacityCPUSpecs,
				EventInterval: DefaultKeycodeCapacityCPUEventInterval,
				EventLevel:    DefaultKeycodeCapacityCPUEventLevel,
			},
		},
	}
	return &config
}

func NewKeycodeMetrics(keycodeCfg *KeycodeConfig, influxCfg *influxdb.Config) []AlertInterface {
	notifiers := make([]AlertInterface, 0)
	if keycodeCfg.Enabled {
		notifiers = append(notifiers, NewKeycodeExpiration(keycodeCfg.Expiration, influxCfg))
		notifiers = append(notifiers, NewKeycodeCapacityCPU(keycodeCfg.Capacity.CPU, influxCfg))
	}
	return notifiers
}
