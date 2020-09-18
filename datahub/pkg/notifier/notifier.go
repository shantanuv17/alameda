package notifier

import (
	"github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/robfig/cron"
)

var (
	scope           = log.RegisterScope("notifier", "notifier-mgt", 0)
	MetricsRegistry *Registry
)

func Init(config *Config, influxCfg *influxdb.Config) {
	metrics.Init(influxCfg)

	MetricsRegistry = NewRegistry()
	if config.Enabled {
		MetricsRegistry.Register(metrics.NewKeycodeMetrics(config.Keycode, influxCfg))
		MetricsRegistry.Register(metrics.NewLicenseMetrics(config.License, influxCfg))
	}

	// NOTE: Metering is always enabled for temporary, move to metering package in advance
	MetricsRegistry.Register(metrics.NewMeteringMetrics(config.Metering, influxCfg))
}

func Run() {
	c := cron.New()

	for _, alertMetrics := range MetricsRegistry.GetAll() {
		if alertMetrics.GetEnabled() == true {
			err := c.AddFunc(alertMetrics.GetSpecs(), alertMetrics.Validate)
			if err != nil {
				scope.Errorf("failed to add cron job of %s: %s", alertMetrics.GetName(), err.Error())
			}
		}
	}

	c.Start()

	select {}
}
