package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/schemas"
)

func (p *Client) CreateSchemas(request *schemas.CreateSchemasRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateSchemas(context.Background(), request)
}

func (p *Client) ListSchemas(request *schemas.ListSchemasRequest) (*schemas.ListSchemasResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListSchemas(context.Background(), request)
}

func (p *Client) DeleteSchemas(request *schemas.DeleteSchemasRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteSchemas(context.Background(), request)
}
