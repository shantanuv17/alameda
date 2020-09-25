package v1alpha1

import (
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	DaoSchema "prophetstor.com/alameda/datahub/pkg/dao/interfaces/schemas"
	FormatRequest "prophetstor.com/alameda/datahub/pkg/formatconversion/requests/schemas"
	FormatResponse "prophetstor.com/alameda/datahub/pkg/formatconversion/responses/schemas"
	SchemaMgt "prophetstor.com/alameda/datahub/pkg/schemamgt"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	ApiSchemas "prophetstor.com/api/datahub/schemas"
)

func (s *ServiceV1alpha1) CreateSchemas(ctx context.Context, in *ApiSchemas.CreateSchemasRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateSchemas grpc function: " + AlamedaUtils.InterfaceToString(in))

	if in.GetSchemas() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	requestExt := FormatRequest.CreateSchemasRequestExtended{CreateSchemasRequest: in}
	if requestExt.Validate() != nil {
		return &status.Status{
			Code: int32(code.Code_INVALID_ARGUMENT),
		}, nil
	}

	// Get schema write lock
	SchemaMgt.RWLock.Lock()
	defer SchemaMgt.RWLock.Unlock()

	schemaDAO := DaoSchema.NewSchemaDAO(*s.Config)
	if err := schemaDAO.CreateSchemas(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to create schemas: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListSchemas(ctx context.Context, in *ApiSchemas.ListSchemasRequest) (*ApiSchemas.ListSchemasResponse, error) {
	scope.Debug("Request received from ListSchemas grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.ListSchemasRequestExtended{ListSchemasRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiSchemas.ListSchemasResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	// Get schema read lock
	SchemaMgt.RWLock.RLock()
	defer SchemaMgt.RWLock.RUnlock()

	schemaDAO := DaoSchema.NewSchemaDAO(*s.Config)
	schemaList, err := schemaDAO.ListSchemas(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ListApplications failed: %+v", err)
		return &ApiSchemas.ListSchemasResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	schemas := make([]*ApiSchemas.Schema, 0)
	for _, schema := range schemaList {
		schemas = append(schemas, FormatResponse.NewSchema(schema))
	}

	response := ApiSchemas.ListSchemasResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Schemas: schemas,
	}

	return &response, nil
}

func (s *ServiceV1alpha1) DeleteSchemas(ctx context.Context, in *ApiSchemas.DeleteSchemasRequest) (*status.Status, error) {
	scope.Debug("Request received from ListSchemas grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.DeleteSchemasRequestExtended{DeleteSchemasRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	// Get schema write lock
	SchemaMgt.RWLock.Lock()
	defer SchemaMgt.RWLock.Unlock()

	schemaDAO := DaoSchema.NewSchemaDAO(*s.Config)
	if err := schemaDAO.DeleteSchemas(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete schemas: %+v", err)
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
