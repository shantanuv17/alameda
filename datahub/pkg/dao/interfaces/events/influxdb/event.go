package influxdb

import (
	"encoding/json"
	DaoEventTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/events/types"
	RepoInfluxEvent "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/events"
	InternalRabbitMQ "prophetstor.com/alameda/internal/pkg/message-queue/rabbitmq"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiEvents "prophetstor.com/api/datahub/events"
)

type Event struct {
	InfluxDBConfig *InfluxDB.Config
	RabbitMQConfig *InternalRabbitMQ.Config
}

func NewEventWithConfig(influxConfig *InfluxDB.Config, rabbitMQConfig *InternalRabbitMQ.Config) DaoEventTypes.EventDAO {
	return &Event{InfluxDBConfig: influxConfig, RabbitMQConfig: rabbitMQConfig}
}

func (e *Event) CreateEvents(in *ApiEvents.CreateEventsRequest) error {
	eventRepo := RepoInfluxEvent.NewEventRepository(e.InfluxDBConfig)
	return eventRepo.CreateEvents(in)
}

func (e *Event) ListEvents(in *ApiEvents.ListEventsRequest) ([]*ApiEvents.Event, error) {
	eventRepo := RepoInfluxEvent.NewEventRepository(e.InfluxDBConfig)
	return eventRepo.ListEvents(in)
}

func (e *Event) SendEvents(in *ApiEvents.CreateEventsRequest) error {
	messageQueue, err := InternalRabbitMQ.NewRabbitMQSender(e.RabbitMQConfig)
	if err != nil {
		return err
	}
	defer messageQueue.Close()

	events, err := json.Marshal(in.GetEvents())
	if err != nil {
		return err
	}

	err = messageQueue.SendJsonString("event", string(events))
	return err
}
