package clusterstatus

import (
	"github.com/containers-ai/alameda/pkg/database/influxdb"
)

const (
	Keycode influxdb.Tag = "keycode"

	KeycodeStatus          influxdb.Field = "status"
	KeycodeType            influxdb.Field = "type"
	KeycodeState           influxdb.Field = "state"
	KeycodeRegistered      influxdb.Field = "registered"
	KeycodeExpireTimestamp influxdb.Field = "expire_timestamp"
	KeycodeRawdata         influxdb.Field = "rawdata"
)

var (
	// List of tags of keycode measurement
	KeycodeTags = []influxdb.Tag{
		Keycode,
	}
	// List of fields of keycode measurement
	KeycodeFields = []influxdb.Field{
		KeycodeStatus,
		KeycodeType,
		KeycodeState,
		KeycodeRegistered,
		KeycodeExpireTimestamp,
		KeycodeRawdata,
	}
)
