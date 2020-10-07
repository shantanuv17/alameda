package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiConfigs "prophetstor.com/api/datahub/configs"
)

func (s *ServiceV1alpha1) CreateConfigs(ctx context.Context, in *ApiConfigs.CreateConfigsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateConfigs grpc function: " + AlamedaUtils.InterfaceToString(in))

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListConfigs(ctx context.Context, in *ApiConfigs.ListConfigsRequest) (*ApiConfigs.ListConfigsResponse, error) {
	scope.Debug("Request received from ListConfigs grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := ApiConfigs.ListConfigsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Configs: []*ApiConfigs.Config{},
	}

	return &response, nil
}

func (s *ServiceV1alpha1) DeleteConfigs(ctx context.Context, in *ApiConfigs.DeleteConfigsRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteConfigs grpc function: " + AlamedaUtils.InterfaceToString(in))
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
