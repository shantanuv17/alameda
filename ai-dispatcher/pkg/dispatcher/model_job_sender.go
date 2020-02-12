package dispatcher

import (
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/metrics"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/queue"
	utils "github.com/containers-ai/alameda/ai-dispatcher/pkg/utils"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_gpu "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/gpu"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/golang/protobuf/ptypes/duration"
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
		for _, grp := range rawDatum.GetGroups() {
			rawDatumColumns := grp.GetColumns()
			for _, row := range grp.GetRows() {
				readData := []*datahub_data.ReadData{}
				for _, metricType := range unit.MetricTypes {
					readData = append(readData, &datahub_data.ReadData{
						MetricType:       metricType,
						ResourceBoundary: datahub_common.ResourceBoundary_RESOURCE_RAW,
						QueryCondition: &datahub_common.QueryCondition{
							Order: datahub_common.QueryCondition_DESC,
							Limit: 1,
							TimeRange: &datahub_common.TimeRange{
								Step: &duration.Duration{
									Seconds: granularity,
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

				for _, lastPredictPointRawData := range readDataRes.GetData().GetRawdata() {
					for _, lastPredictPointRawDataGrp := range lastPredictPointRawData.GetGroups() {
						jobID, _ := utils.GetJobID(unit, row.GetValues(), rawDatumColumns,
							lastPredictPointRawData.GetMetricType(), granularity)
						if len(lastPredictPointRawDataGrp.GetRows()) == 0 {
							scope.Infof("[%s] No prediction found, send model job.", jobID)
							err := queueSender.SendJob(modelQueueName, unit, rawDatumColumns, row.GetValues(),
								lastPredictPointRawData.GetMetricType(), granularity)
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
								err := queueSender.SendJob(modelQueueName, unit, rawDatumColumns, row.GetValues(),
									lastPredictPointRawData.GetMetricType(), granularity)
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
							scope.Debugf("[%s] Use model ID %s to query prediction series to measure dirft", jobID, modelID)
							err = queueSender.SendJob(modelQueueName, unit, rawDatumColumns, row.GetValues(),
								lastPredictPointRawData.GetMetricType(), granularity)
							if err != nil {
								scope.Errorf("[%s] Send model job failed due to %s.", jobID, err.Error())
							}
						}
					}
				}
			}
		}
	}
}
