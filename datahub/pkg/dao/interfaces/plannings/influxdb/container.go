package influxdb

import (
	RepoInfluxPlanning "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	Log "prophetstor.com/alameda/pkg/utils/log"
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

var (
	scope = Log.RegisterScope("planning_dao_implement", "planning dao implement", 0)
)

// Container Implements ContainerOperation interface
type ContainerPlannings struct {
	InfluxDBConfig InfluxDB.Config
}

func NewContainerPlanningsWithConfig(config InfluxDB.Config) *ContainerPlannings {
	return &ContainerPlannings{InfluxDBConfig: config}
}

// AddPodPlannings add pod plannings to database
func (c *ContainerPlannings) AddPodPlannings(in *ApiPlannings.CreatePodPlanningsRequest) error {
	containerRepository := RepoInfluxPlanning.NewContainerRepository(&c.InfluxDBConfig)
	return containerRepository.CreateContainerPlannings(in)
}

// ListPodPlannings list pod plannings
func (c *ContainerPlannings) ListPodPlannings(in *ApiPlannings.ListPodPlanningsRequest) ([]*ApiPlannings.PodPlanning, error) {
	containerRepository := RepoInfluxPlanning.NewContainerRepository(&c.InfluxDBConfig)
	return containerRepository.ListContainerPlannings(in)
}
