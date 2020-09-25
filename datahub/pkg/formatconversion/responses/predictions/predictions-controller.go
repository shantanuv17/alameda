package predictions

import (
	DaoPredictionTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/resources"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
	ApiResources "prophetstor.com/api/datahub/resources"
)

type ControllerPredictionExtended struct {
	*DaoPredictionTypes.ControllerPrediction
}

func (d *ControllerPredictionExtended) ProducePredictions() *ApiPredictions.ControllerPrediction {
	var (
		rawDataChan        = make(chan *ApiPredictions.MetricData)
		upperBoundDataChan = make(chan *ApiPredictions.MetricData)
		lowerBoundDataChan = make(chan *ApiPredictions.MetricData)
		numOfGoroutine     = 0

		datahubControllerPrediction ApiPredictions.ControllerPrediction
	)

	var ctlKind ApiResources.Kind
	if value, ok := ApiResources.Kind_value[d.Kind]; ok {
		ctlKind = ApiResources.Kind(value)
	}

	datahubControllerPrediction = ApiPredictions.ControllerPrediction{
		ObjectMeta: resources.NewObjectMeta(&d.ObjectMeta),
		Kind:       ctlKind,
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
		datahubControllerPrediction.PredictedRawData = append(datahubControllerPrediction.PredictedRawData, receivedPredictionData)
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
		datahubControllerPrediction.PredictedUpperboundData = append(datahubControllerPrediction.PredictedUpperboundData, receivedPredictionData)
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
		datahubControllerPrediction.PredictedLowerboundData = append(datahubControllerPrediction.PredictedLowerboundData, receivedPredictionData)
	}

	return &datahubControllerPrediction
}
