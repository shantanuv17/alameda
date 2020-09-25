package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/data"
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

func (p *Client) WriteMeta(request *data.WriteMetaRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.WriteMeta(context.Background(), request)
}
