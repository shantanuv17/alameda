package dispatcher

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/viper"
	"prophetstor.com/alameda/ai-dispatcher/consts"
	"prophetstor.com/alameda/ai-dispatcher/pkg/metrics"
	"prophetstor.com/alameda/ai-dispatcher/pkg/queue"
	utils "prophetstor.com/alameda/ai-dispatcher/pkg/utils"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	datahub_common "prophetstor.com/api/datahub/common"
	datahub_metrics "prophetstor.com/api/datahub/metrics"
	datahub_predictions "prophetstor.com/api/datahub/predictions"
	datahub_resources "prophetstor.com/api/datahub/resources"
)

type controllerModelJobSender struct {
	datahubClient  *datahubpkg.Client
	modelMapper    *ModelMapper
	metricExporter *metrics.Exporter
}

func NewControllerModelJobSender(datahubClient *datahubpkg.Client, modelMapper *ModelMapper,
	metricExporter *metrics.Exporter) *controllerModelJobSender {
	return &controllerModelJobSender{
		datahubClient:  datahubClient,
		modelMapper:    modelMapper,
		metricExporter: metricExporter,
	}
}

func (sender *controllerModelJobSender) sendModelJobs(controllers []*datahub_resources.Controller,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	for _, controller := range controllers {
		sender.sendControllerModelJobs(controller, queueSender, pdUnit, granularity, predictionStep, &wg)
		err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.model.directory"), granularity))
		if err != nil {
			scope.Error(err.Error())
		}
	}
}

func (sender *controllerModelJobSender) sendControllerModelJobs(controller *datahub_resources.Controller,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64, wg *sync.WaitGroup) {
	dataGranularity := queue.GetGranularityStr(granularity)
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetName()

	lastPredictionMetrics, err := sender.getLastMIdPrediction(sender.datahubClient, controller, granularity)
	if err != nil {
		scope.Infof("[CONTROLLER][%s][%s][%s/%s] Get last prediction failed: %s",
			controller.GetKind().String(), dataGranularity, controllerNS, controllerName, err.Error())
		return
	}
	sender.sendJobByMetrics(controller, queueSender, pdUnit, granularity, predictionStep,
		sender.datahubClient, lastPredictionMetrics)
}

func (sender *controllerModelJobSender) sendJob(controller *datahub_resources.Controller,
	queueSender queue.QueueSender, pdUnit string, granularity int64,
	metricType datahub_common.MetricType) {
	marshaler := jsonpb.Marshaler{}
	dataGranularity := queue.GetGranularityStr(granularity)
	clusterID := controller.GetObjectMeta().GetClusterName()
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetName()
	controllerKindStr := controller.GetKind().String()
	controllerStr, err := marshaler.MarshalToString(controller)
	if err != nil {
		scope.Errorf("[CONTROLLER][%s][%s][%s/%s] Encode pb message failed. %s",
			controllerKindStr, dataGranularity, controllerNS, controllerName, err.Error())
		return
	}

	jb := queue.NewJobBuilder(clusterID, pdUnit, granularity, metricType, controllerStr, nil)
	jobJSONStr, err := jb.GetJobJSONString()
	if err != nil {
		scope.Errorf(
			"[CONTROLLER][%s][%s][%s/%s] Prepare model job payload failed. %s",
			controllerKindStr, dataGranularity, controllerNS, controllerName, err.Error())
		return
	}

	controllerJobStr := fmt.Sprintf("%s/%s/%s/%s/%s/%v/%s", consts.UnitTypeController,
		clusterID, controllerKindStr, controllerNS, controllerName, granularity, metricType)
	scope.Infof("[CONTROLLER][%s][%s][%s/%s] Try to send controller model job: %s",
		controllerKindStr, dataGranularity, controllerNS, controllerName, controllerJobStr)
	err = queueSender.SendJsonString(modelQueueName, jobJSONStr, controllerJobStr, granularity)
	if err == nil {
		sender.modelMapper.AddModelInfo(clusterID, pdUnit, dataGranularity, metricType.String(), map[string]string{
			"namespace": controllerNS,
			"name":      controllerName,
			"kind":      controllerKindStr,
		})
	} else {
		scope.Errorf(
			"[CONTROLLER][%s][%s][%s/%s] Send model job payload failed. %s",
			controllerKindStr, dataGranularity, controllerNS, controllerName, err.Error())
	}

}

