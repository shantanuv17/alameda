package common

import (
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/enumconv"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

func NewFunction(function *ApiCommon.Function) *common.Function {
	if function != nil {
		f := common.Function{}
		f.Type = enumconv.QueryConditionFunctionNameMap[function.GetType()]
		f.Fields = function.GetFields()
		f.Tags = function.GetTags()
		f.IntoClause = function.GetIntoClause()
		f.RegularExpression = function.GetRegularExpression()
		f.Number = function.GetNumber()
		return &f
	}
	return nil
}

func NewWhereCondition(conditions []*ApiCommon.Condition) []*common.Condition {
	if conditions != nil {
		whereCondition := make([]*common.Condition, 0)
		for _, condition := range conditions {
			whereCondition = append(whereCondition, NewCondition(condition))
		}
		return whereCondition
	}
	return nil
}

func NewCondition(condition *ApiCommon.Condition) *common.Condition {
	if condition != nil {
		c := common.Condition{}
		c.Keys = condition.GetKeys()
		c.Values = condition.GetValues()
		c.Operators = condition.GetOperators()
		c.Types = make([]common.DataType, 0)
		return &c
	}
	return nil
}
