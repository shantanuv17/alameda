package v1alpha1

import (
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	KeycodeMgt "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiLicenses "prophetstor.com/api/datahub/licenses"
)

func (s *ServiceV1alpha1) GetLicense(ctx context.Context, in *empty.Empty) (*ApiLicenses.GetLicenseResponse, error) {
	scope.Debug("Request received from GetLicense grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(s.Config.InfluxDB)
	license := &ApiLicenses.License{Valid: keycodeMgt.IsValid()}

	response := &ApiLicenses.GetLicenseResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		License: license,
	}

	scope.Debug("Response sent from GetLicense grpc function: " + AlamedaUtils.InterfaceToString(response))
	return response, nil
}
