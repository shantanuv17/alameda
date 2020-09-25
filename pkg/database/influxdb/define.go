package influxdb

import (
	"prophetstor.com/alameda/pkg/database/common"
)

type Functions = string

// Influxdb function definition
const (
	Last Functions = "LAST"
	Mean Functions = "MEAN"
	Max  Functions = "MAX"
)

var AggregateFuncMap = map[common.AggregateFunction]Functions{
	common.None:        Last,
	common.MaxOverTime: Max,
	common.AvgOverTime: Mean,
}

var FunctionNameMap = map[common.FunctionType]string{
	common.NoneFunction: "",

	// Aggregation Function
	common.FunctionCount:  "COUNT",
	common.FuncDistinct:   "DISTINCT",
	common.FuncIntegral:   "INTEGRAL",
	common.FunctionMean:   "MEAN",
	common.FunctionMedian: "MEDIAN",
	common.FunctionMode:   "MODE",
	common.FunctionSpread: "SPREAD",
	common.FunctionStddev: "STDDEV",
	common.FunctionSum:    "SUM",

	// Selector function
	common.FunctionBottom:     "BOTTOM",
	common.FunctionFirst:      "FIRST",
	common.FunctionLast:       "LAST",
	common.FunctionMax:        "MAX",
	common.FunctionMin:        "MIN",
	common.FunctionPercentile: "PERCENTILE",
	common.FunctionSample:     "SAMPLE",
	common.FunctionTop:        "TOP",

	// Transformation function
	common.FuncDerivative: "DERIVATIVE",
}
