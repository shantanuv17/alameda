package predictions

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewApplicationPredictionsDAO(config config.Config) types.ApplicationPredictionsDAO {
	return influxdb.NewApplicationPredictionsWithConfig(*config.InfluxDB)
}
