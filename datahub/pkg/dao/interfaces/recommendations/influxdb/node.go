package influxdb

import (
	RepoInfluxRecommendation "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiRecommendations "prophetstor.com/api/datahub/recommendations"
)

type NodeRecommendations struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNodeRecommendationsWithConfig(config InfluxDB.Config) *NodeRecommendations {
	return &NodeRecommendations{InfluxDBConfig: config}
}

func (c *NodeRecommendations) CreateRecommendations(recommendations []*ApiRecommendations.NodeRecommendation) error {
	repository := RepoInfluxRecommendation.NewNodeRepository(&c.InfluxDBConfig)
	return repository.CreateRecommendations(recommendations)
}

func (c *NodeRecommendations) ListRecommendations(in *ApiRecommendations.ListNodeRecommendationsRequest) ([]*ApiRecommendations.NodeRecommendation, error) {
	repository := RepoInfluxRecommendation.NewNodeRepository(&c.InfluxDBConfig)
	return repository.ListRecommendations(in)
}
