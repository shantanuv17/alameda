package keycodes

import (
	"time"
)

type KeycodeStatusObject struct {
}

func NewKeycodeStatusObject() *KeycodeStatusObject {
	keycodeStatusObject := KeycodeStatusObject{}
	return &keycodeStatusObject
}

func (c *KeycodeStatusObject) GetStatus() int {
	if c.isNoKeycode() {
		return KeycodeStatusNoKeycode
	}
	if c.isInvalid() {
		return KeycodeStatusInvalid
	}
	if c.isExpired() {
		return KeycodeStatusExpired
	}
	if c.isCapacityCPUCoresExceeded() {
		return KeycodeStatusCapacityCPUCoresExceeded
	}
	if c.isCapacityCPUCoresGracePeriod() {
		return KeycodeStatusCapacityCPUCoresGracePeriod
	}
	if c.isNotActivated() {
		return KeycodeStatusNotActivated
	}
	if c.isValid() {
		return KeycodeStatusValid
	}
	return KeycodeStatusUnknown
}

func (c *KeycodeStatusObject) isNoKeycode() bool {
	if KeycodeList == nil || len(KeycodeList) == 0 {
		return true
	}
	return false
}

func (c *KeycodeStatusObject) isInvalid() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.LicenseState == "Invalid" {
			return true
		}
	}
	return false
}

func (c *KeycodeStatusObject) isExpired() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.LicenseState == "Expired" {
			return true
		}
	}
	return false
}

func (c *KeycodeStatusObject) isNotActivated() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.KeycodeType == "Regular" && KeycodeSummary.LicenseState == "Valid" {
			if KeycodeSummary.Registered == false {
				return true
			}
		}
	}
	return false
}

func (c *KeycodeStatusObject) isValid() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.LicenseState == "Valid" {
			return true
		}
	}
	return false
}

func (c *KeycodeStatusObject) isCapacityCPUCoresGracePeriod() bool {
	// If CPUs is -1 which means no limitation
	if KeycodeSummary.Capacity.CPUs >= 0 {
		if KeycodeCapacityOccupied != nil {
			if KeycodeCapacityOccupied.CPUCores > KeycodeSummary.Capacity.CPUs {
				if KeycodeGracePeriod == 0 {
					return true
				}
				if KeycodeGracePeriod >= time.Now().Unix() {
					return true
				}
			}
		}
	}
	return false
}

func (c *KeycodeStatusObject) isCapacityCPUCoresExceeded() bool {
	// If CPUs is -1 which means no limitation
	if KeycodeSummary.Capacity.CPUs >= 0 {
		if KeycodeCapacityOccupied != nil {
			if KeycodeCapacityOccupied.CPUCores > KeycodeSummary.Capacity.CPUs {
				if KeycodeGracePeriod == 0 {
					return false
				}
				if time.Now().Unix() > KeycodeGracePeriod {
					return true
				}
			}
		}
	}
	return false
}
