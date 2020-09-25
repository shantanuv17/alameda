package types

import (
	ApiEvents "prophetstor.com/api/datahub/events"
)

type EventDAO interface {
	CreateEvents(in *ApiEvents.CreateEventsRequest) error
	ListEvents(in *ApiEvents.ListEventsRequest) ([]*ApiEvents.Event, error)
	SendEvents(in *ApiEvents.CreateEventsRequest) error
}
