package clusterstatus

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewNamespaceDAO(config config.Config) types.NamespaceDAO {
	return influxdb.NewNamespaceWithConfig(*config.InfluxDB)
}
