package queue

import (
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

type QueueSender interface {
	SendJsonString(queueName, jsonStr, msgID string, timeout int64) error
	SendJob(queueName string, unit *config.Unit,
		columns, values []string, metricType datahub_common.MetricType, granularity int64) error
	getRetry() *retry
}

var scope = log.RegisterScope("queue", "job queue", 0)

const (
	DEFAULT_PUBLISH_RETRY_TIME              = 3
	DEFAULT_PUBLISH_RETRY_INTERVAL_MS int64 = 500
	DEFAULT_CONSUME_RETRY_TIME              = 3
	DEFAULT_CONSUME_RETRY_INTERVAL_MS int64 = 500
	DEFAULT_ACK_TIMEOUT_SEC                 = 3
)
