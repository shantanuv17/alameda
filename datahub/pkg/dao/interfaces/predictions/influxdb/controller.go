package influxdb

import (
	DaoPredictionTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/predictions/types"
	RepoInfluxPrediction "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/predictions"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
)

type ControllerPredictions struct {
	InfluxDBConfig InfluxDB.Config
}

func NewControllerPredictionsWithConfig(config InfluxDB.Config) DaoPredictionTypes.ControllerPredictionsDAO {
	return &ControllerPredictions{InfluxDBConfig: config}
}

// CreateControllerPredictions Implementation of prediction dao interface
func (p *ControllerPredictions) CreatePredictions(predictions DaoPredictionTypes.ControllerPredictionMap) error {
	predictionRepo := RepoInfluxPrediction.NewControllerRepositoryWithConfig(p.InfluxDBConfig)

	err := predictionRepo.CreatePredictions(predictions)
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *ControllerPredictions) ListPredictions(request DaoPredictionTypes.ListControllerPredictionsRequest) (DaoPredictionTypes.ControllerPredictionMap, error) {
	controllerPredictionMap := DaoPredictionTypes.NewControllerPredictionMap()

	predictionRepo := RepoInfluxPrediction.NewControllerRepositoryWithConfig(p.InfluxDBConfig)
	controllerPredictions, err := predictionRepo.ListPredictions(request)
	if err != nil {
		scope.Error(err.Error())
		return DaoPredictionTypes.NewControllerPredictionMap(), err
	}
	for _, controllerPrediction := range controllerPredictions {
		controllerPredictionMap.AddControllerPrediction(controllerPrediction)
	}

	return controllerPredictionMap, nil
}
