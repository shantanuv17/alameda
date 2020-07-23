package influxdb

import (
	RepoInfluxPlanning "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	ApiPlannings "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/plannings"
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
