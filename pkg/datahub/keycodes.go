package datahub

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/keycodes"
)

func (p *Client) ActivateRegistrationData(request *keycodes.ActivateRegistrationDataRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ActivateRegistrationData(context.Background(), request)
}

func (p *Client) AddKeycode(request *keycodes.AddKeycodeRequest) (*keycodes.AddKeycodeResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.AddKeycode(context.Background(), request)
}

func (p *Client) DeleteKeycode(request *keycodes.DeleteKeycodeRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteKeycode(context.Background(), request)
}

func (p *Client) ListKeycodes(request *keycodes.ListKeycodesRequest) (*keycodes.ListKeycodesResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListKeycodes(context.Background(), request)
}

func (p *Client) GenerateRegistrationData(request *empty.Empty) (*keycodes.GenerateRegistrationDataResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.GenerateRegistrationData(context.Background(), request)
}
