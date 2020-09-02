package keycodes

import (
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/keycodes"
)

func NewRetention(retention KeycodeMgt.Retention) *ApiKeycode.Retention {
	r := ApiKeycode.Retention{}
	r.ValidMonth = int32(retention.ValidMonth)
	r.Years = int32(retention.Years)
	return &r
}
