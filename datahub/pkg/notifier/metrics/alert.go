package metrics

import (
	"github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/utils/log"
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
	GetCategory() string
	GetSpecs() string
	GetEnabled() bool
	Validate()
	GenerateCriteria()
	MeetCriteria() bool
}
type AlertMetrics struct {
	name     string
	category string
	notifier *Notifier
}

func Init(influxCfg *influxdb.Config) {
	KeycodeMgt = keycodes.NewKeycodeMgt(influxCfg)
	SchemaMgt = schemamgt.NewSchemaManagement()
}

func (c *AlertMetrics) GetName() string {
	return c.name
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

}

func (c *AlertMetrics) MeetCriteria() bool {
	return false
}
