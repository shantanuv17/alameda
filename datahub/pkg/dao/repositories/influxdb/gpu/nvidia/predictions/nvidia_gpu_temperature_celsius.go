package predictions

import (
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	EntityInfluxGpuPrediction "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/gpu/nvidia/predictions"
	DaoGpu "prophetstor.com/alameda/datahub/pkg/dao/interfaces/gpu/influxdb"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	DatahubUtils "prophetstor.com/alameda/datahub/pkg/utils"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	InfluxModels "prophetstor.com/alameda/pkg/database/influxdb/models"
	"strconv"
)

type TemperatureCelsiusRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewTemperatureCelsiusRepositoryWithConfig(cfg InfluxDB.Config) *TemperatureCelsiusRepository {
	return &TemperatureCelsiusRepository{
		influxDB: InfluxDB.NewClient(&cfg),
	}
}

func (r *TemperatureCelsiusRepository) CreatePredictions(predictions []*DaoGpu.GpuPrediction) error {
	points := make([]*InfluxClient.Point, 0)

	for _, prediction := range predictions {
		granularity := int64(30)
		if prediction.Granularity != 0 {
			granularity = prediction.Granularity
		}

		for _, metric := range prediction.Metrics {
			// Parse float string to value
			valueInFloat64, err := DatahubUtils.StringToFloat64(metric.Value)
			if err != nil {
				return errors.Wrap(err, "failed to parse string to float64")
			}

			// Pack influx tags
			tags := map[string]string{
				EntityInfluxGpuPrediction.TemperatureCelsiusHost:        prediction.Metadata.Host,
				EntityInfluxGpuPrediction.TemperatureCelsiusName:        prediction.Name,
				EntityInfluxGpuPrediction.TemperatureCelsiusUuid:        prediction.Uuid,
				EntityInfluxGpuPrediction.TemperatureCelsiusGranularity: strconv.FormatInt(granularity, 10),
			}

			// Pack influx fields
			fields := map[string]interface{}{
				EntityInfluxGpuPrediction.TemperatureCelsiusModelId:      metric.ModelId,
				EntityInfluxGpuPrediction.TemperatureCelsiusPredictionId: metric.PredictionId,
				EntityInfluxGpuPrediction.TemperatureCelsiusMinorNumber:  prediction.Metadata.MinorNumber,
				EntityInfluxGpuPrediction.TemperatureCelsiusValue:        valueInFloat64,
			}

			// Add to influx point list
			point, err := InfluxClient.NewPoint(string(TemperatureCelsius), tags, fields, metric.Timestamp)
			if err != nil {
				return errors.Wrap(err, "failed to instance influxdb data point")
			}
			points = append(points, point)
		}
	}

	// Batch write influxdb data points
	err := r.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.GpuPrediction),
	})
	if err != nil {
		return errors.Wrap(err, "failed to batch write influxdb data points")
	}

	return nil
}

func (r *TemperatureCelsiusRepository) ListPredictions(host, minorNumber, modelId, predictionId, granularity string, condition *DBCommon.QueryCondition) ([]*EntityInfluxGpuPrediction.TemperatureCelsiusEntity, error) {
	entities := make([]*EntityInfluxGpuPrediction.TemperatureCelsiusEntity, 0)

	influxdbStatement := InfluxDB.Statement{
		QueryCondition: condition,
		Measurement:    TemperatureCelsius,
		GroupByTags:    []string{"host", "uuid"},
	}

	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.TemperatureCelsiusHost, "=", host)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.TemperatureCelsiusMinorNumber, "=", minorNumber)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.TemperatureCelsiusModelId, "=", modelId)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.TemperatureCelsiusPredictionId, "=", predictionId)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.TemperatureCelsiusGranularity, "=", granularity)
	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	response, err := r.influxDB.QueryDB(cmd, string(RepoInflux.GpuPrediction))
	if err != nil {
		return entities, errors.Wrap(err, "failed to list nvidia gpu temperature celsius predictions")
	}

	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			group := result.GetGroup(i)
			for j := 0; j < group.GetRowNum(); j++ {
				entity := EntityInfluxGpuPrediction.NewTemperatureCelsiusEntityFromMap(group.GetRow(j))
				entities = append(entities, &entity)
			}
		}
	}

	return entities, nil
}
