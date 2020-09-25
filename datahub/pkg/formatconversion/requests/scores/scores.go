package scores

import (
	DaoScoreTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/scores/types"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/common"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	ApiScores "prophetstor.com/api/datahub/scores"
)

type ListSimulatedSchedulingScoresRequestExtended struct {
	Request *ApiScores.ListSimulatedSchedulingScoresRequest
}

func (r *ListSimulatedSchedulingScoresRequestExtended) ProduceRequest() DaoScoreTypes.ListRequest {
	var (
		queryCondition DBCommon.QueryCondition
	)

	queryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	listRequest := DaoScoreTypes.ListRequest{
		QueryCondition: queryCondition,
	}

	return listRequest
}
