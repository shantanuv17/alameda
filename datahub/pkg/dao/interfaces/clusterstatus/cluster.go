package clusterstatus

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewClusterDAO(config config.Config) types.ClusterDAO {
	return influxdb.NewClusterWithConfig(*config.InfluxDB)
}
