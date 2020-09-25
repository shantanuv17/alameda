package metrics

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
)

type LicenseConfig struct {
	Enabled  bool                   `mapstructure:"enabled"`
	Capacity *LicenseCapacityConfig `mapstructure:"capacity"`
}

func NewDefaultLicenseConfig() *LicenseConfig {
	var config = LicenseConfig{
		Enabled: true,
		Capacity: &LicenseCapacityConfig{
			Enabled: true,
			CPU: &Notifier{
				Enabled:       true,
				Specs:         DefaultLicenseCapacityCPUSpecs,
				EventInterval: DefaultLicenseCapacityCPUEventInterval,
				EventLevel:    DefaultLicenseCapacityCPUEventLevel,
			},
		},
	}
	return &config
}

func NewLicenseMetrics(licenseCfg *LicenseConfig, influxCfg *influxdb.Config) []AlertInterface {
	notifiers := make([]AlertInterface, 0)
	if licenseCfg.Enabled {
		notifiers = append(notifiers, NewLicenseCapacity(licenseCfg.Capacity, influxCfg)...)
	}
	return notifiers
}
