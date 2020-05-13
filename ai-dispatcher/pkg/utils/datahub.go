package utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
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

func ReadData(datahubServiceClnt datahub_v1alpha1.DatahubServiceClient,
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
		data, err := datahubServiceClnt.ReadData(context.Background(),
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
	} else if granularitySec%60 == 0 &&
		granularitySec/60 > 0 && granularitySec/60 < 60 {
		return fmt.Sprintf("%dm", granularitySec/60)
	} else if granularitySec%3600 == 0 &&
		granularitySec/3600 > 0 && granularitySec/3600 <= 24 {
		return fmt.Sprintf("%dh", granularitySec/3600)
	}
	return "30s"
}

func GetGranularitySec(granularityStr string) int64 {
	theUnit := granularityStr[len(granularityStr)-1:]
	valStr := granularityStr[0 : len(granularityStr)-1]
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 30
	}
	if theUnit == "s" {
		return val
	} else if theUnit == "m" {
		return val * 60
	} else if theUnit == "h" {
		return val * 3600
	}

	return 30
}
