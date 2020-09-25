package influxdb

import (
	RepoInfluxPlanning "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

type AppPlannings struct {
	InfluxDBConfig InfluxDB.Config
}

func NewAppPlanningsWithConfig(config InfluxDB.Config) *AppPlannings {
	return &AppPlannings{InfluxDBConfig: config}
}

func (c *AppPlannings) CreatePlannings(in *ApiPlannings.CreateApplicationPlanningsRequest) error {
	repository := RepoInfluxPlanning.NewAppRepository(&c.InfluxDBConfig)
	return repository.CreatePlannings(in)
}

func (c *AppPlannings) ListPlannings(in *ApiPlannings.ListApplicationPlanningsRequest) ([]*ApiPlannings.ApplicationPlanning, error) {
	repository := RepoInfluxPlanning.NewAppRepository(&c.InfluxDBConfig)
	return repository.ListPlannings(in)
}
