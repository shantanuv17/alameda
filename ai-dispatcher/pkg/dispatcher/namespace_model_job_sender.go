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

type namespaceModelJobSender struct {
	datahubClient  *datahubpkg.Client
	modelMapper    *ModelMapper
	metricExporter *metrics.Exporter
}

func NewNamespaceModelJobSender(datahubClient *datahubpkg.Client, modelMapper *ModelMapper,
	metricExporter *metrics.Exporter) *namespaceModelJobSender {
	return &namespaceModelJobSender{
		datahubClient:  datahubClient,
		modelMapper:    modelMapper,
		metricExporter: metricExporter,
	}
}

func (sender *namespaceModelJobSender) sendModelJobs(namespaces []*datahub_resources.Namespace,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	for _, namespace := range namespaces {
		sender.sendNamespaceModelJobs(namespace, queueSender, pdUnit, granularity, predictionStep, &wg)
		err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.model.directory"), granularity))
		if err != nil {
			scope.Error(err.Error())
		}
	}
}

func (sender *namespaceModelJobSender) sendNamespaceModelJobs(namespace *datahub_resources.Namespace,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64, wg *sync.WaitGroup) {
	dataGranularity := queue.GetGranularityStr(granularity)
	namespaceName := namespace.GetObjectMeta().GetName()
	lastPredictionMetrics, err := sender.getLastMIdPrediction(sender.datahubClient, namespace, granularity)
	if err != nil {
		scope.Infof("[NAMESPACE][%s][%s] Get last prediction failed: %s",
			dataGranularity, namespaceName, err.Error())
		return
	}

	sender.sendJobByMetrics(namespace, queueSender, pdUnit, granularity, predictionStep,
		sender.datahubClient, lastPredictionMetrics)
}

func (sender *namespaceModelJobSender) sendJob(namespace *datahub_resources.Namespace,
	queueSender queue.QueueSender, pdUnit string, granularity int64,
	metricType datahub_common.MetricType) {
	clusterID := namespace.GetObjectMeta().GetClusterName()
	namespaceName := namespace.GetObjectMeta().GetName()
	dataGranularity := queue.GetGranularityStr(granularity)
	marshaler := jsonpb.Marshaler{}
	namespaceStr, err := marshaler.MarshalToString(namespace)
	if err != nil {
		scope.Errorf("[NAMESPACE][%s][%s] Encode pb message failed. %s",
			dataGranularity, namespaceName, err.Error())
		return
	}

	jb := queue.NewJobBuilder(clusterID, pdUnit, granularity, metricType, namespaceStr, nil)
	jobJSONStr, err := jb.GetJobJSONString()
	if err != nil {
		scope.Errorf(
			"[NAMESPACE][%s][%s] Prepare model job payload failed. %s",
			dataGranularity, namespaceName, err.Error())
		return
	}

	nsJobStr := fmt.Sprintf("%s/%s/%s/%v/%s", consts.UnitTypeNamespace, clusterID, namespaceName, granularity, metricType)
	scope.Infof("[NAMESPACE][%s][%s] Try to send namespace model job: %s",
		dataGranularity, namespaceName, nsJobStr)
	err = queueSender.SendJsonString(modelQueueName, jobJSONStr, nsJobStr, granularity)
	if err == nil {
		sender.modelMapper.AddModelInfo(clusterID, pdUnit, dataGranularity, metricType.String(), map[string]string{
			"name": namespaceName,
		})
	} else {
		scope.Errorf(
			"[NAMESPACE][%s][%s] Send model job payload failed. %s",
			dataGranularity, namespaceName, err.Error())
	}

}

