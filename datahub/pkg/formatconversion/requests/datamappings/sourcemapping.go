package datamappings

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/enumconv"
	ApiDataMapping "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings"
)

func NewSourceMapping(sourceMapping *ApiDataMapping.SourceMapping) *datamapping.SourceMapping {
	if sourceMapping != nil {
		c := datamapping.SourceMapping{}
		c.Source = enumconv.SourceNameMap[sourceMapping.Source]
		c.Mapping = sourceMapping.Mapping
		return &c
	}
	return nil
}
