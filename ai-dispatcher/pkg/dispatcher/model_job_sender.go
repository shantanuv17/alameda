package dispatcher

import (
	"fmt"
	"strings"
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/metrics"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/queue"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/stats"
	utils "github.com/containers-ai/alameda/ai-dispatcher/pkg/utils"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_gpu "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/gpu"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/viper"
)

type modelJobSender struct {
	datahubClient  *datahubpkg.Client
	modelMapper    *ModelMapper
	metricExporter *metrics.Exporter

	podModelJobSender         *podModelJobSender
	nodeModelJobSender        *nodeModelJobSender
	gpuModelJobSender         *gpuModelJobSender
	applicationModelJobSender *applicationModelJobSender
	namespaceModelJobSender   *namespaceModelJobSender
	clusterModelJobSender     *clusterModelJobSender
	controllerModelJobSender  *controllerModelJobSender
}

func NewModelJobSender(datahubClient *datahubpkg.Client, modelMapper *ModelMapper,
	metricExporter *metrics.Exporter) *modelJobSender {

	return &modelJobSender{
		datahubClient:  datahubClient,
		modelMapper:    modelMapper,
		metricExporter: metricExporter,
		podModelJobSender: NewPodModelJobSender(datahubClient, modelMapper,
			metricExporter),
		nodeModelJobSender: NewNodeModelJobSender(datahubClient, modelMapper,
			metricExporter),
		gpuModelJobSender: NewGPUModelJobSender(datahubClient, modelMapper,
			metricExporter),
		applicationModelJobSender: NewApplicationModelJobSender(datahubClient, modelMapper,
			metricExporter),
		namespaceModelJobSender: NewNamespaceModelJobSender(datahubClient, modelMapper,
			metricExporter),
		clusterModelJobSender: NewClusterModelJobSender(datahubClient, modelMapper,
			metricExporter),
		controllerModelJobSender: NewControllerModelJobSender(datahubClient, modelMapper,
			metricExporter),
	}
}

