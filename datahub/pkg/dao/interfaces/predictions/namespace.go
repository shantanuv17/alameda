package predictions

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewNamespacePredictionsDAO(config config.Config) types.NamespacePredictionsDAO {
	return influxdb.NewNamespacePredictionsWithConfig(*config.InfluxDB)
}
