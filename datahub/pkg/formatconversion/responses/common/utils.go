package common

import (
	"github.com/golang/protobuf/ptypes"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	ApiCommon "prophetstor.com/api/datahub/common"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
)

func ProduceMetricDataFromSamples(metricType ApiCommon.MetricType, samples []FormatTypes.Sample, MetricDataChan chan<- *ApiCommon.MetricData) {
	datahubMetricData := &ApiCommon.MetricData{
		MetricType: metricType,
	}

	for _, sample := range samples {
		// TODO: Send error to caller
		googleTimestamp, err := ptypes.TimestampProto(sample.Timestamp)
		if err != nil {
			googleTimestamp = nil
		}

		datahubSample := ApiCommon.Sample{Time: googleTimestamp, NumValue: sample.Value}
		datahubMetricData.Data = append(datahubMetricData.Data, &datahubSample)
	}

	MetricDataChan <- datahubMetricData
}

func ProducePredictionMetricDataFromSamples(metricType ApiCommon.MetricType, granularity int64, samples []FormatTypes.PredictionSample, MetricDataChan chan<- *ApiPredictions.MetricData) {
	datahubMetricData := &ApiPredictions.MetricData{
		MetricType:  metricType,
		Granularity: granularity,
	}

	for _, sample := range samples {
		// TODO: Send error to caller
		googleTimestamp, err := ptypes.TimestampProto(sample.Timestamp)
		if err != nil {
			googleTimestamp = nil
		}

		datahubSample := ApiPredictions.Sample{Time: googleTimestamp, NumValue: sample.Value, ModelId: sample.ModelId, PredictionId: sample.PredictionId}
		datahubMetricData.Data = append(datahubMetricData.Data, &datahubSample)
	}

	MetricDataChan <- datahubMetricData
}

func MetricMapToDatahubMetricSlice(metricMap map[enumconv.MetricType][]FormatTypes.Sample) []*ApiCommon.MetricData {

	result := make([]*ApiCommon.MetricData, 0, len(metricMap))
	for metricType, samples := range metricMap {

		m := ApiCommon.MetricData{
			MetricType: enumconv.TypeToDatahubMetricType[metricType],
			Data:       make([]*ApiCommon.Sample, len(samples)),
		}
		for i, sample := range samples {
			// TODO: Send error to caller
			t, err := ptypes.TimestampProto(sample.Timestamp)
			if err != nil {
				t = nil
			}
			m.Data[i] = &ApiCommon.Sample{
				Time:     t,
				NumValue: sample.Value,
			}
		}
		result = append(result, &m)
	}
	return result
}
