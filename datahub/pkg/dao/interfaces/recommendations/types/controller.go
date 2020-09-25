package types

import (
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

// ContainerOperation defines container measurement operation of recommendation database
type ControllerRecommendationsDAO interface {
	AddControllerRecommendations([]*ApiRecommendations.ControllerRecommendation) error
	ListControllerRecommendations(in *ApiRecommendations.ListControllerRecommendationsRequest) ([]*ApiRecommendations.ControllerRecommendation, error)
}