func (dispatcher *modelJobSender) SendNodeModelJobs(nodes []*datahub_resources.Node,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.nodeModelJobSender.sendModelJobs(nodes, queueSender, pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendPodModelJobs(pods []*datahub_resources.Pod, queueSender queue.QueueSender,
	pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.podModelJobSender.sendModelJobs(pods, queueSender,
		pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendGPUModelJobs(gpus []*datahub_gpu.Gpu,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.gpuModelJobSender.sendModelJobs(gpus,
		queueSender, pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendApplicationModelJobs(applications []*datahub_resources.Application,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.applicationModelJobSender.sendModelJobs(applications,
		queueSender, pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendNamespaceModelJobs(namespaces []*datahub_resources.Namespace,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.namespaceModelJobSender.sendModelJobs(namespaces,
		queueSender, pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendClusterModelJobs(clusters []*datahub_resources.Cluster,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.clusterModelJobSender.sendModelJobs(clusters,
		queueSender, pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendControllerModelJobs(controllers []*datahub_resources.Controller,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {
	dispatcher.controllerModelJobSender.sendModelJobs(controllers,
		queueSender, pdUnit, granularity, predictionStep)
}

func (dispatcher *modelJobSender) SendModelJobs(rawData []*datahub_data.Rawdata, queueSender queue.QueueSender,
	unit *config.Unit, granularity int64) {

	for _, rawDatum := range rawData {
		err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.model.directory"), granularity))
		if err != nil {
			scope.Error(err.Error())
		}
		for _, grp := range rawDatum.GetGroups() {
			rawDatumColumns := grp.GetColumns()
			for _, row := range grp.GetRows() {
				if unit.UnitValueKeys != nil {
					scaledUnit, err := dispatcher.isUnitWatchedByScaler(unit, row.GetValues(), rawDatumColumns)
					if err != nil || !scaledUnit {
						jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
							datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
						if err != nil {
							scope.Errorf("[%s] Skip sending model job due to get alamedascaler information failed: %s",
								jobID, err.Error())
						} else {
							scope.Infof("[%s] Skip sending model job due to the unit is not watched by any alamedascaler",
								jobID)
						}
						continue
					}
				}

				readData := []*datahub_data.ReadData{}
				for _, metricType := range unit.MetricTypes {
					readData = append(readData, &datahub_data.ReadData{
						MetricType:       metricType,
						ResourceBoundary: datahub_common.ResourceBoundary_RESOURCE_RAW,
						QueryCondition: &datahub_common.QueryCondition{
							Order: datahub_common.QueryCondition_DESC,
							Limit: 1,
							WhereCondition: []*datahub_common.Condition{
								{
									Keys:      []string{unit.Prediction.PredictValueKeys.Granularity},
									Values:    []string{fmt.Sprintf("%v", granularity)},
									Operators: []string{"="},
									Types: []datahub_common.DataType{
										datahub_common.DataType_DATATYPE_STRING,
									},
								},
							},
						},
					})
				}
				readDataRes, err := utils.ReadData(dispatcher.datahubClient, &datahub_schemas.SchemaMeta{
					Scope:    unit.Prediction.Scope,
					Category: unit.Prediction.Category,
					Type:     unit.Prediction.Type,
				}, readData)
				if err != nil {
					jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
						datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
					scope.Errorf("[%s] List last prediction point failed: %s", jobID, err.Error())
					continue
				} else if readDataRes.GetData() == nil {
					for _, unitMetric := range unit.MetricTypes {
						jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
							unitMetric, granularity)
						scope.Infof("[%s] No prediction found, send model job.", jobID)
						err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
							unitMetric, granularity, queueSender, jobID)
						if err != nil {
							scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
						}
					}
					continue
				}

				for _, lastPredictPointRawData := range readDataRes.GetData().GetRawdata() {
					jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
						lastPredictPointRawData.GetMetricType(), granularity)
					if len(lastPredictPointRawData.GetGroups()) == 0 {
						scope.Infof("[%s] No prediction found, send model job.", jobID)
						err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
							lastPredictPointRawData.GetMetricType(), granularity, queueSender, jobID)
						if err != nil {
							scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
						}
						continue
					}
					for _, lastPredictPointRawDataGrp := range lastPredictPointRawData.GetGroups() {
						if len(lastPredictPointRawDataGrp.GetRows()) == 0 {
							scope.Infof("[%s] No prediction found, send model job.", jobID)
							err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
								lastPredictPointRawData.GetMetricType(), granularity, queueSender, jobID)
							if err != nil {
								scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
							}
							continue
						}
						if len(lastPredictPointRawDataGrp.GetRows()) > 1 {
							scope.Errorf("[%s] Get last prediction point but get more than one", jobID)
							continue
						}
						if len(lastPredictPointRawDataGrp.GetRows()) == 1 {
							if time.Now().Unix() >= lastPredictPointRawDataGrp.GetRows()[0].GetTime().GetSeconds() {
								scope.Infof("[%s] Prediction is out of date, send model job", jobID)
								err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
									lastPredictPointRawData.GetMetricType(), granularity, queueSender, jobID)
								if err != nil {
									scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
								}
								continue
							}
							modelID, err := utils.GetRowValue(lastPredictPointRawDataGrp.GetRows()[0].GetValues(),
								lastPredictPointRawDataGrp.GetColumns(), unit.Prediction.PredictValueKeys.ModelID)
							if err != nil {
								scope.Errorf("[%s] Get model ID from last predict point failed", jobID)
								continue
							}

							modelMaxUsedTimes := viper.GetInt64(fmt.Sprintf(
								"granularities.%s.modelMaxUsedTimes", utils.GetGranularityStr(granularity)))
							if modelID != "" && utils.IsModelExpired(
								modelID, granularity, modelMaxUsedTimes) {
								scope.Errorf("[%s] Send model job due to the model (id: %s, model max used times: %d, now: %d) is expired",
									jobID, modelID, modelMaxUsedTimes, time.Now().Unix())
								err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
									lastPredictPointRawData.GetMetricType(), granularity, queueSender, jobID)
								if err != nil {
									scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
								}
								continue
							}

							scope.Infof("[%s] Use model ID %s to query prediction series to measure dirft", jobID, modelID)
							dispatcher.driftEval(modelID, lastPredictPointRawData.GetMetricType(), rawData, queueSender, unit, granularity)
						}
					}
				}
			}
		}
	}
}

func (dispatcher *modelJobSender) driftEval(modelID string, metricType datahub_common.MetricType, rawData []*datahub_data.Rawdata, queueSender queue.QueueSender,
	unit *config.Unit, granularity int64) {
	for _, rawDatum := range rawData {
		for _, grp := range rawDatum.GetGroups() {
			rawDatumColumns := grp.GetColumns()
			for _, row := range grp.GetRows() {
				readData := []*datahub_data.ReadData{{
					MetricType:       metricType,
					ResourceBoundary: datahub_common.ResourceBoundary_RESOURCE_RAW,
					QueryCondition: &datahub_common.QueryCondition{
						Order: datahub_common.QueryCondition_DESC,
						WhereCondition: []*datahub_common.Condition{
							{
								Keys:      []string{unit.Prediction.PredictValueKeys.ModelID, unit.Prediction.PredictValueKeys.Granularity},
								Values:    []string{modelID, fmt.Sprintf("%v", granularity)},
								Operators: []string{"=", "="},
								Types: []datahub_common.DataType{
									datahub_common.DataType_DATATYPE_STRING,
									datahub_common.DataType_DATATYPE_STRING,
								},
							},
						},
					},
				}}

				readDataRes, err := utils.ReadData(dispatcher.datahubClient, &datahub_schemas.SchemaMeta{
					Scope:    unit.Prediction.Scope,
					Category: unit.Prediction.Category,
					Type:     unit.Prediction.Type,
				}, readData)
				if err != nil {
					jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
						datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
					scope.Errorf("[%s] List prediction with model id %s failed: %s", jobID, modelID, err.Error())
					continue
				} else if readDataRes.GetData() == nil {
					for _, unitMetric := range unit.MetricTypes {
						jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
							unitMetric, granularity)
						scope.Infof("[%s] No prediction found, send model job.", jobID)
						err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
							unitMetric, granularity, queueSender, jobID)
						if err != nil {
							scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
						}
					}
					continue
				}

				for _, modelIDPredictData := range readDataRes.GetData().GetRawdata() {
					jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
						modelIDPredictData.GetMetricType(), granularity)
					if len(modelIDPredictData.GetGroups()) == 0 {
						scope.Infof("[%s] No prediction found with model id %s, send model job.", jobID, modelID)
						err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
							modelIDPredictData.GetMetricType(), granularity, queueSender, jobID)
						if err != nil {
							scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
						}
						continue
					}
					for _, modelIDPredictDataGrp := range modelIDPredictData.GetGroups() {

						modelPredictRows := modelIDPredictDataGrp.GetRows()
						if len(modelPredictRows) == 0 {
							scope.Infof("[%s] No prediction found with model id %s, send model job.", jobID, modelID)

							err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
								modelIDPredictData.GetMetricType(), granularity, queueSender, jobID)
							if err != nil {
								scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
							}
							continue
						}
						metricStartT := modelPredictRows[len(modelPredictRows)-1].GetTime().GetSeconds()
						if metricStartT > time.Now().Unix() {
							scope.Errorf("[%s] Cannot query future metrics with timestamp %v", jobID, metricStartT)
							continue
						}

						whereVals := []string{}
						whereOps := []string{}
						whereTypes := []datahub_common.DataType{}
						toQueryMetric := true
						for _, idK := range unit.IDKeys {
							whereVal, err := utils.GetRowValue(row.GetValues(), rawDatumColumns, idK)
							if err != nil {
								scope.Errorf("[%s] Cannot query metric for drift evaluation due to %s", jobID, err.Error())
								toQueryMetric = false
								break
							}
							whereVals = append(whereVals, whereVal)
							whereOps = append(whereOps, "=")
							whereTypes = append(whereTypes, datahub_common.DataType_DATATYPE_STRING)
						}
						if !toQueryMetric {
							continue
						}
						metricReadData := []*datahub_data.ReadData{{
							MetricType: metricType,
							QueryCondition: &datahub_common.QueryCondition{
								Order:   datahub_common.QueryCondition_DESC,
								Selects: []string{unit.Metric.MetricValueKeys.Value},
								Groups:  unit.IDKeys,
								TimeRange: &datahub_common.TimeRange{
									Step: &duration.Duration{
										Seconds: granularity,
									},
									StartTime: &timestamp.Timestamp{
										Seconds: metricStartT,
									},
									AggregateFunction: unit.Metric.Aggregation,
								},
								WhereCondition: []*datahub_common.Condition{
									{
										Keys:      unit.IDKeys,
										Values:    whereVals,
										Operators: whereOps,
										Types:     whereTypes,
									},
								},
							},
						}}

						metricReadDataRes, err := utils.ReadData(dispatcher.datahubClient, &datahub_schemas.SchemaMeta{
							Scope:    unit.Metric.Scope,
							Category: unit.Metric.Category,
							Type:     unit.Metric.Type,
						}, metricReadData)
						if err != nil {
							jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
								datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
							scope.Errorf("[%s] List metric from time %v failed: %s", jobID, metricStartT, err.Error())
							continue
						}
						for _, metricRawDatum := range metricReadDataRes.GetData().GetRawdata() {
							for _, metricRawDatumGrp := range metricRawDatum.GetGroups() {
								metricCols := metricRawDatumGrp.GetColumns()
								metricRows := metricRawDatumGrp.GetRows()
								if len(metricRows) == 0 {
									scope.Errorf("[%s] No metric found from time %v, skip drift evaluation.", jobID, metricStartT)
									continue
								}
								measurementDataSet := stats.NewMeasurementDataSetV2(modelPredictRows, modelIDPredictDataGrp.GetColumns(),
									metricRows, metricCols, unit, granularity)

								currentMeasure := viper.GetString("measurements.current")
								mapeVal, mapeErr := stats.MAPE(measurementDataSet, granularity)
								if mapeErr != nil {
									scope.Errorf("[%s] Calculate MAPE failed due to %s", jobID, mapeErr.Error())
								} else {
									metrics.SetMetricMAPE(jobID, mapeVal)
								}
								rmseVal, rmseErr := stats.RMSE(measurementDataSet, datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
								if rmseErr != nil {
									scope.Errorf("[%s] Calculate RMSE failed due to %s", jobID, rmseErr.Error())
								} else {
									metrics.SetMetricRMSE(jobID, rmseVal)
								}

								if strings.ToLower(strings.TrimSpace(currentMeasure)) == "mape" && mapeErr == nil {
									modelThreshold := viper.GetFloat64("measurements.mape.threshold")
									if mapeVal > modelThreshold {
										scope.Infof("[%s] MAPE of metric %v  %v > %v (threshold), drift is true and start sending model job",
											jobID, metricType, mapeVal, modelThreshold)
										err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
											metricType, granularity, queueSender, jobID)
										if err != nil {
											scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
										} else {
											metrics.AddMetricDrift(jobID, 1.0)
										}
									} else {
										scope.Infof("[%s] MAPE of metric %v  %v <= %v (threshold), drift is false", jobID, metricType, mapeVal, modelThreshold)
									}
								} else if strings.ToLower(strings.TrimSpace(currentMeasure)) == "rmse" && rmseErr == nil {
									modelThreshold := viper.GetFloat64("measurements.rmse.threshold")
									if rmseVal > modelThreshold {
										scope.Infof("[%s] RMSE of metric %v  %v > %v (threshold), drift is true and start sending model job",
											jobID, metricType, rmseVal, modelThreshold)
										err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
											metricType, granularity, queueSender, jobID)
										if err != nil {
											scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
										} else {
											metrics.AddMetricDrift(jobID, 1.0)
										}
									} else {
										scope.Infof("[%s] RMSE of metric %v  %v <= %v (threshold), drift is false",
											jobID, metricType, rmseVal, modelThreshold)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (dispatcher *modelJobSender) tryToJobSending(queueName string, unit *config.Unit,
	rawDataCols []string, rawDataVals []string, metricType datahub_common.MetricType,
	granularity int64, queueSender queue.QueueSender, jobID string) error {
	if dispatcher.modelMapper.IsModelingV2(jobID) && !dispatcher.modelMapper.IsModelTimeoutV2(jobID) {
		return fmt.Errorf("model job with id is processing, do not send duplicated")
	}
	err := queueSender.SendJob(queueName, unit, rawDataCols, rawDataVals,
		metricType, granularity)
	if err != nil {
		return err
	} else {
		dispatcher.modelMapper.AddModelInfoV2(jobID)
	}
	return nil
}

func (dispatcher *modelJobSender) isUnitWatchedByScaler(unit *config.Unit, rowValues []string, columns []string) (bool, error) {
	k8sNS, k8sName, err := GetUnitResourceK8SNSName(unit, rowValues, columns)
	if err != nil {
		return false, err
	}
	if k8sNS == nil || k8sName == nil {
		return true, nil
	}
	if *k8sNS == "" || *k8sName == "" {
		return false, nil
	}
	return true, nil
}
