package keycodes

import (
	KeycodeMgt "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "prophetstor.com/api/datahub/keycodes"
)

func NewRetention(retention KeycodeMgt.Retention) *ApiKeycode.Retention {
	r := ApiKeycode.Retention{}
	r.ValidMonth = int32(retention.ValidMonth)
	r.Years = int32(retention.Years)
	return &r
}
