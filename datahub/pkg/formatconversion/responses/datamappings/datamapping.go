package datamappings

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/schemas"
	ApiDataMapping "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings"
)

func NewDataMapping(dataMapping *datamapping.DataMapping) *ApiDataMapping.DataMapping {
	if dataMapping != nil {
		d := ApiDataMapping.DataMapping{}
		d.SchemaMeta = schemas.NewSchemaMeta(dataMapping.SchemaMeta)
		d.MetricType = enumconv.MetricTypeNameMap[dataMapping.MetricType]
		for _, column := range dataMapping.Columns {
			d.Columns = append(d.Columns, NewColumn(column))
		}
		return &d
	}
	return nil
}
