package predictions

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewPodPredictionsDAO(config config.Config) types.PodPredictionsDAO {
	return influxdb.NewPodPredictionsWithConfig(*config.InfluxDB)
}
