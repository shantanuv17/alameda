package data

import (
	"errors"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/schemas"
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

type DeleteDataRequestRequestExtended struct {
	*data.DeleteDataRequest
}

func (p *DeleteDataRequestRequestExtended) Validate() error {
	if p.GetSchemaMeta() == nil {
		return errors.New("schema meta is not given")
	}

	schemaMgt := schemamgt.NewSchemaManagement()
	schemaMeta := schemas.NewSchemaMeta(p.GetSchemaMeta())

	if err := isSchemaMetaComplete(schemaMeta); err != nil {
		return err
	}

	schema := schemaMgt.GetSchemas(schemaMeta.Scope, schemaMeta.Category, schemaMeta.Type)
	if len(schema) == 0 {
		return errors.New("schema is not found")
	}

	for _, d := range p.GetDeleteData() {
		metricType := enumconv.MetricTypeNameMap[d.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[d.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[d.GetResourceQuota()]
		m, err := isMeasurementFound(schema[0], d.Measurement, metricType, boundary, quota)
		if err != nil {
			return err
		}
		for _, w := range d.GetQueryCondition().GetWhereCondition() {
			if err := m.ColumnSupported(w.GetKeys()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *DeleteDataRequestRequestExtended) ProduceRequest() *types.DeleteDataRequest {
	request := types.DeleteDataRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, d := range p.GetDeleteData() {
		request.DeleteData = append(request.DeleteData, NewDeleteData(d))
	}
	return &request
}