func (sender *controllerModelJobSender) getLastMIdPrediction(datahubServiceClnt *datahubpkg.Client,
	controller *datahub_resources.Controller, granularity int64) ([]*datahub_predictions.MetricData, error) {

	metricData := []*datahub_predictions.MetricData{}
	dataGranularity := queue.GetGranularityStr(granularity)
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetName()

	controllerPredictRes, err := datahubServiceClnt.ListControllerPredictions(
		&datahub_predictions.ListControllerPredictionsRequest{
			ObjectMeta: []*datahub_resources.ObjectMeta{
				{
					Namespace: controllerNS,
					Name:      controllerName,
				},
			},
			Granularity: granularity,
			QueryCondition: &datahub_common.QueryCondition{
				Limit: 1,
				Order: datahub_common.QueryCondition_DESC,
				TimeRange: &datahub_common.TimeRange{
					Step: &duration.Duration{
						Seconds: granularity,
					},
				},
			},
		})
	if err != nil {
		return metricData, err
	}

	lastMid := ""
	if len(controllerPredictRes.GetControllerPredictions()) == 0 {
		return metricData, nil
	}

	lastControllerPrediction := controllerPredictRes.GetControllerPredictions()[0]
	lctrlPDRData := lastControllerPrediction.GetPredictedRawData()
	if lctrlPDRData == nil {
		lctrlPDRData = lastControllerPrediction.GetPredictedLowerboundData()
	}
	if lctrlPDRData == nil {
		lctrlPDRData = lastControllerPrediction.GetPredictedUpperboundData()
	}
	for _, pdRD := range lctrlPDRData {
		for _, theData := range pdRD.GetData() {
			lastMid = theData.GetModelId()
			break
		}
		if lastMid == "" {
			scope.Warnf("[CONTROLLER][%s][%s][%s/%s] Query last model id for metric %s is empty",
				controller.GetKind().String(), dataGranularity, controllerNS, controllerName, pdRD.GetMetricType())
		}

		controllerPredictRes, err = datahubServiceClnt.ListControllerPredictions(
			&datahub_predictions.ListControllerPredictionsRequest{
				ObjectMeta: []*datahub_resources.ObjectMeta{
					{
						Namespace: controllerNS,
						Name:      controllerName,
					},
				},
				Granularity: granularity,
				QueryCondition: &datahub_common.QueryCondition{
					Order: datahub_common.QueryCondition_DESC,
					TimeRange: &datahub_common.TimeRange{
						Step: &duration.Duration{
							Seconds: granularity,
						},
					},
				},
				ModelId: lastMid,
			})

		if err != nil {
			scope.Warnf("[CONTROLLER][%s][%s][%s/%s] Query last model id %v for metric %s failed",
				controller.GetKind().String(), dataGranularity, controllerNS, controllerName, lastMid, pdRD.GetMetricType())
		}

		for _, ctrlPrediction := range controllerPredictRes.GetControllerPredictions() {
			for _, lMIDPdRD := range ctrlPrediction.GetPredictedRawData() {
				if lMIDPdRD.GetMetricType() == pdRD.GetMetricType() {
					metricData = append(metricData, lMIDPdRD)
				}
			}
		}
	}

	return metricData, nil
}

func (sender *controllerModelJobSender) getQueryMetricStartTime(
	metricData *datahub_predictions.MetricData) int64 {
	mD := metricData.GetData()
	if len(mD) > 0 {
		return mD[len(mD)-1].GetTime().GetSeconds()
	}
	return 0
}

