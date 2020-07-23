package influxdb

import (
	RepoInfluxRecommendation "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	ApiRecommendations "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
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
