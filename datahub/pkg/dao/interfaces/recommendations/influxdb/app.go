package influxdb

import (
	RepoInfluxRecommendation "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

type AppRecommendations struct {
	InfluxDBConfig InfluxDB.Config
}

func NewAppRecommendationsWithConfig(config InfluxDB.Config) *AppRecommendations {
	return &AppRecommendations{InfluxDBConfig: config}
}

func (c *AppRecommendations) CreateRecommendations(recommendations []*ApiRecommendations.ApplicationRecommendation) error {
	repository := RepoInfluxRecommendation.NewAppRepository(&c.InfluxDBConfig)
	return repository.CreateRecommendations(recommendations)
}

func (c *AppRecommendations) ListRecommendations(in *ApiRecommendations.ListApplicationRecommendationsRequest) ([]*ApiRecommendations.ApplicationRecommendation, error) {
	repository := RepoInfluxRecommendation.NewAppRepository(&c.InfluxDBConfig)
	return repository.ListRecommendations(in)
}
