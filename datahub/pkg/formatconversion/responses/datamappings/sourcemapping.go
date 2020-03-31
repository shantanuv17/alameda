package datamappings

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/enumconv"
	ApiDataMapping "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings"
)

func NewSourceMapping(sourceMapping *datamapping.SourceMapping) *ApiDataMapping.SourceMapping {
	if sourceMapping != nil {
		s := ApiDataMapping.SourceMapping{}
		s.Source = enumconv.SourceNameMap[sourceMapping.Source]
		s.Mapping = sourceMapping.Mapping
		return &s
	}
	return nil
}
