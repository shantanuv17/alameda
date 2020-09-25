package event

import (
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/code"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	datahub_events "prophetstor.com/api/datahub/events"
)

type eventSender struct {
	datahubClient *datahubpkg.Client
}

func NewEventSender(datahubClient *datahubpkg.Client) *eventSender {
	return &eventSender{
		datahubClient: datahubClient,
	}
}

func (evtSender *eventSender) SendEvents(events []*datahub_events.Event) error {
	if len(events) == 0 {
		return nil
	}

	request := datahub_events.CreateEventsRequest{
		Events: events,
	}
	status, err := evtSender.datahubClient.CreateEvents(&request)
	if err != nil {
		return errors.Errorf("send events to Datahub failed: %s", err.Error())
	} else if status == nil {
		return errors.Errorf("send events to Datahub failed: receive nil status")
	} else if status.Code != int32(code.Code_OK) {
		return errors.Errorf("send events to Datahub failed: statusCode: %d, message: %s",
			status.Code, status.Message)
	}

	return nil
}
