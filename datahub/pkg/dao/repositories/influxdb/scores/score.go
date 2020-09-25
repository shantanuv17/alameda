package scores

import (
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	EntityInfluxScore "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/scores"
	DaoScoreTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/scores/types"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
)

// SimulatedSchedulingScoreRepository Repository of simulated_scheduling_score data
type SimulatedSchedulingScoreRepository struct {
	influxDB *InfluxDB.InfluxClient
}

// NewRepositoryWithConfig New SimulatedSchedulingScoreRepository with influxdb configuration
func NewRepositoryWithConfig(cfg InfluxDB.Config) SimulatedSchedulingScoreRepository {
	return SimulatedSchedulingScoreRepository{
		influxDB: InfluxDB.NewClient(&cfg),
	}
}

// ListScoresByRequest List scores from influxDB
func (r SimulatedSchedulingScoreRepository) ListScoresByRequest(request DaoScoreTypes.ListRequest) ([]*EntityInfluxScore.SimulatedSchedulingScoreEntity, error) {

	var (
		err error

		results      []InfluxClient.Result
		influxdbRows []*InfluxDB.InfluxRow
		scores       = make([]*EntityInfluxScore.SimulatedSchedulingScoreEntity, 0)
	)

	queryCondition := DBCommon.QueryCondition{
		StartTime:      request.QueryCondition.StartTime,
		EndTime:        request.QueryCondition.EndTime,
		StepTime:       request.QueryCondition.StepTime,
		TimestampOrder: request.QueryCondition.TimestampOrder,
		Limit:          request.QueryCondition.Limit,
	}

	influxdbStatement := InfluxDB.Statement{
		QueryCondition: &queryCondition,
		Measurement:    SimulatedSchedulingScore,
	}

	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	results, err = r.influxDB.QueryDB(cmd, string(RepoInflux.Score))
	if err != nil {
		return scores, errors.Wrap(err, "list scores failed")
	}

	influxdbRows = InfluxDB.PackMap(results)
	for _, influxdbRow := range influxdbRows {
		for _, data := range influxdbRow.Data {
			scoreEntity := EntityInfluxScore.NewSimulatedSchedulingScoreEntityFromMap(data)
			scores = append(scores, &scoreEntity)
		}
	}

	return scores, nil

}

// CreateScores Create simulated_scheduling_score data points into influxdb
func (r SimulatedSchedulingScoreRepository) CreateScores(scores []*DaoScoreTypes.SimulatedSchedulingScore) error {

	var (
		err error

		points = make([]*InfluxClient.Point, 0)
	)

	for _, score := range scores {

		time := score.Timestamp
		scoreBefore := score.ScoreBefore
		scoreAfter := score.ScoreAfter
		entity := EntityInfluxScore.SimulatedSchedulingScoreEntity{
			Time:        time,
			ScoreBefore: &scoreBefore,
			ScoreAfter:  &scoreAfter,
		}

		point, err := entity.InfluxDBPoint(string(SimulatedSchedulingScore))
		if err != nil {
			return errors.Wrap(err, "create scores failed")
		}
		points = append(points, point)
	}

	err = r.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.Score),
	})
	if err != nil {
		return errors.Wrap(err, "create scores failed")
	}

	return nil
}
