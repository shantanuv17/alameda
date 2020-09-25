package predictions

import (
	DaoPredictionTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/resources"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
)

type ClusterPredictionExtended struct {
	*DaoPredictionTypes.ClusterPrediction
}

func (d *ClusterPredictionExtended) ProducePredictions() *ApiPredictions.ClusterPrediction {
	var (
		rawDataChan        = make(chan *ApiPredictions.MetricData)
		upperBoundDataChan = make(chan *ApiPredictions.MetricData)
		lowerBoundDataChan = make(chan *ApiPredictions.MetricData)
		numOfGoroutine     = 0

		datahubClusterPrediction ApiPredictions.ClusterPrediction
	)

	datahubClusterPrediction = ApiPredictions.ClusterPrediction{
		ObjectMeta: resources.NewObjectMeta(&d.ObjectMeta),
	}

	// Handle prediction raw data
	numOfGoroutine = 0
	for metricType, samples := range d.PredictionRaw {
		if datahubMetricType, exist := FormatEnum.TypeToDatahubMetricType[metricType]; exist {
			numOfGoroutine++
			go common.ProducePredictionMetricDataFromSamples(datahubMetricType, samples.Granularity, samples.Data, rawDataChan)
		}
	}
	for i := 0; i < numOfGoroutine; i++ {
		receivedPredictionData := <-rawDataChan
		datahubClusterPrediction.PredictedRawData = append(datahubClusterPrediction.PredictedRawData, receivedPredictionData)
	}

	// Handle prediction upper bound data
	numOfGoroutine = 0
	for metricType, samples := range d.PredictionUpperBound {
		if datahubMetricType, exist := FormatEnum.TypeToDatahubMetricType[metricType]; exist {
			numOfGoroutine++
			go common.ProducePredictionMetricDataFromSamples(datahubMetricType, samples.Granularity, samples.Data, upperBoundDataChan)
		}
	}
	for i := 0; i < numOfGoroutine; i++ {
		receivedPredictionData := <-upperBoundDataChan
		datahubClusterPrediction.PredictedUpperboundData = append(datahubClusterPrediction.PredictedUpperboundData, receivedPredictionData)
	}

	// Handle prediction lower bound data
	numOfGoroutine = 0
	for metricType, samples := range d.PredictionLowerBound {
		if datahubMetricType, exist := FormatEnum.TypeToDatahubMetricType[metricType]; exist {
			numOfGoroutine++
			go common.ProducePredictionMetricDataFromSamples(datahubMetricType, samples.Granularity, samples.Data, lowerBoundDataChan)
		}
	}
	for i := 0; i < numOfGoroutine; i++ {
		receivedPredictionData := <-lowerBoundDataChan
		datahubClusterPrediction.PredictedLowerboundData = append(datahubClusterPrediction.PredictedLowerboundData, receivedPredictionData)
	}

	return &datahubClusterPrediction
}
