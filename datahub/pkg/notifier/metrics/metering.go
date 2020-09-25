package metrics

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
)

type MeteringConfig struct {
	Enabled     bool      `mapstructure:"enabled"`
	Federatorai *Notifier `mapstructure:"federatorai"`
}

func NewDefaultMeteringConfig() *MeteringConfig {
	var config = MeteringConfig{
		Enabled: true,
		Federatorai: &Notifier{
			Enabled: DefaultMeteringFederatoraiEnabled,
			Specs:   DefaultMeteringFederatoraiSpecs,
		},
	}
	return &config
}

func NewMeteringMetrics(meteringCfg *MeteringConfig, influxCfg *influxdb.Config) []AlertInterface {
	notifiers := make([]AlertInterface, 0)
	if meteringCfg.Enabled {
		notifiers = append(notifiers, NewMeteringFederatorai(meteringCfg.Federatorai, influxCfg))
	}
	return notifiers
}
