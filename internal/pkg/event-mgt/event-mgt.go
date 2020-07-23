package eventmgt

import (
	RabbitMQ "github.com/containers-ai/alameda/internal/pkg/message-queue/rabbitmq"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	ApiEvents "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
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
