package influxdb

import (
	RepoInfluxRecommendation "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

type NamespaceRecommendations struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNamespaceRecommendationsWithConfig(config InfluxDB.Config) *NamespaceRecommendations {
	return &NamespaceRecommendations{InfluxDBConfig: config}
}

func (c *NamespaceRecommendations) CreateRecommendations(recommendations []*ApiRecommendations.NamespaceRecommendation) error {
	repository := RepoInfluxRecommendation.NewNamespaceRepository(&c.InfluxDBConfig)
	return repository.CreateRecommendations(recommendations)
}

func (c *NamespaceRecommendations) ListRecommendations(in *ApiRecommendations.ListNamespaceRecommendationsRequest) ([]*ApiRecommendations.NamespaceRecommendation, error) {
	repository := RepoInfluxRecommendation.NewNamespaceRepository(&c.InfluxDBConfig)
	return repository.ListRecommendations(in)
}
