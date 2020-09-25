package influxdb

import (
	"github.com/pkg/errors"
	EntityInfluxScore "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/scores"
	DaoScore "prophetstor.com/alameda/datahub/pkg/dao/interfaces/scores/types"
	RepoInfluxScore "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/scores"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
)

type Score struct {
	InfluxDBConfig InfluxDB.Config
}

// NewWithConfig New influxdb score dao implement
func NewScoreWithConfig(config InfluxDB.Config) DaoScore.ScoreDAO {
	return &Score{InfluxDBConfig: config}
}

// ListSimulatedSchedulingScores Function implementation of score dao
func (s *Score) ListSimulatedSchedulingScores(request DaoScore.ListRequest) ([]*DaoScore.SimulatedSchedulingScore, error) {

	var (
		err error

		scoreRepository       RepoInfluxScore.SimulatedSchedulingScoreRepository
		influxdbScoreEntities []*EntityInfluxScore.SimulatedSchedulingScoreEntity
		scores                = make([]*DaoScore.SimulatedSchedulingScore, 0)
	)

	scoreRepository = RepoInfluxScore.NewRepositoryWithConfig(s.InfluxDBConfig)
	influxdbScoreEntities, err = scoreRepository.ListScoresByRequest(request)
	if err != nil {
		return scores, errors.Wrap(err, "list simulated scheduing scores failed")
	}

	for _, influxdbScoreEntity := range influxdbScoreEntities {

		score := DaoScore.SimulatedSchedulingScore{
			Timestamp: influxdbScoreEntity.Time,
		}

		if scoreBefore := influxdbScoreEntity.ScoreBefore; scoreBefore != nil {
			score.ScoreBefore = *scoreBefore
		}

		if scoreAfter := influxdbScoreEntity.ScoreAfter; scoreAfter != nil {
			score.ScoreAfter = *scoreAfter
		}

		scores = append(scores, &score)
	}

	return scores, nil
}

// CreateSimulatedSchedulingScores Function implementation of score dao
func (s *Score) CreateSimulatedSchedulingScores(scores []*DaoScore.SimulatedSchedulingScore) error {

	var (
		err error

		scoreRepository RepoInfluxScore.SimulatedSchedulingScoreRepository
	)

	scoreRepository = RepoInfluxScore.NewRepositoryWithConfig(s.InfluxDBConfig)
	err = scoreRepository.CreateScores(scores)
	if err != nil {
		return errors.Wrap(err, "create simulated scheduing scores failed")
	}

	return nil
}
