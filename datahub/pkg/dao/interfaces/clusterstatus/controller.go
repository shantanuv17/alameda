package clusterstatus

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewControllerDAO(config config.Config) types.ControllerDAO {
	return influxdb.NewControllerWithConfig(*config.InfluxDB)
}
