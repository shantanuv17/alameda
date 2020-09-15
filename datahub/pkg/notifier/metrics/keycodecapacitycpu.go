package metrics

import (
	"fmt"
	Keycodes "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	"strconv"
	"strings"
)

const (
	DefaultKeycodeCapacityCPUEnabled       = true
	DefaultKeycodeCapacityCPUSpecs         = "0 0 * * * *"
	DefaultKeycodeCapacityCPUEventInterval = "50,70,80,90,100"
	DefaultKeycodeCapacityCPUEventLevel    = "Info,Info,Warn,Warn,Error"
)

const (
	CPUCoresExceed          = "CPU cores (%d) exceed the licensed number of CPU cores (%d)"
	CPUCoresPercentageUsage = "CPU cores utilization exceeds %d percents"
)

type KeycodeCapacityCPU struct {
	AlertMetrics

	eventLevel      map[int]events.EventLevel
	eventPosted     map[int]bool
	percentageLevel []int
	percentage      int
	percentageMeet  int
	coreOccupied    int64
	coreCapacity    int64
	influxCfg       *influxdb.Config
}

func NewKeycodeCapacityCPU(notifier *Notifier, influxCfg *influxdb.Config) *KeycodeCapacityCPU {
	cpu := KeycodeCapacityCPU{}
	cpu.name = "capacity_cpu"
	cpu.category = "keycode"
	cpu.notifier = notifier
	cpu.eventLevel = make(map[int]events.EventLevel, 0)
	cpu.eventPosted = make(map[int]bool, 0)
	cpu.percentageLevel = make([]int, 0)
	cpu.percentage = 0
	cpu.coreOccupied = 0
	cpu.coreCapacity = 0
	cpu.influxCfg = influxCfg
	cpu.GenerateCriteria()
	return &cpu
}

func (c *KeycodeCapacityCPU) Validate() {
	if c.MeetCriteria() == true {
		if c.eventPosted[c.percentageMeet] == false {
			c.eventPosted[c.percentageMeet] = true
			c.PostEvent()
		}
	}
}

func (c *KeycodeCapacityCPU) GenerateCriteria() {
	eventInterval := strings.Split(c.notifier.EventInterval, ",")

	for index, level := range strings.Split(c.notifier.EventLevel, ",") {
		interval, _ := strconv.Atoi(eventInterval[index])
		c.percentageLevel = append(c.percentageLevel, interval)
		switch level {
		case "Info":
			c.eventLevel[interval] = events.EventLevel_EVENT_LEVEL_INFO
		case "Warn":
			c.eventLevel[interval] = events.EventLevel_EVENT_LEVEL_WARNING
		case "Error":
			c.eventLevel[interval] = events.EventLevel_EVENT_LEVEL_ERROR
		}
		c.eventPosted[interval] = false
	}
}

func (c *KeycodeCapacityCPU) MeetCriteria() bool {
	// Refresh keycode cache before getting keycode information
	KeycodeMgt.Refresh(true)

	// Get cluster CPU capacity
	keycodes, summary, err := KeycodeMgt.GetAllKeycodes()
	if err != nil {
		scope.Error("failed to check criteria when validate license")
		return false
	}

	// If no keycode is applied, we do not need to check the cpu capacity
	if len(keycodes) == 0 {
		c.percentage = 0
		c.percentageMeet = 0
		c.coreOccupied = 0
		c.coreCapacity = 0
		c.clearEventPosted()
		return false
	}

	// Get cluster CPU occupied from influxdb
	cores, err := KeycodeMgt.GetCPUCoresOccupied()
	if err != nil {
		scope.Errorf("Failed to refresh keycode CPU info, unable to get CPU info: %s", err.Error())
		return false
	}

	// If cpu capacity is different from previous which means user may add or delete keycodes
	// So we need to clear event posted list
	if c.coreCapacity != summary.Capacity.CPUs || c.coreOccupied != cores {
		c.clearEventPosted()
	}

	c.coreCapacity = summary.Capacity.CPUs
	c.coreOccupied = cores

	// Check if CPU capacity exceeds keycode limitation
	if c.coreCapacity >= 0 {
		if c.coreOccupied >= c.coreCapacity {
			c.percentage = 100
			return true
		}

		c.percentage = int((float64(c.coreOccupied) / float64(c.coreCapacity)) * 100)
		inRange, percentageMeet := c.percentageInRange(c.percentage)
		if inRange == true {
			c.percentageMeet = percentageMeet
			return true
		}
	}

	return false
}

func (c *KeycodeCapacityCPU) PostEvent() error {
	eventLevel := c.eventLevel[c.percentageMeet]
	message := ""

	if c.percentage >= 100 {
		message = fmt.Sprintf(CPUCoresExceed, c.coreOccupied, c.coreCapacity)
	} else {
		message = fmt.Sprintf(CPUCoresPercentageUsage, c.percentage)
	}

	err := Keycodes.PostEvent(eventLevel, message)
	if err != nil {
		scope.Error(err.Error())
		scope.Error("failed to post event when check keycode cpu capacity")
		return err
	}

	return nil
}

func (c *KeycodeCapacityCPU) clearEventPosted() {
	for k := range c.eventPosted {
		c.eventPosted[k] = false
	}
}

func (c *KeycodeCapacityCPU) percentageInRange(percentage int) (bool, int) {
	for i, p := range c.percentageLevel {
		// Lower bound check
		if i == 0 {
			if percentage <= p {
				return true, p
			}
		}
		// Upper bound check
		if i == len(c.percentageLevel)-1 {
			if percentage >= p {
				return true, p
			}
		}
		// In range check
		if (percentage > c.percentageLevel[i]) && (percentage <= c.percentageLevel[i+1]) {
			fmt.Println(c.percentageLevel[i+1])
			return true, c.percentageLevel[i+1]
		}
	}
	return false, 0
}
