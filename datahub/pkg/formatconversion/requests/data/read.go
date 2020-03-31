package data

import (
	"errors"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	EnumConv "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/metrics"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/schemas"
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

type ReadDataRequestRequestExtended struct {
	*data.ReadDataRequest
}

func (p *ReadDataRequestRequestExtended) Validate() error {
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

	for _, r := range p.GetReadData() {
		metricType := enumconv.MetricTypeNameMap[r.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[r.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[r.GetResourceQuota()]
		if err := isMeasurementFound(schema[0], r.Measurement, metricType, boundary, quota); err != nil {
			return err
		}
	}

	return nil
}

func (p *ReadDataRequestRequestExtended) ProduceRequest() *types.ReadDataRequest {
	request := types.ReadDataRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, r := range p.GetReadData() {
		request.ReadData = append(request.ReadData, NewReadData(r))
	}
	p.postProcessRequest(&request)
	return &request
}

func (p *ReadDataRequestRequestExtended) postProcessRequest(request *types.ReadDataRequest) {
	if p.SchemaMeta.Category == "cluster_status" {
		for i, w := range p.ReadData {
			p.processColumnMetric(w, request.ReadData[i])
			p.processColumnKind(w, request.ReadData[i])
		}
	}
}

func (p *ReadDataRequestRequestExtended) processColumnMetric(before *data.ReadData, after *types.ReadData) {
	metric := metrics.MetricTypeNameMap[before.MetricType]

	for _, condition := range after.QueryCondition.WhereCondition {
		found := false
		index := 0

		// Check if field "metric" is already filled
		for i, key := range condition.Keys {
			if key == "metric" {
				found = true
				index = i
				break
			}
		}

		// Fill the field "metric" if necessary
		if found {
			condition.Values[index] = metric
		} else {
			condition.Keys = append(condition.Keys, "metric")
			condition.Values = append(condition.Values, metric)
			condition.Operators = append(condition.Operators, "=")
			condition.Types = append(condition.Types, common.String)
		}
	}

	if len(after.QueryCondition.WhereCondition) == 0 {
		condition := common.Condition{}
		condition.Keys = append(condition.Keys, "metric")
		condition.Values = append(condition.Values, metric)
		condition.Operators = append(condition.Operators, "=")
		condition.Types = append(condition.Types, common.String)
		after.QueryCondition.WhereCondition = append(after.QueryCondition.WhereCondition, &condition)
	}
}

func (p *ReadDataRequestRequestExtended) processColumnKind(before *data.ReadData, after *types.ReadData) {
	kind := EnumConv.MetricKindNameMap[before.ResourceBoundary]

	for _, condition := range after.QueryCondition.WhereCondition {
		found := false
		index := 0

		// Check if field "kind" is already filled
		for i, key := range condition.Keys {
			if key == "kind" {
				found = true
				index = i
				break
			}
		}

		// Fill the field "kind" if necessary
		if found {
			condition.Values[index] = kind
		} else {
			condition.Keys = append(condition.Keys, "kind")
			condition.Values = append(condition.Values, kind)
			condition.Operators = append(condition.Operators, "=")
			condition.Types = append(condition.Types, common.String)
		}
	}

	if len(after.QueryCondition.WhereCondition) == 0 {
		condition := common.Condition{}
		condition.Keys = append(condition.Keys, "kind")
		condition.Values = append(condition.Values, kind)
		condition.Operators = append(condition.Operators, "=")
		condition.Types = append(condition.Types, common.String)
		after.QueryCondition.WhereCondition = append(after.QueryCondition.WhereCondition, &condition)
	}
}
