package clusterstatus

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewPodDAO(config config.Config) types.PodDAO {
	return influxdb.NewPodWithConfig(*config.InfluxDB)
}
