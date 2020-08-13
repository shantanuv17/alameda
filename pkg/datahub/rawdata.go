package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/rawdata"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) ReadRawdata(request *rawdata.ReadRawdataRequest) (*rawdata.ReadRawdataResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ReadRawdata(context.Background(), request)
}

func (p *Client) WriteRawdata(request *rawdata.WriteRawdataRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.WriteRawdata(context.Background(), request)
}
