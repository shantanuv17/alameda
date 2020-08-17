package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) ReadData(request *data.ReadDataRequest) (*data.ReadDataResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ReadData(context.Background(), request)
}

func (p *Client) WriteData(request *data.WriteDataRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.WriteData(context.Background(), request)
}

func (p *Client) DeleteData(request *data.DeleteDataRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteData(context.Background(), request)
}
