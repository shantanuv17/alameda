package predictions

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewControllerPredictionsDAO(config config.Config) types.ControllerPredictionsDAO {
	return influxdb.NewControllerPredictionsWithConfig(*config.InfluxDB)
}
