package keycodes

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	KeycodeMgt "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	ApiKeycode "prophetstor.com/api/datahub/keycodes"
)

func NewKeycode(keycode *KeycodeMgt.Keycode) *ApiKeycode.Keycode {
	k := ApiKeycode.Keycode{}
	k.Keycode = keycode.Keycode
	k.KeycodeType = keycode.KeycodeType
	k.KeycodeVersion = int32(keycode.KeycodeVersion)
	k.Registered = keycode.Registered
	k.LicenseState = keycode.LicenseState
	k.Capacity = NewCapacity(keycode.Capacity)
	k.Functionality = NewFunctionality(keycode.Functionality)
	k.Retention = NewRetention(keycode.Retention)
	k.ServiceAgreement = NewServiceAgreement(keycode.ServiceAgreement)

	if keycode.ApplyTimestamp == 0 {
		k.ApplyTime = &timestamp.Timestamp{Seconds: 0}
	} else if keycode.ApplyTimestamp == -1 {
		k.ApplyTime = &timestamp.Timestamp{Seconds: time.Date(2039, 12, 31, 0, 0, 0, 0, time.UTC).Unix()}
	} else {
		k.ApplyTime = &timestamp.Timestamp{Seconds: keycode.ApplyTimestamp}
	}

	if keycode.ExpireTimestamp == 0 {
		k.ExpireTime = &timestamp.Timestamp{Seconds: 0}
	} else if keycode.ExpireTimestamp == -1 {
		k.ExpireTime = &timestamp.Timestamp{Seconds: time.Date(2039, 12, 31, 0, 0, 0, 0, time.UTC).Unix()}
	} else {
		k.ExpireTime = &timestamp.Timestamp{Seconds: keycode.ExpireTimestamp}
	}

	return &k
}

func NewKeycodeList(keycodes []*KeycodeMgt.Keycode) []*ApiKeycode.Keycode {
	keycodeList := make([]*ApiKeycode.Keycode, 0)

	for _, keycode := range keycodes {
		keycodeList = append(keycodeList, NewKeycode(keycode))
	}

	return keycodeList
}
