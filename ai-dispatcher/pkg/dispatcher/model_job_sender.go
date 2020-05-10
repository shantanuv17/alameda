package dispatcher

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/metrics"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/queue"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/stats"
	utils "github.com/containers-ai/alameda/ai-dispatcher/pkg/utils"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_gpu "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/gpu"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type modelJobSender struct {
	datahubServiceClnt datahub_v1alpha1.DatahubServiceClient
	datahubGrpcCn      *grpc.ClientConn
	modelMapper        *ModelMapper
	metricExporter     *metrics.Exporter

	podModelJobSender         *podModelJobSender
	nodeModelJobSender        *nodeModelJobSender
	gpuModelJobSender         *gpuModelJobSender
	applicationModelJobSender *applicationModelJobSender
	namespaceModelJobSender   *namespaceModelJobSender
	clusterModelJobSender     *clusterModelJobSender
	controllerModelJobSender  *controllerModelJobSender
}

func NewModelJobSender(datahubGrpcCn *grpc.ClientConn, modelMapper *ModelMapper,
	metricExporter *metrics.Exporter) *modelJobSender {

	return &modelJobSender{
		datahubGrpcCn:      datahubGrpcCn,
		modelMapper:        modelMapper,
		metricExporter:     metricExporter,
		datahubServiceClnt: datahub_v1alpha1.NewDatahubServiceClient(datahubGrpcCn),

		podModelJobSender: NewPodModelJobSender(datahubGrpcCn, modelMapper,
			metricExporter),
		nodeModelJobSender: NewNodeModelJobSender(datahubGrpcCn, modelMapper,
			metricExporter),
		gpuModelJobSender: NewGPUModelJobSender(datahubGrpcCn, modelMapper,
			metricExporter),
		applicationModelJobSender: NewApplicationModelJobSender(datahubGrpcCn, modelMapper,
			metricExporter),
		namespaceModelJobSender: NewNamespaceModelJobSender(datahubGrpcCn, modelMapper,
			metricExporter),
		clusterModelJobSender: NewClusterModelJobSender(datahubGrpcCn, modelMapper,
			metricExporter),
		controllerModelJobSender: NewControllerModelJobSender(datahubGrpcCn, modelMapper,
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
								&datahub_common.Condition{
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
				readDataRes, err := utils.ReadData(dispatcher.datahubServiceClnt, &datahub_schemas.SchemaMeta{
					Scope:    unit.Prediction.Scope,
					Category: unit.Prediction.Category,
					Type:     unit.Prediction.Type,
				}, readData)
				if err != nil {
					jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
						datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
					scope.Errorf("[%s] List last prediction point failed: %s", jobID, err.Error())
					continue
				}

				if len(readDataRes.GetData().GetRawdata()) == 0 {
					for _, mt := range unit.MetricTypes {
						scope.Infof("[%s] No prediction found, send model jobs")
						jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
							mt, granularity)
						err := dispatcher.tryToJobSending(modelQueueName, unit, rawDatumColumns, row.GetValues(),
							mt, granularity, queueSender, jobID)
						if err != nil {
							scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
						}
					}
					return
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
							scope.Infof("[%s] Use model ID %s to query prediction series to measure dirft", jobID, modelID)
							dispatcher.driftEval(modelID, lastPredictPointRawData.GetMetricType(), rawData, queueSender, unit, granularity)
						}
					}
				}
			}
		}
	}
}

