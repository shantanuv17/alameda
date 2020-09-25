package metrics

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
)

type LicenseCapacityConfig struct {
	Enabled bool      `mapstructure:"enabled"`
	CPU     *Notifier `mapstructure:"cpu"`
}

func NewLicenseCapacity(capacityCfg *LicenseCapacityConfig, influxCfg *influxdb.Config) []AlertInterface {
	notifiers := make([]AlertInterface, 0)
	if capacityCfg.Enabled {
		notifiers = append(notifiers, NewLicenseCapacityCPU(capacityCfg.CPU, influxCfg))
	}
	return notifiers
}