func (sender *controllerModelJobSender) sendJobByMetrics(controller *datahub_resources.Controller, queueSender queue.QueueSender,
	pdUnit string, granularity int64, predictionStep int64, datahubServiceClnt *datahubpkg.Client,
	lastPredictionMetrics []*datahub_predictions.MetricData) {

	dataGranularity := queue.GetGranularityStr(granularity)
	clusterID := controller.GetObjectMeta().GetClusterName()
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetName()
	nowSeconds := time.Now().Unix()

	for _, metricType := range []datahub_common.MetricType{
		datahub_common.MetricType_MEMORY_BYTES_USAGE,
		datahub_common.MetricType_CPU_MILLICORES_USAGE,
	} {
		mtNotFound := true
		for _, lastPredictionMetric := range lastPredictionMetrics {
			if lastPredictionMetric.GetMetricType() == metricType {
				mtNotFound = false
				break
			}
		}
		if mtNotFound {
			scope.Infof("[CONTROLLER][%s][%s][%s/%s] No prediction metric %s found, send model jobs",
				controller.GetKind().String(), dataGranularity, controllerNS, controllerName, metricType)
			sender.sendJob(controller, queueSender, pdUnit, granularity, metricType)
		}

	}

	for _, lastPredictionMetric := range lastPredictionMetrics {
		if len(lastPredictionMetric.GetData()) == 0 {
			scope.Infof("[CONTROLLER][%s][%s][%s/%s] No prediction metric %s found, send model jobs",
				controller.GetKind().String(), dataGranularity, controllerNS, controllerName, lastPredictionMetric.GetMetricType().String())
			sender.sendJob(controller, queueSender, pdUnit, granularity, lastPredictionMetric.GetMetricType())
			continue
		} else {
			lastPrediction := lastPredictionMetric.GetData()[0]
			lastPredictionTime := lastPredictionMetric.GetData()[0].GetTime().GetSeconds()

			if lastPrediction != nil && lastPredictionTime <= nowSeconds {
				scope.Infof("[CONTROLLER][%s][%s][%s/%s] Send model job due to no predict metric %s found or is out of date",
					controller.GetKind().String(), dataGranularity, controllerNS, controllerName, lastPredictionMetric.GetMetricType())
				sender.sendJob(controller, queueSender, pdUnit, granularity, lastPredictionMetric.GetMetricType())
				continue
			}

			mID := lastPredictionMetric.GetData()[0].ModelId
			modelMaxUsedTimes := viper.GetInt64(fmt.Sprintf(
				"granularities.%s.modelMaxUsedTimes", utils.GetGranularityStr(granularity)))
			if mID != "" && utils.IsModelExpired(
				mID, granularity, modelMaxUsedTimes) {
				scope.Infof("[CONTROLLER][%s][%s][%s/%s] Send model job due to the model (id: %s, model max used times: %d, now: %d) of metric %s is expired",
					controller.GetKind().String(), dataGranularity, controllerNS, controllerName, mID, modelMaxUsedTimes, time.Now().Unix(), lastPredictionMetric.GetMetricType())
				sender.sendJob(controller, queueSender, pdUnit, granularity, lastPredictionMetric.GetMetricType())
				continue
			}

			queryStartTime := time.Now().Unix() - predictionStep*granularity
			firstPDTime := sender.getQueryMetricStartTime(lastPredictionMetric)
			if firstPDTime > 0 && firstPDTime <= time.Now().Unix() {
				queryStartTime = firstPDTime
			}

			controllerMetricsRes, err := datahubServiceClnt.ListControllerMetrics(
				&datahub_metrics.ListControllerMetricsRequest{
					QueryCondition: &datahub_common.QueryCondition{
						Order: datahub_common.QueryCondition_DESC,
						TimeRange: &datahub_common.TimeRange{
							StartTime: &timestamp.Timestamp{
								Seconds: queryStartTime,
							},
							Step: &duration.Duration{
								Seconds: granularity,
							},
							AggregateFunction: datahub_common.TimeRange_AVG,
						},
					},
					ObjectMeta: []*datahub_resources.ObjectMeta{
						{
							Namespace: controllerNS,
							Name:      controllerName,
						},
					},
					MetricTypes: []datahub_common.MetricType{
						lastPredictionMetric.GetMetricType(),
					},
				})

			if err != nil {
				scope.Errorf("[CONTROLLER][%s][%s][%s/%s] List metric for sending model job failed: %s",
					controller.GetKind().String(), dataGranularity, controllerNS, controllerName, err.Error())
				continue
			}
			controllerMetrics := controllerMetricsRes.GetControllerMetrics()

			predictRawData := lastPredictionMetrics
			for _, predictRawDatum := range predictRawData {
				for _, controllerMetric := range controllerMetrics {
					metricData := controllerMetric.GetMetricData()
					for _, metricDatum := range metricData {
						mData := metricDatum.GetData()
						pData := []*datahub_predictions.Sample{}
						if metricDatum.GetMetricType() == predictRawDatum.GetMetricType() {
							pData = append(pData, predictRawDatum.GetData()...)
							metricsNeedToModel, drift := DriftEvaluation(consts.UnitTypeController, predictRawDatum.GetMetricType(), granularity, mData, pData, map[string]string{
								"clusterID":      clusterID,
								"controllerNS":   controllerNS,
								"controllerName": controllerName,
								"controllerKind": controller.GetKind().String(),
								"targetDisplayName": fmt.Sprintf("[CONTROLLER][%s][%s][%s/%s]",
									controller.GetKind().String(), dataGranularity, controllerNS, controllerName),
							}, sender.metricExporter)

							for _, mntm := range metricsNeedToModel {
								if drift {
									scope.Infof("[CONTROLLER][%s][%s][%s/%s] Export metric %s drift counter",
										controller.GetKind().String(), dataGranularity, controllerNS, controllerName, mntm)
									sender.metricExporter.AddControllerMetricDrift(clusterID, controllerNS, controllerName,
										controller.GetKind().String(), queue.GetGranularityStr(granularity), mntm.String(), time.Now().Unix(), 1.0)
								}
								isModeling := sender.modelMapper.IsModeling(clusterID, pdUnit, dataGranularity, mntm.String(), map[string]string{
									"namespace": controllerNS,
									"name":      controllerName,
									"kind":      controller.GetKind().String(),
								})
								if !isModeling || (isModeling && sender.modelMapper.IsModelTimeout(
									clusterID, pdUnit, dataGranularity, mntm.String(), map[string]string{
										"namespace": controllerNS,
										"name":      controllerName,
										"kind":      controller.GetKind().String(),
									})) {
									sender.sendJob(controller, queueSender, pdUnit, granularity, mntm)
								}
							}
						}
					}
				}
			}
		}
	}
}
