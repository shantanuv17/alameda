package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/events"
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
