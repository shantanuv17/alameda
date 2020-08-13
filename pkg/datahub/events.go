package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) CreateEvents(request *events.CreateEventsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateEvents(context.Background(), request)
}

func (p *Client) ListEvents(request *events.ListEventsRequest) (*events.ListEventsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListEvents(context.Background(), request)
}
