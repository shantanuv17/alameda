package datahub

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"prophetstor.com/api/datahub/licenses"
)

func (p *Client) GetLicense(request *empty.Empty) (*licenses.GetLicenseResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.GetLicense(context.Background(), request)
}
