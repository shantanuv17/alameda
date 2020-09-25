package influxdb

import (
	RepoInfluxPlanning "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiPlannings "prophetstor.com/api/datahub/plannings"
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
