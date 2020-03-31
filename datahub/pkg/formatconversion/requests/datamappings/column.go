package datamappings

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/schemas"
	ApiDataMapping "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings"
)

func NewColumn(column *ApiDataMapping.Column) *datamapping.Column {
	if column != nil {
		c := datamapping.Column{}
		c.ColumnMeta = schemas.NewColumn(column.ColumnMeta)
		c.Name = column.Name
		for _, sourceMapping := range column.SourceMappings {
			c.SourceMappings = append(c.SourceMappings, NewSourceMapping(sourceMapping))
		}
		return &c
	}
	return nil
}
