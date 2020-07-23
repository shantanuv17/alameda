package influxdb

import (
	RepoInfluxRecommendation "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	ApiRecommendations "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
)

type ControllerRecommendations struct {
	InfluxDBConfig InfluxDB.Config
}

func NewControllerRecommendationsWithConfig(config InfluxDB.Config) *ControllerRecommendations {
	return &ControllerRecommendations{InfluxDBConfig: config}
}

func (c *ControllerRecommendations) CreateControllerRecommendations(controllerRecommendations []*ApiRecommendations.ControllerRecommendation) error {
	controllerRepository := RepoInfluxRecommendation.NewControllerRepository(&c.InfluxDBConfig)
	return controllerRepository.CreateControllerRecommendations(controllerRecommendations)
}

func (c *ControllerRecommendations) ListControllerRecommendations(in *ApiRecommendations.ListControllerRecommendationsRequest) ([]*ApiRecommendations.ControllerRecommendation, error) {
	controllerRepository := RepoInfluxRecommendation.NewControllerRepository(&c.InfluxDBConfig)
	return controllerRepository.ListControllerRecommendations(in)
}
