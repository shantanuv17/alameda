package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	"prophetstor.com/alameda/datahub/pkg/schemamgt"
	"prophetstor.com/alameda/pkg/database/influxdb"
	"prophetstor.com/alameda/pkg/utils/log"
	"prophetstor.com/api/datahub/events"
	"strconv"
	"strings"
)

var (
	scope      = log.RegisterScope("notifier", "notifier-alerts", 0)
	KeycodeMgt *keycodes.KeycodeMgt
	SchemaMgt  *schemamgt.SchemaManagement
)

type Notifier struct {
	Enabled       bool   `mapstructure:"enabled"`
	Specs         string `mapstructure:"specs"`
	EventInterval string `mapstructure:"eventInterval"`
	EventLevel    string `mapstructure:"eventLevel"`
}

type AlertInterface interface {
	GetName() string
	GetType() string
	GetCategory() string
	GetSpecs() string
	GetEnabled() bool
	MeetCriteria() bool
	Validate()
	GenerateCriteria()
	ClearEventPosted()
}
type AlertMetrics struct {
	notifier     *Notifier
	name         string
	alertType    string
	category     string
	criteriaType CriteriaType
	eventList    []int
	eventLevel   map[int]events.EventLevel
	eventPosted  map[int]bool
}

func Init(influxCfg *influxdb.Config) {
	KeycodeMgt = keycodes.NewKeycodeMgt(influxCfg)
	SchemaMgt = schemamgt.NewSchemaManagement()
}

func (c *AlertMetrics) GetName() string {
	return c.name
}

func (c *AlertMetrics) GetType() string {
	return c.alertType
}

func (c *AlertMetrics) GetCategory() string {
	return c.category
}

func (c *AlertMetrics) GetSpecs() string {
	return c.notifier.Specs
}

func (c *AlertMetrics) GetEnabled() bool {
	return c.notifier.Enabled
}

func (c *AlertMetrics) Validate() {
}

func (c *AlertMetrics) GenerateCriteria() {
	c.eventLevel = make(map[int]events.EventLevel, 0)
	c.eventPosted = make(map[int]bool, 0)
	c.eventList = make([]int, 0)

	switch c.criteriaType {
	case CriteriaTypeGauge:
		c.parseCriteriaGauge()
	case CriteriaTypeContinuous:
		c.parseCriteriaContinuous()
	}
}

func (c *AlertMetrics) ClearEventPosted() {
	for k := range c.eventPosted {
		c.eventPosted[k] = false
	}
}

func (c *AlertMetrics) MeetCriteria() bool {
	return false
}

func (c *AlertMetrics) parseCriteriaGauge() {
	eventInterval := strings.Split(c.notifier.EventInterval, ",")

	for index, level := range strings.Split(c.notifier.EventLevel, ",") {
		interval, _ := strconv.Atoi(eventInterval[index])
		c.eventList = append(c.eventList, interval)
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

func (c *AlertMetrics) parseCriteriaContinuous() {
	eventMap := map[int]events.EventLevel{}
	for _, level := range strings.Split(c.notifier.EventLevel, ",") {
		day, _ := strconv.Atoi(strings.Split(level, ":")[0])
		value := strings.Split(level, ":")[1]

		switch value {
		case "Info":
			eventMap[day] = events.EventLevel_EVENT_LEVEL_INFO
		case "Warn":
			eventMap[day] = events.EventLevel_EVENT_LEVEL_WARNING
		case "Error":
			eventMap[day] = events.EventLevel_EVENT_LEVEL_ERROR
		}
	}

	nowDay := 0
	for _, dayStr := range strings.Split(c.notifier.EventInterval, ",") {
		day, _ := strconv.Atoi(dayStr)
		if _, ok := eventMap[day]; ok {
			nowDay = day
		}
		c.eventLevel[day] = eventMap[nowDay]
		c.eventPosted[day] = false
	}
}
