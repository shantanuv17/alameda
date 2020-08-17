package utils

import (
	"fmt"
	"time"

	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/spf13/viper"
)

func GetRowValue(values, columns []string, field string) (string, error) {
	if len(values) != len(columns) {
		return "", fmt.Errorf("number of values and columns are not the same")
	}
	for idx, column := range columns {
		if column == field {
			return values[idx], nil
		}
	}
	return "", fmt.Errorf("no matched field %s found", field)
}

func ReadData(datahubServiceClnt *datahubpkg.Client,
	schemaMeta *datahub_schemas.SchemaMeta, readData []*datahub_data.ReadData) (
	*datahub_data.ReadDataResponse, error) {
	retry := 5
	var retryIntervalSec int64 = 30
	var err error
	if viper.IsSet("datahub.query.retry") {
		retry = viper.GetInt("datahub.query.retry")
	}
	if viper.IsSet("datahub.query.retryInterval") {
		retryIntervalSec = viper.GetInt64("datahub.query.retryInterval")
	}
	for i := 1; i < retry; i++ {
		data, err := datahubServiceClnt.ReadData(
			&datahub_data.ReadDataRequest{
				SchemaMeta: schemaMeta,
				ReadData:   readData,
			})
		if err == nil {
			return data, err
		}
		time.Sleep(time.Duration(retryIntervalSec) * time.Second)
	}
	return nil, err
}

func GetGranularityStr(granularitySec int64) string {
	if granularitySec == 30 {
		return "30s"
	} else if granularitySec == 60 {
		return "1m"
	} else if granularitySec == 3600 {
		return "1h"
	} else if granularitySec == 21600 {
		return "6h"
	} else if granularitySec == 86400 {
		return "24h"
	}
	return "30s"
}

func GetGranularitySec(granularityStr string) int64 {
	if granularityStr == "30s" {
		return 30
	} else if granularityStr == "1m" {
		return 60
	} else if granularityStr == "1h" {
		return 3600
	} else if granularityStr == "6h" {
		return 21600
	} else if granularityStr == "24h" {
		return 86400
	}
	return 30
}
