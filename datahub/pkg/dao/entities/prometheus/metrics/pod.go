package metrics

import (
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
)

type PodCPUUsageMillicoresEntity struct {
	NamespaceName string
	PodName       string
	Samples       []FormatTypes.Sample
}

type PodMemoryUsageBytesEntity struct {
	NamespaceName string
	PodName       string
	Samples       []FormatTypes.Sample
}
