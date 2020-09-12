package notifier

import (
	"github.com/containers-ai/alameda/datahub/pkg/notifier/metrics"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/robfig/cron"
)

var (
	scope             = log.RegisterScope("notifier", "notifier-mgt", 0)
	MetricsManagement *Management
)

func NotifierInit(config *Config, influxCfg *influxdb.Config) {
	metrics.Init(influxCfg)

	MetricsManagement = NewManagement()
	if config.Enabled {
		MetricsManagement.AddMetrics(metrics.NewKeycodeMetrics(config.Keycode, influxCfg))
	}
}

func Run() {
	c := cron.New()

	for _, alertMetrics := range MetricsManagement.GetAllMetrics() {
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
