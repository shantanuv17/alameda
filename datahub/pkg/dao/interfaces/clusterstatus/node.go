package clusterstatus

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewNodeDAO(config config.Config) types.NodeDAO {
	return influxdb.NewNodeWithConfig(*config.InfluxDB)
}
