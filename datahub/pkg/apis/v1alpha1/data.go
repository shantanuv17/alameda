package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoData "prophetstor.com/alameda/datahub/pkg/dao/interfaces/data"
	FormatRequest "prophetstor.com/alameda/datahub/pkg/formatconversion/requests/data"
	FormatResponse "prophetstor.com/alameda/datahub/pkg/formatconversion/responses/data"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiData "prophetstor.com/api/datahub/data"
)

func (s *ServiceV1alpha1) ReadData(ctx context.Context, in *ApiData.ReadDataRequest) (*ApiData.ReadDataResponse, error) {
	scope.Debug("Request received from ReadData grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.ReadDataRequestRequestExtended{ReadDataRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiData.ReadDataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	if in.GetReadData() == nil {
		return &ApiData.ReadDataResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Data: nil,
		}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	readData, err := dataDAO.ReadData(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ReadData failed: %+v", err)
		return &ApiData.ReadDataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	response := ApiData.ReadDataResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Data: FormatResponse.NewData(readData),
	}

	return &response, nil
}

func (s *ServiceV1alpha1) WriteData(ctx context.Context, in *ApiData.WriteDataRequest) (*status.Status, error) {
	scope.Debug("Request received from WriteData grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.WriteDataRequestRequestExtended{WriteDataRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	if in.GetWriteData() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	if err := dataDAO.WriteData(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to write data: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}

func (s *ServiceV1alpha1) DeleteData(ctx context.Context, in *ApiData.DeleteDataRequest) (*status.Status, error) {
	scope.Debug("Request received from WriteData grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.DeleteDataRequestRequestExtended{DeleteDataRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	if in.GetDeleteData() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	if err := dataDAO.DeleteData(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete data: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}

func (s *ServiceV1alpha1) WriteMeta(ctx context.Context, in *ApiData.WriteMetaRequest) (*status.Status, error) {
	scope.Debug("Request received from WriteMeta grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.WriteMetaRequestRequestExtended{WriteMetaRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	if in.GetWriteMeta() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	if err := dataDAO.WriteMeta(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to write meta: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}
