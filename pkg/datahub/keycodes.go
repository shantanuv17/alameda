package datahub

import (
	"context"
	"github.com/containers-ai/api/datahub/keycodes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) ActivateRegistrationData(request *keycodes.ActivateRegistrationDataRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.KeycodesServiceClient.ActivateRegistrationData(context.Background(), request)
}

func (p *Client) AddKeycode(request *keycodes.AddKeycodeRequest) (*keycodes.AddKeycodeResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.KeycodesServiceClient.AddKeycode(context.Background(), request)
}

func (p *Client) DeleteKeycode(request *keycodes.DeleteKeycodeRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.KeycodesServiceClient.DeleteKeycode(context.Background(), request)
}

func (p *Client) ListKeycodes(request *keycodes.ListKeycodesRequest) (*keycodes.ListKeycodesResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.KeycodesServiceClient.ListKeycodes(context.Background(), request)
}

func (p *Client) GenerateRegistrationData(request *empty.Empty) (*keycodes.GenerateRegistrationDataResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.KeycodesServiceClient.GenerateRegistrationData(context.Background(), request)
}
