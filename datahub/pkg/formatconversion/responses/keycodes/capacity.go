package keycodes

import (
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/keycodes"
)

func NewCapacity(capacity KeycodeMgt.Capacity) *ApiKeycode.Capacity {
	c := ApiKeycode.Capacity{}
	c.Users = int32(capacity.Users)
	c.Hosts = int32(capacity.Hosts)
	c.Disks = int32(capacity.Disks)
	return &c
}
