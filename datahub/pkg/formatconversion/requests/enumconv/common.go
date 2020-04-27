package enumconv

import (
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var AggregateFunctionNameMap = map[ApiCommon.TimeRange_AggregateFunction]DBCommon.AggregateFunction{
	ApiCommon.TimeRange_NONE: DBCommon.None,
	ApiCommon.TimeRange_MAX:  DBCommon.MaxOverTime,
	ApiCommon.TimeRange_AVG:  DBCommon.AvgOverTime,
}

var DataTypeNameMap = map[ApiCommon.DataType]DBCommon.DataType{
	ApiCommon.DataType_DATATYPE_UNDEFINED: DBCommon.Invalid,
	ApiCommon.DataType_DATATYPE_BOOL:      DBCommon.Bool,
	ApiCommon.DataType_DATATYPE_INT:       DBCommon.Int,
	ApiCommon.DataType_DATATYPE_INT8:      DBCommon.Int8,
	ApiCommon.DataType_DATATYPE_INT16:     DBCommon.Int16,
	ApiCommon.DataType_DATATYPE_INT32:     DBCommon.Int32,
	ApiCommon.DataType_DATATYPE_INT64:     DBCommon.Int64,
	ApiCommon.DataType_DATATYPE_UINT:      DBCommon.Uint,
	ApiCommon.DataType_DATATYPE_UINT8:     DBCommon.Uint8,
	ApiCommon.DataType_DATATYPE_UINT16:    DBCommon.Uint16,
	ApiCommon.DataType_DATATYPE_UINT32:    DBCommon.Uint32,
	ApiCommon.DataType_DATATYPE_UTIN64:    DBCommon.Uint64,
	ApiCommon.DataType_DATATYPE_FLOAT32:   DBCommon.Float32,
	ApiCommon.DataType_DATATYPE_FLOAT64:   DBCommon.Float64,
	ApiCommon.DataType_DATATYPE_STRING:    DBCommon.String,
}

var QueryConditionOrderNameMap = map[ApiCommon.QueryCondition_Order]DBCommon.Order{
	ApiCommon.QueryCondition_NONE: DBCommon.NoneOrder,
	ApiCommon.QueryCondition_ASC:  DBCommon.Asc,
	ApiCommon.QueryCondition_DESC: DBCommon.Desc,
}

var ResourceBoundaryNameMap = map[ApiCommon.ResourceBoundary]schemas.ResourceBoundary{
	ApiCommon.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED: schemas.ResourceBoundaryUndefined,
	ApiCommon.ResourceBoundary_RESOURCE_RAW:                schemas.ResourceRaw,
	ApiCommon.ResourceBoundary_RESOURCE_UPPER_BOUND:        schemas.ResourceUpperBound,
	ApiCommon.ResourceBoundary_RESOURCE_LOWER_BOUND:        schemas.ResourceLowerBound,
}

var ResourceQuotaNameMap = map[ApiCommon.ResourceQuota]schemas.ResourceQuota{
	ApiCommon.ResourceQuota_RESOURCE_QUOTA_UNDEFINED: schemas.ResourceQuotaUndefined,
	ApiCommon.ResourceQuota_RESOURCE_LIMIT:           schemas.ResourceLimit,
	ApiCommon.ResourceQuota_RESOURCE_REQUEST:         schemas.ResourceRequest,
	ApiCommon.ResourceQuota_RESOURCE_INITIAL_LIMIT:   schemas.ResourceInitialLimit,
	ApiCommon.ResourceQuota_RESOURCE_INITIAL_REQUEST: schemas.ResourceInitialRequest,
}
