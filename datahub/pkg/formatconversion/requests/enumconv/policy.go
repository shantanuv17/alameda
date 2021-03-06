package enumconv

import (
	ApiResources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
)

const (
	Stable  string = "Stable"
	Compact string = "Compact"
)

var RecommendationPolicyName = map[ApiResources.RecommendationPolicy]string{
	ApiResources.RecommendationPolicy_STABLE:  Stable,
	ApiResources.RecommendationPolicy_COMPACT: Compact,
}

var RecommendationPolicyValue = map[string]ApiResources.RecommendationPolicy{
	Stable:  ApiResources.RecommendationPolicy_STABLE,
	Compact: ApiResources.RecommendationPolicy_COMPACT,
}
