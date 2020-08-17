package datahub

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) Ping(request *empty.Empty) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.Ping(context.Background(), request)
}
