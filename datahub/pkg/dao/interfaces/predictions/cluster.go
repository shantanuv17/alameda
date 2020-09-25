package predictions

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewClusterPredictionsDAO(config config.Config) types.ClusterPredictionsDAO {
	return influxdb.NewClusterPredictionsWithConfig(*config.InfluxDB)
}
