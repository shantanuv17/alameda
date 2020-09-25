package common

import (
	"prophetstor.com/api/datahub/common"
)

var FunctionValueMap = map[string]common.FunctionType{
	"average":    common.FunctionType_FUNCTIONTYPE_MEAN,
	"max":        common.FunctionType_FUNCTIONTYPE_MAX,
	"percentile": common.FunctionType_FUNCTIONTYPE_PERCENTILE,
}
