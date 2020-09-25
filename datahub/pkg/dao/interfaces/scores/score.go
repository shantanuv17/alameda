package scores

import (
	"prophetstor.com/alameda/datahub/pkg/config"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/scores/influxdb"
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/scores/types"
)

func NewScoreDAO(config config.Config) types.ScoreDAO {
	return influxdb.NewScoreWithConfig(*config.InfluxDB)
}
