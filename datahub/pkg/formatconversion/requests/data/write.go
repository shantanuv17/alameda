package data

import (
	"errors"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/data/types"
	EnumConv "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/metrics"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/schemas"
	"prophetstor.com/alameda/datahub/pkg/schemamgt"
	"prophetstor.com/api/datahub/data"
	ApiSchema "prophetstor.com/api/datahub/schemas"
)

type WriteDataRequestRequestExtended struct {
	*data.WriteDataRequest
}

func (p *WriteDataRequestRequestExtended) Validate() error {
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

	for _, w := range p.GetWriteData() {
		metricType := enumconv.MetricTypeNameMap[w.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[w.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[w.GetResourceQuota()]
		m, err := isMeasurementFound(schema[0], w.GetMeasurement(), metricType, boundary, quota)
		if err != nil {
			return err
		}
		if err := m.ColumnRequired(w.GetColumns()); err != nil {
			return err
		}
		if err := m.ColumnSupported(w.GetColumns()); err != nil {
			return err
		}
	}

	return nil
}

func (p *WriteDataRequestRequestExtended) ProduceRequest() *types.WriteDataRequest {
	request := types.WriteDataRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, w := range p.GetWriteData() {
		request.WriteData = append(request.WriteData, NewWriteData(w))
	}
	p.postProcessRequest(&request)
	return &request
}

func (p *WriteDataRequestRequestExtended) postProcessRequest(request *types.WriteDataRequest) {
	switch p.SchemaMeta.Scope {
	case ApiSchema.Scope_SCOPE_PREDICTION:
		if p.SchemaMeta.Category == "cluster_status" {
			for i, w := range p.WriteData {
				p.processColumnMetric(w, request.WriteData[i])
				p.processColumnKind(w, request.WriteData[i])
			}
		}
	}
}

func (p *WriteDataRequestRequestExtended) processColumnMetric(before *data.WriteData, after *types.WriteData) {
	found := false
	index := 0
	metric := metrics.MetricTypeNameMap[before.MetricType]

	// Check if field "metric" is already filled
	for i, column := range after.Columns {
		if column == "metric" {
			found = true
			index = i
			break
		}
	}

	// Fill the field "metric" if necessary
	if found {
		for _, row := range after.Rows {
			row.Values[index] = metric
		}
	} else {
		after.Columns = append(after.Columns, "metric")
		for _, row := range after.Rows {
			row.Values = append(row.Values, metric)
		}
	}
}

func (p *WriteDataRequestRequestExtended) processColumnKind(before *data.WriteData, after *types.WriteData) {
	found := false
	index := 0
	kind := EnumConv.MetricKindNameMap[before.ResourceBoundary]

	// Check if field "kind" is already filled
	for i, column := range after.Columns {
		if column == "kind" {
			found = true
			index = i
			break
		}
	}

	// Fill the field "kind" if necessary
	if found {
		for _, row := range after.Rows {
			row.Values[index] = kind
		}
	} else {
		after.Columns = append(after.Columns, "kind")
		for _, row := range after.Rows {
			row.Values = append(row.Values, kind)
		}
	}
}
