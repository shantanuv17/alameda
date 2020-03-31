package v1alpha1

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt"
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	FormatRequest "github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/datamappings"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/enumconv"
	FormatResponse "github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/datamappings"
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	ApiMappings "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	"reflect"
)

func (s *ServiceV1alpha1) CreateDataMappings(ctx context.Context, in *ApiMappings.CreateDataMappingsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateDataMappings grpc function: " + AlamedaUtils.InterfaceToString(in))

	if in.GetDataMappings() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	// Get schema write lock
	datamappingmgt.RWLock.Lock()
	defer datamappingmgt.RWLock.Unlock()

	dataMappings := make([]*datamapping.DataMapping, 0)
	for _, dataMapping := range in.GetDataMappings() {
		dataMappings = append(dataMappings, FormatRequest.NewDataMapping(dataMapping))
	}

	dataMappingMgt := datamappingmgt.NewDataMappingManagement()
	dataMappingMgt.AddDataMappings(dataMappings)
	err := dataMappingMgt.Flush()
	if err != nil {
		scope.Errorf("failed to create data mappings: %+v", err.Error())
		return &status.Status{Code: int32(code.Code_INTERNAL), Message: err.Error()}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}

func (s *ServiceV1alpha1) ReadDataMappings(ctx context.Context, in *ApiMappings.ReadDataMappingsRequest) (*ApiMappings.ReadDataMappingsResponse, error) {
	scope.Debug("Request received from ListDataMappings grpc function: " + AlamedaUtils.InterfaceToString(in))

	table := enumconv.ScopeNameMap[in.GetSchemaMeta().GetScope()]
	category := in.GetSchemaMeta().GetCategory()
	schemaType := in.GetSchemaMeta().GetType()
	metricType := enumconv.MetricTypeNameMap[in.GetMetricType()]

	// Get data mapping read lock
	datamappingmgt.RWLock.RLock()
	defer datamappingmgt.RWLock.RUnlock()

	// Get schema read lock
	schemamgt.RWLock.RLock()
	defer schemamgt.RWLock.RUnlock()

	dataMappingMgt := datamappingmgt.NewDataMappingManagement()
	dataMappings := dataMappingMgt.GetDataMappings(table, category, schemaType, metricType)

	schemaMgt := schemamgt.NewSchemaManagement()
	schemas := schemaMgt.GetSchemas(table, category, schemaType)

	for _, dataMapping := range dataMappings {
		for _, schema := range schemas {
			if reflect.DeepEqual(dataMapping.SchemaMeta, schema.SchemaMeta) {
				for _, measurement := range schema.Measurements {
					if dataMapping.MetricType == measurement.MetricType {
						for _, column := range dataMapping.Columns {
							for _, c := range measurement.Columns {
								if column.Name == c.Name {
									column.ColumnMeta = c
									break
								}
							}
						}
					}
				}
			}
		}
	}

	dataMappingList := make([]*ApiMappings.DataMapping, 0)
	for _, dataMapping := range dataMappings {
		dataMappingList = append(dataMappingList, FormatResponse.NewDataMapping(dataMapping))
	}

	// Pack response
	response := ApiMappings.ReadDataMappingsResponse{}
	response.Status = &status.Status{Code: int32(code.Code_OK)}
	response.DataMappings = dataMappingList

	return &response, nil
}

func (s *ServiceV1alpha1) DeleteDataMappings(ctx context.Context, in *ApiMappings.DeleteDataMappingsRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteDataMappings grpc function: " + AlamedaUtils.InterfaceToString(in))

	table := enumconv.ScopeNameMap[in.GetSchemaMeta().GetScope()]
	category := in.GetSchemaMeta().GetCategory()
	schemaType := in.GetSchemaMeta().GetType()
	metricType := enumconv.MetricTypeNameMap[in.GetMetricType()]

	// Get schema write lock
	datamappingmgt.RWLock.Lock()
	defer datamappingmgt.RWLock.Unlock()

	dataMappingMgt := datamappingmgt.NewDataMappingManagement()
	err := dataMappingMgt.DeleteDataMappings(table, category, schemaType, metricType)
	if err != nil {
		scope.Errorf("failed to delete data mappings: %+v", err.Error())
		return &status.Status{Code: int32(code.Code_INTERNAL), Message: err.Error()}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}
