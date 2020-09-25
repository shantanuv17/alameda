package keycodes

import (
	KeycodeMgt "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "prophetstor.com/api/datahub/keycodes"
)

func NewFunctionality(functionality KeycodeMgt.Functionality) *ApiKeycode.Functionality {
	f := ApiKeycode.Functionality{}
	f.DiskProphet = functionality.Diskprophet
	f.Workload = functionality.Workload
	return &f
}