func (sender *namespaceModelJobSender) getLastMIdPrediction(datahubServiceClnt *datahubpkg.Client,
	namespace *datahub_resources.Namespace, granularity int64) ([]*datahub_predictions.MetricData, error) {

	metricData := []*datahub_predictions.MetricData{}
	dataGranularity := queue.GetGranularityStr(granularity)
	namespaceName := namespace.GetObjectMeta().GetName()

	namespacePredictRes, err := datahubServiceClnt.ListNamespacePredictions(
		&datahub_predictions.ListNamespacePredictionsRequest{
			ObjectMeta: []*datahub_resources.ObjectMeta{
				{
					Name: namespaceName,
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
	if len(namespacePredictRes.GetNamespacePredictions()) == 0 {
		return metricData, nil
	}

	lastNamespacePrediction := namespacePredictRes.GetNamespacePredictions()[0]
	lnsPDRData := lastNamespacePrediction.GetPredictedRawData()
	if lnsPDRData == nil {
		return metricData, nil
	}

	for _, pdRD := range lnsPDRData {
		for _, theData := range pdRD.GetData() {
			lastMid = theData.GetModelId()
			break
		}
		if lastMid == "" {
			scope.Warnf("[NAMESPACE][%s][%s] Query last model id for metric %s is empty",
				dataGranularity, namespaceName, pdRD.GetMetricType())
		}

		namespacePredictRes, err = datahubServiceClnt.ListNamespacePredictions(
			&datahub_predictions.ListNamespacePredictionsRequest{
				ObjectMeta: []*datahub_resources.ObjectMeta{
					{
						Name: namespaceName,
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
			scope.Errorf("[NAMESPACE][%s][%s] Query last model id %v for metric %s failed",
				dataGranularity, namespaceName, lastMid, pdRD.GetMetricType())
			continue
		}

		for _, nsPrediction := range namespacePredictRes.GetNamespacePredictions() {
			for _, lMIDPdRD := range nsPrediction.GetPredictedRawData() {
				if lMIDPdRD.GetMetricType() == pdRD.GetMetricType() {
					metricData = append(metricData, lMIDPdRD)
				}
			}
		}
	}

	return metricData, nil
}

func (sender *namespaceModelJobSender) getQueryMetricStartTime(
	metricData *datahub_predictions.MetricData) int64 {
	mD := metricData.GetData()
	if len(mD) > 0 {
		return mD[len(mD)-1].GetTime().GetSeconds()
	}

	return 0
}

func (sender *namespaceModelJobSender) sendJobByMetrics(namespace *datahub_resources.Namespace, queueSender queue.QueueSender,
	pdUnit string, granularity int64, predictionStep int64, datahubServiceClnt *datahubpkg.Client,
	lastPredictionMetrics []*datahub_predictions.MetricData) {

	clusterID := namespace.GetObjectMeta().GetClusterName()
	namespaceName := namespace.GetObjectMeta().GetName()
	dataGranularity := queue.GetGranularityStr(granularity)
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
			scope.Infof("[NAMESPACE][%s][%s] No prediction metric %s found, send model jobs.",
				dataGranularity, namespaceName, metricType)
			sender.sendJob(namespace, queueSender, pdUnit, granularity, metricType)
		}
	}

	for _, lastPredictionMetric := range lastPredictionMetrics {
		if len(lastPredictionMetric.GetData()) == 0 {
			scope.Infof("[NAMESPACE][%s][%s] No prediction metric %s found, send model jobs",
				dataGranularity, namespaceName, lastPredictionMetric.GetMetricType().String())
			sender.sendJob(namespace, queueSender, pdUnit, granularity, lastPredictionMetric.GetMetricType())
			continue
		} else {
			lastPrediction := lastPredictionMetric.GetData()[0]
			lastPredictionTime := lastPredictionMetric.GetData()[0].GetTime().GetSeconds()
			if lastPrediction != nil && lastPredictionTime <= nowSeconds {
				scope.Infof("[NAMESPACE][%s][%s] Send model job due to no predict metric %s found or is out of date",
					dataGranularity, namespaceName, lastPredictionMetric.GetMetricType().String())
				sender.sendJob(namespace, queueSender, pdUnit, granularity, lastPredictionMetric.GetMetricType())
				continue
			}

			mID := lastPredictionMetric.GetData()[0].ModelId
			modelMaxUsedTimes := viper.GetInt64(fmt.Sprintf(
				"granularities.%s.modelMaxUsedTimes", utils.GetGranularityStr(granularity)))
			if mID != "" && utils.IsModelExpired(
				mID, granularity, modelMaxUsedTimes) {
				scope.Infof("[NAMESPACE][%s][%s] Send model job due to the model (id: %s, model max used times: %d, now: %d) of metric %s is expired",
					dataGranularity, namespaceName, mID, modelMaxUsedTimes, time.Now().Unix(), lastPredictionMetric.GetMetricType().String())
				sender.sendJob(namespace, queueSender, pdUnit, granularity, lastPredictionMetric.GetMetricType())
				continue
			}

			queryStartTime := time.Now().Unix() - predictionStep*granularity
			firstPDTime := sender.getQueryMetricStartTime(lastPredictionMetric)
			if firstPDTime > 0 && firstPDTime <= time.Now().Unix() {
				queryStartTime = firstPDTime
			}

			namespaceMetricsRes, err := datahubServiceClnt.ListNamespaceMetrics(
				&datahub_metrics.ListNamespaceMetricsRequest{
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
							Name: namespaceName,
						},
					},
					MetricTypes: []datahub_common.MetricType{
						lastPredictionMetric.GetMetricType(),
					},
				})
			if err != nil {
				scope.Errorf("[NAMESPACE][%s][%s] List metric for sending model job failed: %s",
					dataGranularity, namespaceName, err.Error())
				continue
			}
			namespaceMetrics := namespaceMetricsRes.GetNamespaceMetrics()
			predictRawData := lastPredictionMetrics
			for _, predictRawDatum := range predictRawData {
				for _, namespaceMetric := range namespaceMetrics {
					metricData := namespaceMetric.GetMetricData()
					for _, metricDatum := range metricData {
						mData := metricDatum.GetData()
						pData := []*datahub_predictions.Sample{}
						if metricDatum.GetMetricType() == predictRawDatum.GetMetricType() {
							pData = append(pData, predictRawDatum.GetData()...)
							metricsNeedToModel, drift := DriftEvaluation(consts.UnitTypeNamespace, predictRawDatum.GetMetricType(),
								granularity, mData, pData, map[string]string{
									"clusterID":         clusterID,
									"namespaceName":     namespaceName,
									"targetDisplayName": fmt.Sprintf("[NAMESPACE][%s][%s]", dataGranularity, namespaceName),
								}, sender.metricExporter)
							for _, mntm := range metricsNeedToModel {
								if drift {
									scope.Infof("[NAMESPACE][%s][%s] Export metric %s drift counter",
										dataGranularity, namespaceName, mntm)
									sender.metricExporter.AddNamespaceMetricDrift(clusterID, namespaceName,
										queue.GetGranularityStr(granularity), mntm.String(), time.Now().Unix(), 1.0)
								}
								isModeling := sender.modelMapper.IsModeling(clusterID, pdUnit, dataGranularity, mntm.String(), map[string]string{
									"name": namespaceName,
								})
								if !isModeling || (isModeling && sender.modelMapper.IsModelTimeout(
									clusterID, pdUnit, dataGranularity, mntm.String(), map[string]string{
										"name": namespaceName,
									})) {
									sender.sendJob(namespace, queueSender, pdUnit, granularity, mntm)
								}
							}
						}
					}
				}
			}
		}
	}
}
