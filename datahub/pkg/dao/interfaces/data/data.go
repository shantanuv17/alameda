package data

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/data/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/data/types"
)

func NewDataDAO(config config.Config) types.DataDAO {
	return influxdb.NewDataWithConfig(*config.InfluxDB)
}
