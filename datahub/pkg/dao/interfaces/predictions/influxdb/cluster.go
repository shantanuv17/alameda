package influxdb

import (
	DaoPredictionTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
	RepoInfluxPrediction "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/predictions"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
)

type ClusterPredictions struct {
	InfluxDBConfig InfluxDB.Config
}

func NewClusterPredictionsWithConfig(config InfluxDB.Config) DaoPredictionTypes.ClusterPredictionsDAO {
	return &ClusterPredictions{InfluxDBConfig: config}
}

// CreateClusterPredictions Implementation of prediction dao interface
func (p *ClusterPredictions) CreatePredictions(predictions DaoPredictionTypes.ClusterPredictionMap) error {
	predictionRepo := RepoInfluxPrediction.NewClusterRepositoryWithConfig(p.InfluxDBConfig)

	err := predictionRepo.CreatePredictions(predictions)
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *ClusterPredictions) ListPredictions(request DaoPredictionTypes.ListClusterPredictionsRequest) (DaoPredictionTypes.ClusterPredictionMap, error) {
	clusterPredictionMap := DaoPredictionTypes.NewClusterPredictionMap()

	predictionRepo := RepoInfluxPrediction.NewClusterRepositoryWithConfig(p.InfluxDBConfig)
	clusterPredictions, err := predictionRepo.ListPredictions(request)
	if err != nil {
		scope.Error(err.Error())
		return DaoPredictionTypes.NewClusterPredictionMap(), err
	}
	for _, clusterPrediction := range clusterPredictions {
		clusterPredictionMap.AddClusterPrediction(clusterPrediction)
	}

	return clusterPredictionMap, nil
}
