package metrics

import (
	"fmt"
	Keycodes "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"math"
	"time"
)

const (
	DefaultLicenseCapacityCPUSpecs         = "0 0 * * * *"
	DefaultLicenseCapacityCPUEventInterval = "14,13,12,11,10,9,8,7,6,5,4,3,2,1,0"
	DefaultLicenseCapacityCPUEventLevel    = "14:Warn,0:Error"
)

const (
	CPUCoresGracePeriod = "Number of CPU cores exceed the licensed number of CPU cores, grace period will expire on %s"
	CPUCoresExceed      = "Number of CPU cores exceed the licensed number of CPU cores"
)

type LicenseCapacityCPU struct {
	AlertMetrics

	days      int
	status    int
	graceTime time.Time
}

func NewLicenseCapacityCPU(notifier *Notifier, influxCfg *influxdb.Config) *LicenseCapacityCPU {
	alert := LicenseCapacityCPU{}
	alert.notifier = notifier
	alert.name = "cpu"
	alert.alertType = "capacity"
	alert.category = "license"
	alert.criteriaType = CriteriaTypeContinuous
	alert.days = 0
	alert.status = Keycodes.KeycodeStatusUnknown
	alert.GenerateCriteria()
	return &alert
}

func (c *LicenseCapacityCPU) Validate() {
	scope.Info("check licensed capacity")
	if c.MeetCriteria() == true {
		if c.eventPosted[c.days] == false {
			c.eventPosted[c.days] = true
			c.PostEvent()
		}
	}
}

func (c *LicenseCapacityCPU) MeetCriteria() bool {
	currentTimestamp := time.Now().Unix()

	// Refresh keycode cache before getting keycode information
	KeycodeMgt.Refresh(true)

	// Get cluster CPU capacity
	keycodes, _, err := KeycodeMgt.GetAllKeycodes()
	if err != nil {
		scope.Error("failed to check criteria when validate license")
		return false
	}

	// If no keycode is applied, we do not need to check the cpu capacity
	if len(keycodes) == 0 {
		c.days = 0
		c.status = Keycodes.KeycodeStatusUnknown
		c.ClearEventPosted()
		return false
	}

	c.status = KeycodeMgt.GetStatus()
	if c.status != Keycodes.KeycodeStatusCapacityCPUCoresGracePeriod && c.status != Keycodes.KeycodeStatusCapacityCPUCoresExceeded {
		c.days = 0
		c.ClearEventPosted()
		return false
	}

	gracePeriod, err := KeycodeMgt.GetKeycodeGracePeriod()
	if err != nil {
		scope.Errorf("failed to get grace period: %s", err.Error())
		return false
	}

	if currentTimestamp >= gracePeriod {
		c.days = 0
		return true
	}

	currentTime := time.Unix(currentTimestamp, 0)
	c.graceTime = time.Unix(gracePeriod, 0)

	diff := c.graceTime.Sub(currentTime)
	c.days = int(math.Floor(diff.Hours() / 24))

	if _, ok := c.eventLevel[c.days]; ok {
		return true
	}

	return false
}

func (c *LicenseCapacityCPU) PostEvent() error {
	eventLevel := c.eventLevel[c.days]
	message := ""

	if c.status == Keycodes.KeycodeStatusCapacityCPUCoresGracePeriod {
		message = fmt.Sprintf(CPUCoresGracePeriod, c.graceTime.Format(time.RFC3339))
	} else {
		message = CPUCoresExceed
	}

	err := Keycodes.PostEvent(eventLevel, message)
	if err != nil {
		scope.Error(err.Error())
		scope.Error("failed to post event when check keycode cpu capacity")
		return err
	}

	return nil
}
