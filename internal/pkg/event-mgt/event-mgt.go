package eventmgt

import (
	RabbitMQ "prophetstor.com/alameda/internal/pkg/message-queue/rabbitmq"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiEvents "prophetstor.com/api/datahub/events"
)

var (
	gInfluxDBCfg    = InfluxDB.NewDefaultConfig()
	gRabbitMQConfig = RabbitMQ.NewDefaultConfig()
)

type EventMgt struct {
	RabbitMQConfig *RabbitMQ.Config
	influxDB       *InfluxDB.InfluxClient
}

func InitEventMgt(influxDBCfg *InfluxDB.Config, rabbitMQConfig *RabbitMQ.Config) {
	gInfluxDBCfg = influxDBCfg
	gRabbitMQConfig = rabbitMQConfig
}

func NewEventMgt(influxDBCfg *InfluxDB.Config, rabbitMQConfig *RabbitMQ.Config) *EventMgt {
	return &EventMgt{
		influxDB: &InfluxDB.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
		RabbitMQConfig: rabbitMQConfig,
	}
}

func PostEvents(in *ApiEvents.CreateEventsRequest) error {
	eventMgt := NewEventMgt(gInfluxDBCfg, gRabbitMQConfig)
	return eventMgt.PostEvents(in)
}

func ListEvents(in *ApiEvents.ListEventsRequest) ([]*ApiEvents.Event, error) {
	eventMgt := NewEventMgt(gInfluxDBCfg, gRabbitMQConfig)
	return eventMgt.ListEvents(in)
}
