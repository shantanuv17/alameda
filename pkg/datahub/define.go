package datahub

// Order enumerator
type Order = int

// Function type enumerator
type FunctionType = int

// Sort order definition
const (
	NoneOrder Order = iota
	Asc  // Represent ascending order
	Desc // Represent descending order
)

const (
	NoneFunction FunctionType = iota

	// Aggregation Function
	FunctionCount  // Returns the number of non-null field values
    FunctionMean   // Returns the arithmetic mean (average) of field values.
    FunctionMedian // Returns the middle value from a sorted list of field values.
    FunctionStddev // Returns the standard deviation of field values.
    FunctionSum    // Returns the sum of field values.

    // Selector function
    FunctionBottom // Returns the smallest N field values.
    FunctionFirst  // Returns the field value with the oldest timestamp.
    FunctionLast   // Returns the field value with the most recent timestamp.
    FunctionMax    // Returns the greatest field value.
    FunctionMin    // Returns the lowest field value.
    FunctionTop    // Returns the greatest N field values.
)
