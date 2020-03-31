package enumconv

import (
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	ApiDataMapping "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/datamappings"
)

var SourceNameMap = map[ApiDataMapping.Source]datamapping.Source{
	ApiDataMapping.Source_SOURCE_UNDEFINED:  datamapping.SourceUndefined,
	ApiDataMapping.Source_SOURCE_DATADOG:    datamapping.Datadog,
	ApiDataMapping.Source_SOURCE_DYNATRACE:  datamapping.Dynatrace,
	ApiDataMapping.Source_SOURCE_K8S:        datamapping.K8s,
	ApiDataMapping.Source_SOURCE_PROMETHEUS: datamapping.Prometheus,
	ApiDataMapping.Source_SOURCE_SYSDIG:     datamapping.Sysdig,
}
