package influxdb

import (
	RepoInfluxRecommendation "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

type ClusterRecommendations struct {
	InfluxDBConfig InfluxDB.Config
}

func NewClusterRecommendationsWithConfig(config InfluxDB.Config) *ClusterRecommendations {
	return &ClusterRecommendations{InfluxDBConfig: config}
}

func (c *ClusterRecommendations) CreateRecommendations(recommendations []*ApiRecommendations.ClusterRecommendation) error {
	repository := RepoInfluxRecommendation.NewClusterRepository(&c.InfluxDBConfig)
	return repository.CreateRecommendations(recommendations)
}

func (c *ClusterRecommendations) ListRecommendations(in *ApiRecommendations.ListClusterRecommendationsRequest) ([]*ApiRecommendations.ClusterRecommendation, error) {
	repository := RepoInfluxRecommendation.NewClusterRepository(&c.InfluxDBConfig)
	return repository.ListRecommendations(in)
}
