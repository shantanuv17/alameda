package common

import (
	"github.com/golang/protobuf/ptypes"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/enumconv"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	ApiCommon "prophetstor.com/api/datahub/common"
)

func NewQueryCondition(queryCondition *ApiCommon.QueryCondition) *DBCommon.QueryCondition {
	if queryCondition != nil {
		qc := DBCommon.QueryCondition{}
		qc.TimestampOrder = enumconv.QueryConditionOrderNameMap[queryCondition.GetOrder()]
		qc.WhereClause = queryCondition.GetWhereClause()
		qc.WhereCondition = NewWhereCondition(queryCondition.GetWhereCondition())
		qc.Selects = queryCondition.GetSelects()
		qc.Groups = queryCondition.GetGroups()
		qc.Limit = int(queryCondition.GetLimit())
		qc.Function = NewFunction(queryCondition.GetFunction())
		qc.Into = NewInto(queryCondition.GetInto())
		if queryCondition.GetTimeRange() != nil {
			timeRange := queryCondition.GetTimeRange()
			if timeRange.GetStartTime() != nil {
				ts, _ := ptypes.Timestamp(timeRange.GetStartTime())
				qc.StartTime = &ts
			}
			if timeRange.GetEndTime() != nil {
				ts, _ := ptypes.Timestamp(timeRange.GetEndTime())
				qc.EndTime = &ts
			}
			if timeRange.GetStep() != nil {
				ts, _ := ptypes.Duration(timeRange.GetStep())
				qc.StepTime = &ts
			}
			qc.AggregateOverTimeFunction = enumconv.AggregateFunctionNameMap[timeRange.GetAggregateFunction()]
		}
		if queryCondition.GetSubQuery() != nil {
			qc.SubQuery = NewQueryCondition(queryCondition.GetSubQuery())
		}
		return &qc
	}
	return nil
}
