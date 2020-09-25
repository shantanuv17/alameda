package schemas

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/schemas/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/schemas/types"
)

func NewSchemaDAO(config config.Config) types.SchemaDAO {
	return influxdb.NewSchemaWithConfig(*config.InfluxDB)
}
