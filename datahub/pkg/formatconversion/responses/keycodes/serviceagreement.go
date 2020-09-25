package keycodes

import (
	KeycodeMgt "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "prophetstor.com/api/datahub/keycodes"
)

func NewServiceAgreement(serviceAgreement KeycodeMgt.ServiceAgreement) *ApiKeycode.ServiceAgreement {
	s := ApiKeycode.ServiceAgreement{}
	return &s
}
