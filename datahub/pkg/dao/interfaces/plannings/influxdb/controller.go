package influxdb

import (
	RepoInfluxPlanning "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

type ControllerPlannings struct {
	InfluxDBConfig InfluxDB.Config
}

func NewControllerPlanningsWithConfig(config InfluxDB.Config) *ControllerPlannings {
	return &ControllerPlannings{InfluxDBConfig: config}
}

func (c *ControllerPlannings) AddControllerPlannings(in *ApiPlannings.CreateControllerPlanningsRequest) error {
	controllerRepository := RepoInfluxPlanning.NewControllerRepository(&c.InfluxDBConfig)
	return controllerRepository.CreateControllerPlannings(in)
}

func (c *ControllerPlannings) ListControllerPlannings(in *ApiPlannings.ListControllerPlanningsRequest) ([]*ApiPlannings.ControllerPlanning, error) {
	controllerRepository := RepoInfluxPlanning.NewControllerRepository(&c.InfluxDBConfig)
	return controllerRepository.ListControllerPlannings(in)
}
