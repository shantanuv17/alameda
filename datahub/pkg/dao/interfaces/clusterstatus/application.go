package clusterstatus

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewApplicationDAO(config config.Config) types.ApplicationDAO {
	return influxdb.NewApplicationWithConfig(*config.InfluxDB)
}
