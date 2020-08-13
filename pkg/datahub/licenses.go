package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/licenses"
	"github.com/golang/protobuf/ptypes/empty"
)

func (p *Client) GetLicense(request *empty.Empty) (*licenses.GetLicenseResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.GetLicense(context.Background(), request)
}
