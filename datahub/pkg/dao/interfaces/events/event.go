package events

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/events/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/events/types"
)

func NewEventDAO(config config.Config) types.EventDAO {
	return influxdb.NewEventWithConfig(config.InfluxDB, config.RabbitMQ)
}
