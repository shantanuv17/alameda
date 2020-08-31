package influxdb

import (
	"github.com/containers-ai/alameda/pkg/database/common"
	"github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	"time"
)

type InfluxData struct {
	Time   *time.Time
	Tags   map[string]string
	Fields map[string]string
}

func NewInfluxData(columnNames []string, row *common.Row, columns []*schemas.Column) *InfluxData {
	record := InfluxData{Time: row.Time}
	record.Tags = make(map[string]string)
	record.Fields = make(map[string]string)

	for i, name := range columnNames {
		for _, column := range columns {
			if name == column.Name {
				if column.ColumnType == schemas.Tag {
					record.Tags[name] = row.Values[i]
				} else {
					record.Fields[name] = row.Values[i]
				}
				break
			}
		}
	}

	return &record
}

func CompareInfluxDataByTags(data1, data2 *InfluxData) bool {
	// NOTE: compare the second data MAINLY with the FIRST data
	for k, v := range data1.Tags {
		if val, ok := data2.Tags[k]; !ok {
			if v != "" {
				return false
			}
		} else {
			if v != val {
				return false
			}
		}
	}
	return true
}

func GenerateQueryConditionByInfluxData(records []*InfluxData) *common.QueryCondition {
	queryCondition := common.QueryCondition{}
	for _, record := range records {
		condition := common.Condition{}
		for k, v := range record.Tags {
			condition.Keys = append(condition.Keys, k)
			condition.Values = append(condition.Values, v)
			condition.Operators = append(condition.Operators, "=")
		}
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, &condition)
	}
	return &queryCondition
}
