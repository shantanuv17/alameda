package predictions

import (
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	EntityInfluxGpuPrediction "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/gpu/nvidia/predictions"
	DaoGpu "prophetstor.com/alameda/datahub/pkg/dao/interfaces/gpu/influxdb"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	DatahubUtils "prophetstor.com/alameda/datahub/pkg/utils"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	"strconv"
)

type MemoryUsagePercentageRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewMemoryUsagePercentageRepositoryWithConfig(cfg InfluxDB.Config) *MemoryUsagePercentageRepository {
	return &MemoryUsagePercentageRepository{
		influxDB: InfluxDB.NewClient(&cfg),
	}
}

func (r *MemoryUsagePercentageRepository) CreatePredictions(predictions []*DaoGpu.GpuPrediction) error {
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
				EntityInfluxGpuPrediction.MemoryUsageHost:        prediction.Metadata.Host,
				EntityInfluxGpuPrediction.MemoryUsageName:        prediction.Name,
				EntityInfluxGpuPrediction.MemoryUsageUuid:        prediction.Uuid,
				EntityInfluxGpuPrediction.MemoryUsageGranularity: strconv.FormatInt(granularity, 10),
			}

			// Pack influx fields
			fields := map[string]interface{}{
				EntityInfluxGpuPrediction.MemoryUsageModelId:      metric.ModelId,
				EntityInfluxGpuPrediction.MemoryUsagePredictionId: metric.PredictionId,
				EntityInfluxGpuPrediction.MemoryUsageMinorNumber:  prediction.Metadata.MinorNumber,
				EntityInfluxGpuPrediction.MemoryUsageValue:        valueInFloat64,
			}

			// Add to influx point list
			point, err := InfluxClient.NewPoint(string(MemoryUsagePercentage), tags, fields, metric.Timestamp)
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
