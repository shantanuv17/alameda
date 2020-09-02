package keycodes

import (
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/keycodes"
)

func NewServiceAgreement(serviceAgreement KeycodeMgt.ServiceAgreement) *ApiKeycode.ServiceAgreement {
	s := ApiKeycode.ServiceAgreement{}
	return &s
}
