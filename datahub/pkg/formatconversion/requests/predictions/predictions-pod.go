package predictions

import (
	"github.com/golang/protobuf/ptypes"
	DaoPredictionTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/metrics"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/resources"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	Metadata "prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
)

type CreatePodPredictionsRequestExtended struct {
	*ApiPredictions.CreatePodPredictionsRequest
}

func (r *CreatePodPredictionsRequestExtended) Validate() error {
	return nil
}

func (r *CreatePodPredictionsRequestExtended) ProducePredictions() DaoPredictionTypes.PodPredictionMap {
	podPredictionMap := DaoPredictionTypes.NewPodPredictionMap()

	for _, pod := range r.GetPodPredictions() {
		podName := pod.GetObjectMeta().GetName()
		namespace := pod.GetObjectMeta().GetNamespace()
		nodeName := pod.GetObjectMeta().GetNodeName()
		clusterName := pod.GetObjectMeta().GetClusterName()

		podPrediction := DaoPredictionTypes.NewPodPrediction()
		podPrediction.ObjectMeta.Name = podName
		podPrediction.ObjectMeta.Namespace = namespace
		podPrediction.ObjectMeta.NodeName = nodeName
		podPrediction.ObjectMeta.ClusterName = clusterName

		for _, container := range pod.GetContainerPredictions() {
			containerName := container.GetName()

			containerPrediction := DaoPredictionTypes.NewContainerPrediction()
			containerPrediction.ContainerName = containerName
			containerPrediction.PodName = podName
			containerPrediction.Namespace = namespace
			containerPrediction.NodeName = nodeName
			containerPrediction.ClusterName = clusterName

			// Handle predicted raw data
			for _, data := range container.GetPredictedRawData() {
				metricType := metrics.MetricTypeNameMap[data.GetMetricType()]
				granularity := data.GetGranularity()
				for _, sample := range data.GetData() {
					timestamp, err := ptypes.Timestamp(sample.GetTime())
					if err != nil {
						scope.Error(" failed: " + err.Error())
					}
					sample := FormatTypes.PredictionSample{
						Timestamp:    timestamp,
						Value:        sample.GetNumValue(),
						ModelId:      sample.GetModelId(),
						PredictionId: sample.GetPredictionId(),
					}
					containerPrediction.AddRawSample(metricType, granularity, sample)
				}
			}

			// Handle predicted upper bound data
			for _, data := range container.GetPredictedUpperboundData() {
				metricType := metrics.MetricTypeNameMap[data.GetMetricType()]
				granularity := data.GetGranularity()
				for _, sample := range data.GetData() {
					timestamp, err := ptypes.Timestamp(sample.GetTime())
					if err != nil {
						scope.Error(" failed: " + err.Error())
					}
					sample := FormatTypes.PredictionSample{
						Timestamp:    timestamp,
						Value:        sample.GetNumValue(),
						ModelId:      sample.GetModelId(),
						PredictionId: sample.GetPredictionId(),
					}
					containerPrediction.AddUpperBoundSample(metricType, granularity, sample)
				}
			}

			// Handle predicted lower bound data
			for _, data := range container.GetPredictedLowerboundData() {
				metricType := metrics.MetricTypeNameMap[data.GetMetricType()]
				granularity := data.GetGranularity()
				for _, sample := range data.GetData() {
					timestamp, err := ptypes.Timestamp(sample.GetTime())
					if err != nil {
						scope.Error(" failed: " + err.Error())
					}
					sample := FormatTypes.PredictionSample{
						Timestamp:    timestamp,
						Value:        sample.GetNumValue(),
						ModelId:      sample.GetModelId(),
						PredictionId: sample.GetPredictionId(),
					}
					containerPrediction.AddLowerBoundSample(metricType, granularity, sample)
				}
			}

			podPrediction.ContainerPredictionMap.AddContainerPrediction(containerPrediction)
		}

		podPredictionMap.AddPodPrediction(podPrediction)
	}

	return podPredictionMap
}

type ListPodPredictionsRequestExtended struct {
	Request *ApiPredictions.ListPodPredictionsRequest
}

func (r *ListPodPredictionsRequestExtended) Validate() error {
	return nil
}

func (r *ListPodPredictionsRequestExtended) ProduceRequest() DaoPredictionTypes.ListPodPredictionsRequest {
	request := DaoPredictionTypes.NewListPodPredictionsRequest()
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	request.Granularity = 30
	request.FillDays = r.Request.GetFillDays()
	request.ModelId = r.Request.GetModelId()
	request.PredictionId = r.Request.GetPredictionId()
	if r.Request.GetGranularity() != 0 {
		request.Granularity = r.Request.GetGranularity()
	}
	if r.Request.GetObjectMeta() != nil {
		for _, meta := range r.Request.GetObjectMeta() {
			// Normalize request
			objectMeta := resources.NewObjectMeta(meta)

			if objectMeta.IsEmpty() {
				request.ObjectMeta = make([]Metadata.ObjectMeta, 0)
				return request
			}
			request.ObjectMeta = append(request.ObjectMeta, objectMeta)
		}
	}
	return request
}
