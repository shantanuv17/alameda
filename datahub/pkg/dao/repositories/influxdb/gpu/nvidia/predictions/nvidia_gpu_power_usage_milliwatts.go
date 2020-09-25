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

type PowerUsageMilliWattsRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewPowerUsageMilliWattsRepositoryWithConfig(cfg InfluxDB.Config) *PowerUsageMilliWattsRepository {
	return &PowerUsageMilliWattsRepository{
		influxDB: InfluxDB.NewClient(&cfg),
	}
}

func (r *PowerUsageMilliWattsRepository) CreatePredictions(predictions []*DaoGpu.GpuPrediction) error {
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
				EntityInfluxGpuPrediction.PowerUsageMilliWattsHost:        prediction.Metadata.Host,
				EntityInfluxGpuPrediction.PowerUsageMilliWattsName:        prediction.Name,
				EntityInfluxGpuPrediction.PowerUsageMilliWattsUuid:        prediction.Uuid,
				EntityInfluxGpuPrediction.PowerUsageMilliWattsGranularity: strconv.FormatInt(granularity, 10),
			}

			// Pack influx fields
			fields := map[string]interface{}{
				EntityInfluxGpuPrediction.PowerUsageMilliWattsModelId:      metric.ModelId,
				EntityInfluxGpuPrediction.PowerUsageMilliWattsPredictionId: metric.PredictionId,
				EntityInfluxGpuPrediction.PowerUsageMilliWattsMinorNumber:  prediction.Metadata.MinorNumber,
				EntityInfluxGpuPrediction.PowerUsageMilliWattsValue:        valueInFloat64,
			}

			// Add to influx point list
			point, err := InfluxClient.NewPoint(string(PowerUsageMilliWatts), tags, fields, metric.Timestamp)
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

func (r *PowerUsageMilliWattsRepository) ListPredictions(host, minorNumber, modelId, predictionId, granularity string, condition *DBCommon.QueryCondition) ([]*EntityInfluxGpuPrediction.PowerUsageMilliWattsEntity, error) {
	entities := make([]*EntityInfluxGpuPrediction.PowerUsageMilliWattsEntity, 0)

	influxdbStatement := InfluxDB.Statement{
		QueryCondition: condition,
		Measurement:    PowerUsageMilliWatts,
		GroupByTags:    []string{"host", "uuid"},
	}

	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.PowerUsageMilliWattsHost, "=", host)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.PowerUsageMilliWattsMinorNumber, "=", minorNumber)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.PowerUsageMilliWattsModelId, "=", modelId)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.PowerUsageMilliWattsPredictionId, "=", predictionId)
	influxdbStatement.AppendWhereClause("AND", EntityInfluxGpuPrediction.PowerUsageMilliWattsGranularity, "=", granularity)
	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	response, err := r.influxDB.QueryDB(cmd, string(RepoInflux.GpuPrediction))
	if err != nil {
		return entities, errors.Wrap(err, "failed to list nvidia gpu power usage milli watts predictions")
	}

	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			group := result.GetGroup(i)
			for j := 0; j < group.GetRowNum(); j++ {
				entity := EntityInfluxGpuPrediction.NewPowerUsageMilliWattsEntityFromMap(group.GetRow(j))
				entities = append(entities, &entity)
			}
		}
	}

	return entities, nil
}
