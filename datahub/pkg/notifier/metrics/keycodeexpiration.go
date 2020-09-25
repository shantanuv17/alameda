package metrics

import (
	"fmt"
	"math"
	Keycodes "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	"prophetstor.com/alameda/pkg/database/influxdb"
	"time"
)

const (
	DefaultKeycodeExpirationSpecs         = "0 0 * * * *"
	DefaultKeycodeExpirationEventInterval = "90,60,30,15,7,6,5,4,3,2,1,0,-1,-2,-3,-4,-5,-6,-7"
	DefaultKeycodeExpirationEventLevel    = "90:Info,15:Warn,0:Error"
)

type KeycodeExpiration struct {
	AlertMetrics

	days    int
	expired bool
}

func NewKeycodeExpiration(notifier *Notifier, influxCfg *influxdb.Config) *KeycodeExpiration {
	alert := KeycodeExpiration{}
	alert.notifier = notifier
	alert.name = "expiration"
	alert.alertType = "expiration"
	alert.category = "keycode"
	alert.criteriaType = CriteriaTypeContinuous
	alert.days = 0
	alert.expired = false
	alert.GenerateCriteria()
	return &alert
}

func (c *KeycodeExpiration) Validate() {
	scope.Info("check if keycode is expired")
	if c.MeetCriteria() == true {
		if c.eventPosted[c.days] == false {
			c.eventPosted[c.days] = true
			c.PostEvent()
		}
	}
}

func (c *KeycodeExpiration) MeetCriteria() bool {
	currentTimestamp := time.Now().Unix()

	// Refresh keycode cache before getting keycode information
	KeycodeMgt.Refresh(true)

	keycodes, summary, err := KeycodeMgt.GetAllKeycodes()
	if err != nil {
		scope.Error("failed to check criteria when validate license")
		return false
	}

	// If no keycode is applied, we do not need to check the remaining days
	if len(keycodes) == 0 {
		c.days = 0
		c.expired = false
		c.ClearEventPosted()
		return false
	}

	// Check if license is already expired
	isExpired := false
	if currentTimestamp >= summary.ExpireTimestamp {
		isExpired = true
	}

	// If expired status is changed from "true" to "false", we need to clear eventPosted map
	if c.expired != isExpired {
		c.expired = isExpired
		if isExpired == false {
			c.ClearEventPosted()
		}
	}

	currentTime := time.Unix(currentTimestamp, 0)
	expireTime := time.Unix(summary.ExpireTimestamp, 0)

	diff := expireTime.Sub(currentTime)
	c.days = int(math.Floor(diff.Hours() / 24))

	if _, ok := c.eventLevel[c.days]; ok {
		return true
	}

	return false
}

func (c *KeycodeExpiration) PostEvent() error {
	eventLevel := c.eventLevel[c.days]

	days := 0
	if c.days >= 0 {
		days = c.days
	} else {
		days = 0
	}

	message := fmt.Sprintf("Expired in %d day(s).", days)

	err := Keycodes.PostEvent(eventLevel, message)
	if err != nil {
		scope.Error(err.Error())
		scope.Error("failed to post event when validate keycode")
		return err
	}

	return nil
}
