package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoPlannings "prophetstor.com/alameda/datahub/pkg/dao/interfaces/plannings"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

func (s *ServiceV1alpha1) CreateNodePlannings(ctx context.Context, in *ApiPlannings.CreateNodePlanningsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateNodePlannings grpc function: " + AlamedaUtils.InterfaceToString(in))

	nodeDAO := DaoPlannings.NewNodePlanningsDAO(*s.Config)
	err := nodeDAO.CreatePlannings(in)

	if err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, err
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListNodePlannings(ctx context.Context, in *ApiPlannings.ListNodePlanningsRequest) (*ApiPlannings.ListNodePlanningsResponse, error) {
	scope.Debug("Request received from ListNodePlannings grpc function: " + AlamedaUtils.InterfaceToString(in))

	nodeDAO := DaoPlannings.NewNodePlanningsDAO(*s.Config)
	nodePlannings, err := nodeDAO.ListPlannings(in)
	if err != nil {
		scope.Errorf("api ListNodePlannings failed: %v", err)
		response := &ApiPlannings.ListNodePlanningsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
			NodePlannings: nodePlannings,
		}
		return response, nil
	}

	response := &ApiPlannings.ListNodePlanningsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		NodePlannings: nodePlannings,
	}

	return response, nil
}
