package types

import (
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

// ContainerOperation defines container measurement operation of recommendation database
type ContainerRecommendationsDAO interface {
	AddPodRecommendations(in *ApiRecommendations.CreatePodRecommendationsRequest) error
	ListPodRecommendations(in *ApiRecommendations.ListPodRecommendationsRequest) ([]*ApiRecommendations.PodRecommendation, error)
	ListAvailablePodRecommendations(*ApiRecommendations.ListPodRecommendationsRequest) ([]*ApiRecommendations.PodRecommendation, error)
}
