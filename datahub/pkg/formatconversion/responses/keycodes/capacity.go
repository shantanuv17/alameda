package keycodes

import (
	KeycodeMgt "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "prophetstor.com/api/datahub/keycodes"
)

func NewCapacity(capacity KeycodeMgt.Capacity) *ApiKeycode.Capacity {
	c := ApiKeycode.Capacity{}
	c.Users = int32(capacity.Users)
	c.Hosts = int32(capacity.Hosts)
	c.Disks = int32(capacity.Disks)
	return &c
}
