package prometheus

import (
	"fmt"
	"github.com/pkg/errors"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
)

func WrapQueryExpression(queryExpression string, aggregateFunc DBCommon.AggregateFunction, aggregateOverSeconds int64) (string, error) {
	if aggregateFunc == DBCommon.None {
		return queryExpression, nil
	}

	if funcName, exist := DBCommon.AggregationOverTime[aggregateFunc]; !exist {
		scope.Errorf("no mapping function for aggregation: %d", aggregateFunc)
		scope.Error("failed to wrap prometheus query expression with aggregation")
		return queryExpression, errors.Errorf("no mapping function for aggregation: %d", aggregateFunc)
	} else {
		queryExpression = fmt.Sprintf("%s(%s[%ds])", funcName, queryExpression, aggregateOverSeconds)
	}

	return queryExpression, nil
}
