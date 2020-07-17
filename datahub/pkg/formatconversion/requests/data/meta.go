package data

import (
	"errors"
	"fmt"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/schemas"
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

type WriteMetaRequestRequestExtended struct {
	*data.WriteMetaRequest
}

func (p *WriteMetaRequestRequestExtended) Validate() error {
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

	for _, w := range p.GetWriteMeta() {
		metricType := enumconv.MetricTypeNameMap[w.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[w.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[w.GetResourceQuota()]
		m, err := isMeasurementFound(schema[0], w.GetMeasurement(), metricType, boundary, quota)
		if err != nil {
			return err
		}
		if m.IsTS {
			return errors.New(fmt.Sprintf("measurement(%s) is not for metadata usage", m.Name))
		}
		err = m.ColumnTag(w.GetCondition().GetKeys())
		if err != nil {
			return err
		}
		if w.GetColumns() != nil {
			if err := m.ColumnRequired(w.GetColumns()); err != nil {
				return err
			}
			if err := m.ColumnSupported(w.GetColumns()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *WriteMetaRequestRequestExtended) ProduceRequest() *types.WriteMetaRequest {
	request := types.WriteMetaRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, m := range p.GetWriteMeta() {
		request.WriteMeta = append(request.WriteMeta, NewWriteMeta(m))
	}
	return &request
}
