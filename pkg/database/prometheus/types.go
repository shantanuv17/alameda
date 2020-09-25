package prometheus

import (
	DBCommon "prophetstor.com/alameda/pkg/database/common"
)

type nodeMetricsFetchingFunction func(nodeName string, options ...DBCommon.Option) ([]Entity, error)
