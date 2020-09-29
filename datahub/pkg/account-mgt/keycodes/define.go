package keycodes

import (
	"prophetstor.com/api/datahub/events"
)

const (
	KeycodeStatusUnknown                     = 0
	KeycodeStatusNoKeycode                   = 1
	KeycodeStatusInvalid                     = 2
	KeycodeStatusExpired                     = 3
	KeycodeStatusNotActivated                = 4
	KeycodeStatusValid                       = 5
	KeycodeStatusCapacityCPUCoresGracePeriod = 6
	KeycodeStatusCapacityCPUCoresExceeded    = 7
)

var KeycodeStatusName = map[int]string{
	0: "Unknown",
	1: "No Keycode",
	2: "Invalid",
	3: "Expired",
	4: "Not Activated",
	5: "Valid",
	6: "Grace Period",
	7: "Exceed Licensed Capacity",
}

var KeycodeStatusMessage = map[int]string{
	0: "Unknown keycode is detected",
	1: "No keycode is applied",
	2: "Invalid keycode is detected",
	3: "Keycode is expired",
	4: "Keycode is not activated",
	5: "A valid keycode is applied",
	6: "Number of CPU cores exceeds the licensed number of CPU cores, will expire in grace period",
	7: "Number of CPU cores exceeds the licensed number of CPU cores",
}

var KeycodeEventLevelMap = map[int]events.EventLevel{
	KeycodeStatusUnknown:                     events.EventLevel_EVENT_LEVEL_ERROR,
	KeycodeStatusNoKeycode:                   events.EventLevel_EVENT_LEVEL_ERROR,
	KeycodeStatusInvalid:                     events.EventLevel_EVENT_LEVEL_ERROR,
	KeycodeStatusExpired:                     events.EventLevel_EVENT_LEVEL_ERROR,
	KeycodeStatusNotActivated:                events.EventLevel_EVENT_LEVEL_WARNING,
	KeycodeStatusValid:                       events.EventLevel_EVENT_LEVEL_INFO,
	KeycodeStatusCapacityCPUCoresGracePeriod: events.EventLevel_EVENT_LEVEL_WARNING,
	KeycodeStatusCapacityCPUCoresExceeded:    events.EventLevel_EVENT_LEVEL_ERROR,
}

var KeycodeInfluxTargetMap = map[int]string{
	KeycodeStatusUnknown:                     "Summary",
	KeycodeStatusNoKeycode:                   "N/A",
	KeycodeStatusInvalid:                     "Summary",
	KeycodeStatusExpired:                     "Summary",
	KeycodeStatusNotActivated:                "Summary",
	KeycodeStatusValid:                       "Summary",
	KeycodeStatusCapacityCPUCoresGracePeriod: "Summary",
	KeycodeStatusCapacityCPUCoresExceeded:    "Summary",
}
