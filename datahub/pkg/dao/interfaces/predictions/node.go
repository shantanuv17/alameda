package predictions

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewNodePredictionsDAO(config config.Config) types.NodePredictionsDAO {
	return influxdb.NewNodePredictionsWithConfig(*config.InfluxDB)
}
