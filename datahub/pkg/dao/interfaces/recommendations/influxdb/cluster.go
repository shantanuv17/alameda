package influxdb

import (
	RepoInfluxRecommendation "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	ApiRecommendations "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
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
