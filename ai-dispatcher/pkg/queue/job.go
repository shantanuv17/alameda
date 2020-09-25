package queue

import (
	"encoding/json"
	"time"

	"github.com/streadway/amqp"
	"prophetstor.com/alameda/ai-dispatcher/consts"
	utils "prophetstor.com/alameda/ai-dispatcher/pkg/utils"
	datahub_common "prophetstor.com/api/datahub/common"
)

type job struct {
	UnitType          string                    `json:"unitType"`
	Granularity       string                    `json:"granularity"`
	GranularitySec    int64                     `json:"granularitySec"`
	PayloadJSONString string                    `json:"payloadJSONString"`
	CreateTimeStamp   int64                     `json:"createTimestamp"`
	ClusterID         string                    `json:"clusterID"`
	MetricType        datahub_common.MetricType `json:"metricType"`
	MetricTypeString  string                    `json:"metricTypeString"`
	ContainerName     string                    `json:"containerName"`
}

type jobBuilder struct {
	job *job
}

func NewJobBuilder(clusterID string, unitType string, granularitySec int64,
	metricType datahub_common.MetricType, payloadJSONString string,
	extraJobInfo map[string]string) *jobBuilder {
	granularity := GetGranularityStr(granularitySec)
	job := &job{
		UnitType:          unitType,
		GranularitySec:    granularitySec,
		Granularity:       granularity,
		PayloadJSONString: payloadJSONString,
		CreateTimeStamp:   time.Now().Unix(),
		ClusterID:         clusterID,
		MetricType:        metricType,
		MetricTypeString:  metricType.String(),
	}
	if unitType == consts.UnitTypePod && extraJobInfo != nil {
		ctName, ok := extraJobInfo["containerName"]
		if ok {
			job.ContainerName = ctName
		}
	}
	return &jobBuilder{job: job}
}

func (jobBuilder *jobBuilder) GetJobJSONString() (string, error) {
	jobJSONBin, err := json.Marshal(jobBuilder.job)
	if err != nil {
		return "", err
	}
	return string(jobJSONBin), err
}

func GetGranularityStr(granularitySec int64) string {
	return utils.GetGranularityStr(granularitySec)
}

func GetGranularitySec(granularityStr string) int64 {
	return utils.GetGranularitySec(granularityStr)
}

func GetQueueConn(queueURL string, retryItvMS int64) *amqp.Connection {
	for {
		queueConn, err := amqp.Dial(queueURL)
		if err != nil {
			scope.Errorf("Queue connection constructs failed and will retry after %v milliseconds. %s", retryItvMS, err.Error())
			time.Sleep(time.Duration(retryItvMS) * time.Millisecond)
			continue
		}
		return queueConn
	}
}
