package keycodes

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

func (c *KeycodeStatusObject) isCapacityCPUCoresExceeded() bool {
	// If CPUs is -1 which means no limitation
	if KeycodeSummary.Capacity.CPUs >= 0 {
		if CPUCoresOccupied > KeycodeSummary.Capacity.CPUs {
			return true
		}
	}
	return false
}
