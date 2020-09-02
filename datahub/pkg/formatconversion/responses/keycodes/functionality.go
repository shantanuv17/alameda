package keycodes

import (
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/keycodes"
)

func NewFunctionality(functionality KeycodeMgt.Functionality) *ApiKeycode.Functionality {
	f := ApiKeycode.Functionality{}
	f.DiskProphet = functionality.Diskprophet
	f.Workload = functionality.Workload
	return &f
}
