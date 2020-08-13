package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"google.golang.org/genproto/googleapis/rpc/status"
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
