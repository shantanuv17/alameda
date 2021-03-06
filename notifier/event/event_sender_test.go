package event

import (
	"testing"
	"time"

	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_events "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func Test_eventSender_SendEvents(t *testing.T) {
	datahubAddr := "localhost:50050"
	datahubServiceClient := datahubpkg.NewClient(datahubAddr)
	sender := NewEventSender(datahubServiceClient)
	type args struct {
		events []*datahub_events.Event
	}
	tests := []struct {
		name      string
		evtSender *eventSender
		args      args
		wantErr   bool
	}{
		{
			name:      "send email warning",
			evtSender: sender,
			args: args{
				events: []*datahub_events.Event{
					{
						Time: &timestamp.Timestamp{
							Seconds: time.Now().Unix(),
						},
						ClusterId: "cluster id",
						Source: &datahub_events.EventSource{
							Host:      "email warning host",
							Component: "email warning component",
						},
						Type:    datahub_events.EventType_EVENT_TYPE_EMAIL_NOTIFICATION,
						Version: datahub_events.EventVersion_EVENT_VERSION_V1,
						Level:   datahub_events.EventLevel_EVENT_LEVEL_WARNING,
						Subject: &datahub_events.K8SObjectReference{
							Kind:      "Pod",
							Namespace: "federatorai",
							Name:      "alameda-notifier-7d6b779c47-f6t7q",
						},
						Message: "send email warning message",
						Data:    "send email warning data",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, ttt := range tests {
		tt := ttt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.evtSender.SendEvents(tt.args.events); (err != nil) != tt.wantErr {
				t.Errorf("eventSender.sendEvents() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
