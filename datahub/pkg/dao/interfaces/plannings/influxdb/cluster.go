package influxdb

import (
	RepoInfluxPlanning "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	ApiPlannings "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/plannings"
)

type ClusterPlannings struct {
	InfluxDBConfig InfluxDB.Config
}

func NewClusterPlanningsWithConfig(config InfluxDB.Config) *ClusterPlannings {
	return &ClusterPlannings{InfluxDBConfig: config}
}

func (c *ClusterPlannings) CreatePlannings(in *ApiPlannings.CreateClusterPlanningsRequest) error {
	repository := RepoInfluxPlanning.NewClusterRepository(&c.InfluxDBConfig)
	return repository.CreatePlannings(in)
}

func (c *ClusterPlannings) ListPlannings(in *ApiPlannings.ListClusterPlanningsRequest) ([]*ApiPlannings.ClusterPlanning, error) {
	repository := RepoInfluxPlanning.NewClusterRepository(&c.InfluxDBConfig)
	return repository.ListPlannings(in)
}
