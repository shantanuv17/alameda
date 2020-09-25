package influxdb

import (
	RepoInfluxRecommendation "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
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
