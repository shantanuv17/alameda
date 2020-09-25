package probe

import (
	"os"

	"github.com/spf13/viper"
	"prophetstor.com/alameda/pkg/utils/log"
)

var scope = log.RegisterScope("probe", "ai dispatcher health probe", 0)

func LivenessProbe(cfg *LivenessProbeConfig) {
	queueURL := viper.GetString("rabbitmq.url")
	err := CheckRBMQEventQueueAccess(queueURL)
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func ReadinessProbe(cfg *ReadinessProbeConfig) {
	queueURL := viper.GetString("rabbitmq.url")
	err := connQueue(queueURL)
	if err != nil {
		scope.Errorf("Readiness probe: query queue with url (%s) failed due to %s", queueURL, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
