package common

import (
	Common "prophetstor.com/api/datahub/common"
	"time"
)

// Order enumerator
type Order = int

// Aggregate function enumerator
type AggregateFunction = int

// Function type enumerator
type FunctionType = int

// Data type enumerator
type DataType int

// Sort order definition
const (
	NoneOrder Order = iota
	Asc             // Represent ascending order
	Desc            // Represent descending order
)

// Aggregation function definition
const (
	None        AggregateFunction = 0
	MaxOverTime AggregateFunction = 1
	AvgOverTime AggregateFunction = 2
)

const (
	NoneFunction FunctionType = iota

	// Aggregation Function
	FunctionCount  // Returns the number of non-null field values
	FuncDistinct   // Returns the list of unique field values
	FuncIntegral   // Returns the area under the curve for subsequent field values
	FunctionMean   // Returns the arithmetic mean (average) of field values
	FunctionMedian // Returns the middle value from a sorted list of field values
	FunctionMode   // Returns the most frequent value in a list of field values
	FunctionSpread // Returns the difference between the minimum and maximum field values
	FunctionStddev // Returns the standard deviation of field values
	FunctionSum    // Returns the sum of field values

	// Selector function
	FunctionBottom     // Returns the smallest N field values
	FunctionFirst      // Returns the field value with the oldest timestamp
	FunctionLast       // Returns the field value with the most recent timestamp
	FunctionMax        // Returns the greatest field value
	FunctionMin        // Returns the lowest field value
	FunctionPercentile // Returns the Nth percentile field value
	FunctionSample     // Returns a random sample of N field values. SAMPLE() uses reservoir sampling to generate the random points
	FunctionTop        // Returns the greatest N field values

	// Transformation function
	FuncDerivative // Returns the rate of change between subsequent field values
)

// Data type definition
const (
	Invalid DataType = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
)

type Row struct {
	Time   *time.Time
	Values []string
}

type Group struct {
	Columns []string
	Rows    []*Row
}

var (
	AggregationOverTime = map[AggregateFunction]string{
		MaxOverTime: "max_over_time",
		AvgOverTime: "avg_over_time",
	}

	TimeRange2AggregationOverTime = map[Common.TimeRange_AggregateFunction]AggregateFunction{
		Common.TimeRange_NONE: None,
		Common.TimeRange_MAX:  MaxOverTime,
		Common.TimeRange_AVG:  AvgOverTime,
	}
	FunctionTypeMap = map[AggregateFunction]FunctionType{
		None:        NoneFunction,
		MaxOverTime: FunctionMax,
		AvgOverTime: FunctionMean,
	}
)
