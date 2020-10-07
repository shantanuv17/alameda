package datahub

import (
	"context"

	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/configs"
)

func (p *Client) CreateConfigs(request *configs.CreateConfigsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateConfigs(context.Background(), request)
}

func (p *Client) ListConfigs(request *configs.ListConfigsRequest) (*configs.ListConfigsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListConfigs(context.Background(), request)
}

func (p *Client) DeleteConfigs(request *configs.DeleteConfigsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteConfigs(context.Background(), request)
}
