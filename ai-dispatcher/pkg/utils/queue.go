package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

func GetJobID(unit *config.Unit, rowVals []string, columns []string,
	metricType datahub_common.MetricType, granularity int64) (string, error) {

	idVals := []string{}
	for _, idKey := range unit.IDKeys {
		rowVal, err := GetRowValue(rowVals, columns, idKey)
		if err != nil {
			return "", err
		}
		idVals = append(idVals, rowVal)
	}
	unitIDStr := strings.Join(idVals, "/")

	jobID := fmt.Sprintf("%s/%s/%s/%s/%v/%s", unit.Scope, unit.Category, unit.Type, unitIDStr, granularity, metricType)

	return jobID, nil
}

func GetJobMap(rowVals []string, columns []string,
	metricType datahub_common.MetricType, granularity int64) map[string]string {
	jobMap := map[string]string{}
	for idx, col := range columns {
		jobMap[col] = rowVals[idx]
	}
	return jobMap
}

func GetJobStr(rowVals []string, columns []string,
	metricType datahub_common.MetricType, granularity int64) (string, error) {
	jm := GetJobMap(rowVals, columns,
		metricType, granularity)
	jobJSONBin, err := json.Marshal(jm)
	if err != nil {
		return "", err
	}
	return string(jobJSONBin), err
}