func (dispatcher *modelJobSender) driftEval(modelID string,
	metricType datahub_common.MetricType, rawData []*datahub_data.Rawdata,
	queueSender queue.QueueSender, unit *config.Unit, granularity int64) {
	for _, rawDatum := range rawData {
		for _, grp := range rawDatum.GetGroups() {
			rawDatumColumns := grp.GetColumns()
			for _, row := range grp.GetRows() {
				readData := []*datahub_data.ReadData{&datahub_data.ReadData{
					MetricType:       metricType,
					ResourceBoundary: datahub_common.ResourceBoundary_RESOURCE_RAW,
					QueryCondition: &datahub_common.QueryCondition{
						Order: datahub_common.QueryCondition_DESC,
						WhereCondition: []*datahub_common.Condition{
							&datahub_common.Condition{
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

				readDataRes, err := utils.ReadData(dispatcher.datahubServiceClnt, &datahub_schemas.SchemaMeta{
					Scope:    unit.Prediction.Scope,
					Category: unit.Prediction.Category,
					Type:     unit.Prediction.Type,
				}, readData)
				if err != nil {
					jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
						datahub_common.MetricType_METRICS_TYPE_UNDEFINED, granularity)
					scope.Errorf("[%s] List prediction with model id %s failed: %s", jobID, modelID, err.Error())
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

						metricReadDataRes, err := dispatcher.formatMetrics(
							metricStartT, jobID, row, rawDatumColumns,
							metricType, unit, granularity)
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

func (dispatcher *modelJobSender) formatMetrics(metricStartT int64, jobID string,
	row *datahub_common.Row, rawDatumColumns []string, metricType datahub_common.MetricType,
	unit *config.Unit, granularity int64) (*datahub_data.ReadDataResponse, error) {
	if unit.Category == "cluster_autoscaler" && unit.Type == "machinegroup" &&
		unit.UnitParameters != nil && unit.UnitValueKeys != nil {
		return dispatcher.formatMachineGroupMetrics(metricStartT, jobID,
			row, rawDatumColumns, metricType,
			unit, granularity)
	}
	whereVals := []string{}
	whereOps := []string{}
	whereTypes := []datahub_common.DataType{}
	for _, idK := range unit.IDKeys {
		whereVal, err := utils.GetRowValue(row.GetValues(), rawDatumColumns, idK)
		if err != nil {
			return nil, fmt.Errorf("[%s] Cannot query metric for drift evaluation due to %s", jobID, err.Error())
		}
		whereVals = append(whereVals, whereVal)
		whereOps = append(whereOps, "=")
		whereTypes = append(whereTypes, datahub_common.DataType_DATATYPE_STRING)
	}

	metricReadData := []*datahub_data.ReadData{&datahub_data.ReadData{
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
				&datahub_common.Condition{
					Keys:      unit.IDKeys,
					Values:    whereVals,
					Operators: whereOps,
					Types:     whereTypes,
				},
			},
		},
	}}

	return utils.ReadData(dispatcher.datahubServiceClnt, &datahub_schemas.SchemaMeta{
		Scope:    unit.Metric.Scope,
		Category: unit.Metric.Category,
		Type:     unit.Metric.Type,
	}, metricReadData)
}

func (dispatcher *modelJobSender) formatMachineGroupMetrics(metricStartT int64, jobID string,
	row *datahub_common.Row, rawDatumColumns []string, metricType datahub_common.MetricType,
	unit *config.Unit, granularity int64) (*datahub_data.ReadDataResponse, error) {

	nodes := []string{}
	clusterName, err := utils.GetRowValue(row.GetValues(),
		rawDatumColumns, unit.UnitValueKeys.ClusterName)
	if err != nil {
		return nil, err
	}
	name, err := utils.GetRowValue(row.GetValues(),
		rawDatumColumns, unit.UnitValueKeys.Name)
	if err != nil {
		return nil, err
	}
	namespace, err := utils.GetRowValue(row.GetValues(),
		rawDatumColumns, unit.UnitValueKeys.Namespace)
	if err != nil {
		return nil, err
	}

	machinesetReadData := []*datahub_data.ReadData{&datahub_data.ReadData{
		QueryCondition: &datahub_common.QueryCondition{
			Order: datahub_common.QueryCondition_DESC,
			WhereCondition: []*datahub_common.Condition{
				&datahub_common.Condition{
					Keys:      unit.UnitParameters.MachineSetQueryKeys,
					Values:    []string{clusterName, namespace, name},
					Operators: []string{"=", "=", "="},
					Types: []datahub_common.DataType{
						datahub_common.DataType_DATATYPE_STRING,
						datahub_common.DataType_DATATYPE_STRING,
						datahub_common.DataType_DATATYPE_STRING,
					},
				},
			},
		},
	}}
	machineSetReadDataRes, err := utils.ReadData(dispatcher.datahubServiceClnt, &datahub_schemas.SchemaMeta{
		Scope:    unit.Scope,
		Category: unit.Category,
		Type:     unit.UnitParameters.MachineSetType,
	}, machinesetReadData)
	if err != nil {
		return nil, err
	}
	for _, machineSetReadDatum := range machineSetReadDataRes.GetData().GetRawdata() {
		for _, machineSetReadDatumGrp := range machineSetReadDatum.GetGroups() {
			machineSetColumns := machineSetReadDatumGrp.GetColumns()
			for _, machinesetRow := range machineSetReadDatumGrp.GetRows() {
				machineSetNamespace, err := utils.GetRowValue(machinesetRow.GetValues(),
					machineSetColumns, unit.UnitValueKeys.Namespace)
				if err != nil {
					return nil, err
				}
				machineSetName, err := utils.GetRowValue(machinesetRow.GetValues(),
					machineSetColumns, unit.UnitValueKeys.Name)
				if err != nil {
					return nil, err
				}
				nodeReadData := []*datahub_data.ReadData{&datahub_data.ReadData{
					QueryCondition: &datahub_common.QueryCondition{
						Order: datahub_common.QueryCondition_DESC,
						WhereCondition: []*datahub_common.Condition{
							&datahub_common.Condition{
								Keys:      unit.UnitParameters.NodeQueryKeys,
								Values:    []string{clusterName, machineSetNamespace, machineSetName},
								Operators: []string{"=", "=", "="},
								Types: []datahub_common.DataType{
									datahub_common.DataType_DATATYPE_STRING,
									datahub_common.DataType_DATATYPE_STRING,
									datahub_common.DataType_DATATYPE_STRING,
								},
							},
						},
					},
				}}
				nodeReadDataRes, err := utils.ReadData(dispatcher.datahubServiceClnt, &datahub_schemas.SchemaMeta{
					Scope:    unit.Scope,
					Category: unit.UnitParameters.ClusterStatusCategory,
					Type:     unit.UnitParameters.NodeType,
				}, nodeReadData)
				if err != nil {
					return nil, err
				}
				for _, nodeReadDatum := range nodeReadDataRes.GetData().GetRawdata() {
					for _, nodeReadDatumGrp := range nodeReadDatum.GetGroups() {
						nodeColumns := nodeReadDatumGrp.GetColumns()
						for _, nodeRow := range nodeReadDatumGrp.GetRows() {
							nodeName, err := utils.GetRowValue(nodeRow.GetValues(), nodeColumns, unit.UnitValueKeys.Name)
							if err != nil {
								return nil, err
							}
							nodes = append(nodes, nodeName)
						}
					}
				}
			}
		}
	}
	metricReadData := []*datahub_data.ReadData{}
	for _, node := range nodes {
		metricReadData = append(metricReadData, &datahub_data.ReadData{
			MetricType: metricType,
			QueryCondition: &datahub_common.QueryCondition{
				Order:   datahub_common.QueryCondition_DESC,
				Selects: []string{unit.Metric.MetricValueKeys.Value},
				Groups:  []string{unit.UnitValueKeys.ClusterName, unit.UnitValueKeys.Name},
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
					&datahub_common.Condition{
						Keys:      []string{unit.UnitValueKeys.ClusterName, unit.UnitValueKeys.Name},
						Values:    []string{clusterName, node},
						Operators: []string{"=", "="},
						Types: []datahub_common.DataType{
							datahub_common.DataType_DATATYPE_STRING,
							datahub_common.DataType_DATATYPE_STRING,
						},
					},
				},
			},
		})
	}

	metricReadDataRes, err := utils.ReadData(dispatcher.datahubServiceClnt, &datahub_schemas.SchemaMeta{
		Scope:    unit.Metric.Scope,
		Category: unit.UnitParameters.ClusterStatusCategory,
		Type:     unit.UnitParameters.NodeType,
	}, metricReadData)

	if err != nil {
		return nil, err
	}
	mainRows := []*datahub_common.Row{}
	if len(metricReadDataRes.GetData().GetRawdata()) == 1 {
		return metricReadDataRes, err
	} else if len(metricReadDataRes.GetData().GetRawdata()) == 0 {
		return nil, nil
	}
	for _, metricReadDatum := range metricReadDataRes.GetData().GetRawdata() {
		for _, metricReadDatumGrp := range metricReadDatum.GetGroups() {
			for _, metricReadDatumRow := range metricReadDatumGrp.GetRows() {
				metricValueStr, err := utils.GetRowValue(metricReadDatumRow.GetValues(),
					metricReadDatumGrp.GetColumns(), unit.Metric.MetricValueKeys.Value)
				if err != nil {
					return nil, err
				}
				timeFound := false
				for mainRowIdx, mainRow := range mainRows {
					if mainRow.GetTime().GetSeconds() == metricReadDatumRow.GetTime().GetSeconds() {
						timeFound = true
						metricValue, err := strconv.ParseFloat(metricValueStr, 64)
						if err != nil {
							return nil, err
						}
						existingMetricValue, err := strconv.ParseFloat(mainRows[mainRowIdx].Values[0], 64)
						if err != nil {
							return nil, err
						}
						mainRows[mainRowIdx].Values[0] = fmt.Sprintf("%f", existingMetricValue+metricValue)
						break
					}
				}
				if !timeFound {
					mainRows = append(mainRows, &datahub_common.Row{
						Time:   metricReadDatumRow.GetTime(),
						Values: []string{metricValueStr},
					})
				}
			}
		}
	}

	sort.SliceStable(mainRows, func(i, j int) bool {
		if metricReadData[0].QueryCondition.Order == datahub_common.QueryCondition_ASC {
			return mainRows[i].GetTime().GetSeconds() < mainRows[j].GetTime().GetSeconds()
		}
		return mainRows[i].GetTime().GetSeconds() > mainRows[j].GetTime().GetSeconds()
	})

	machineScalerReadDataRes := &datahub_data.ReadDataResponse{
		Data: &datahub_data.Data{
			SchemaMeta: &datahub_schemas.SchemaMeta{},
			Rawdata: []*datahub_data.Rawdata{
				&datahub_data.Rawdata{
					MetricType: metricType,
					Groups: []*datahub_common.Group{
						&datahub_common.Group{
							Columns: []string{unit.Metric.MetricValueKeys.Value},
							Rows:    mainRows,
						},
					},
				},
			},
		},
	}

	return machineScalerReadDataRes, nil
}
