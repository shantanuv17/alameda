package influxdb

import (
	RepoInfluxPlanning "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

type NodePlannings struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNodePlanningsWithConfig(config InfluxDB.Config) *NodePlannings {
	return &NodePlannings{InfluxDBConfig: config}
}

func (c *NodePlannings) CreatePlannings(in *ApiPlannings.CreateNodePlanningsRequest) error {
	repository := RepoInfluxPlanning.NewNodeRepository(&c.InfluxDBConfig)
	return repository.CreatePlannings(in)
}

func (c *NodePlannings) ListPlannings(in *ApiPlannings.ListNodePlanningsRequest) ([]*ApiPlannings.NodePlanning, error) {
	repository := RepoInfluxPlanning.NewNodeRepository(&c.InfluxDBConfig)
	return repository.ListPlannings(in)
}
