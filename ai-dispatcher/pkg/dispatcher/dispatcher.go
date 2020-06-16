package dispatcher

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/consts"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/metrics"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/queue"
	utils "github.com/containers-ai/alameda/ai-dispatcher/pkg/utils"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_gpu "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/gpu"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const queueName = "predict"
const modelQueueName = "model"

var (
	modelHasVPA   = false
	predictHasVPA = false
)

var scope = log.RegisterScope("dispatcher", "dispatcher dispatch jobs", 0)

type Dispatcher struct {
	svcGranularities []string
	svcPredictUnits  []string
	datahubGrpcCn    *grpc.ClientConn
	queueConn        *amqp.Connection

	modelJobSender   *modelJobSender
	predictJobSender *predictJobSender
	cfg              *config.Config
}

func NewDispatcher(datahubGrpcCn *grpc.ClientConn, granularities []string,
	predictUnits []string, modelMapper *ModelMapper, metricExporter *metrics.Exporter,
	cfg *config.Config) *Dispatcher {
	modelJobSender := NewModelJobSender(datahubGrpcCn, modelMapper, metricExporter)
	predictJobSender := NewPredictJobSender(datahubGrpcCn)
	dispatcher := &Dispatcher{
		svcGranularities: granularities,
		svcPredictUnits:  predictUnits,
		datahubGrpcCn:    datahubGrpcCn,
		modelJobSender:   modelJobSender,
		predictJobSender: predictJobSender,
		cfg:              cfg,
	}
	dispatcher.validCfg()
	dispatcher.initExportedMetrics(cfg)
	return dispatcher
}

var wg sync.WaitGroup

func (dispatcher *Dispatcher) Start() {
	// generate len(dispatcher.svcGranularities) senders to publish job,
	// each sender use distinct channel which is not thread safe.
	// all jobs are published to the same queue.
	for _, granularity := range dispatcher.svcGranularities {
		predictionStep := viper.GetInt(fmt.Sprintf("granularities.%s.predictionSteps",
			granularity))
		if predictionStep == 0 {
			scope.Warnf("Prediction step of Granularity %v is not defined or set incorrect.",
				granularity)
			continue
		}
		wg.Add(1)
		go dispatcher.dispatch(granularity, int64(predictionStep),
			"predictionJobSendIntervalSec")
		wg.Add(1)
		go dispatcher.dispatch(granularity, int64(predictionStep),
			"modelJobSendIntervalSec")
	}
	wg.Wait()
}

func (dispatcher *Dispatcher) validCfg() {
	if len(dispatcher.svcGranularities) == 0 {
		scope.Errorf("no setting of granularities of service")
		os.Exit(1)
	}

	if len(dispatcher.svcPredictUnits) == 0 {
		scope.Errorf("no setting of predict units of service")
		os.Exit(1)
	}
}

func (dispatcher *Dispatcher) dispatch(granularity string, predictionStep int64,
	queueJobType string) {
	defer wg.Done()
	granularitySec := int64(viper.GetInt(
		fmt.Sprintf("granularities.%s.dataGranularitySec", granularity)))
	if granularitySec == 0 {
		scope.Warnf("Granularity %v is not defined or set incorrect.", granularitySec)
		return
	}
	queueJobSendIntervalSec := viper.GetInt(
		fmt.Sprintf("granularities.%s.%s", granularity, queueJobType))
	queueURL := viper.GetString("queue.url")
	queueConnRetryItvMS := viper.GetInt64("queue.retry.connectIntervalMs")
	if queueConnRetryItvMS == 0 {
		queueConnRetryItvMS = 3000
	}
	for {
		queueSender, queueConn := queue.NewRabbitMQSender(queueURL, queueConnRetryItvMS)
		// Node will send model/predict job with granularity 30s if modelHasVPA/predictHasVPA is true
		if granularitySec == 30 {
			modelHasVPA = false
			predictHasVPA = false
		}
		// New API section
		for _, unit := range dispatcher.cfg.GetUnits() {
			if queueJobType == "predictionJobSendIntervalSec" {
				err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.predict.directory"), granularitySec))
				if err != nil {
					scope.Error(err.Error())
				}
			} else if queueJobType == "modelJobSendIntervalSec" {
				err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.model.directory"), granularitySec))
				if err != nil {
					scope.Error(err.Error())
				}
			}

			if !unit.Enabled {
				continue
			}
			granularityFound := false
			for _, granu := range unit.Granularities {
				if granu == granularity {
					granularityFound = true
					break
				}
			}
			if !granularityFound {
				continue
			}

			unitScope := unit.Scope
			category := unit.Category
			unitType := unit.Type

			if queueJobType == "predictionJobSendIntervalSec" {
				scope.Infof(
					"Start dispatching prediction unit with (scope: %s, category: %s, type: %s) with granularity %v seconds and cycle %v seconds",
					unitScope, category, unitType, granularitySec, queueJobSendIntervalSec)
			} else if queueJobType == "modelJobSendIntervalSec" {
				scope.Infof(
					"Start dispatching model unit with (scope %s, category %s, type: %s) with granularity %v seconds and cycle %v seconds",
					unitScope, category, unitType, granularitySec, queueJobSendIntervalSec)
			}
			dispatcher.getAndPushJobsV2(queueSender, &unit, granularitySec,
				predictionStep, queueJobType)
		}

		if granularity != "1m" && granularity != "3m" {
			for _, pdUnit := range dispatcher.svcPredictUnits {
				if queueJobType == "predictionJobSendIntervalSec" {
					err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.predict.directory"), granularitySec))
					if err != nil {
						scope.Error(err.Error())
					}
				} else if queueJobType == "modelJobSendIntervalSec" {
					err := utils.TouchFile(fmt.Sprintf("%s/%v", viper.GetString("watchdog.model.directory"), granularitySec))
					if err != nil {
						scope.Error(err.Error())
					}
				}

				if dispatcher.skipJobSending(pdUnit, granularitySec) {
					continue
				}

				pdUnitType := viper.GetString(fmt.Sprintf("predictUnits.%s.type", pdUnit))

				if pdUnitType == "" {
					scope.Warnf("Unit %s is not defined or set incorrect.", pdUnit)
					continue
				}

				if queueJobType == "predictionJobSendIntervalSec" {
					scope.Infof(
						"Start dispatching prediction unit %s with granularity %v seconds and cycle %v seconds",
						pdUnitType, granularitySec, queueJobSendIntervalSec)
				} else if queueJobType == "modelJobSendIntervalSec" {
					scope.Infof(
						"Start dispatching model unit %s with granularity %v seconds and cycle %v seconds",
						pdUnitType, granularitySec, queueJobSendIntervalSec)
				}

				dispatcher.getAndPushJobs(queueSender, pdUnit, granularitySec,
					predictionStep, queueJobType)
			}
		}
		queueConn.Close()
		time.Sleep(time.Duration(queueJobSendIntervalSec) * time.Second)
	}
}

func (dispatcher *Dispatcher) getAndPushJobsV2(queueSender queue.QueueSender,
	pdUnit *config.Unit, granularity int64, predictionStep int64, queueJobType string) {
	datahubServiceClnt := datahub_v1alpha1.NewDatahubServiceClient(dispatcher.datahubGrpcCn)
	data, err := utils.ReadData(datahubServiceClnt, &datahub_schemas.SchemaMeta{
		Scope:    pdUnit.Scope,
		Category: pdUnit.Category,
		Type:     pdUnit.Type,
	}, []*datahub_data.ReadData{
		{
			Measurement: pdUnit.Measurement,
		},
	})

	if err != nil {
		scope.Errorf("List units with (scope %s, category %s, type: %s) failed: %s",
			pdUnit.Scope, pdUnit.Category, pdUnit.Type, err.Error())
		return
	}
	readData := []*datahub_data.ReadData{}
	for _, metricType := range pdUnit.MetricTypes {
		readData = append(readData, &datahub_data.ReadData{
			MetricType: metricType,
		})
	}

	if queueJobType == "predictionJobSendIntervalSec" {
		dispatcher.predictJobSender.SendPredictJobs(data.GetData().GetRawdata(), queueSender, pdUnit, granularity)
	}
	if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
		dispatcher.modelJobSender.SendModelJobs(data.GetData().GetRawdata(), queueSender, pdUnit, granularity)
	}
}

func (dispatcher *Dispatcher) getAndPushJobs(queueSender queue.QueueSender,
	pdUnit string, granularity int64, predictionStep int64, queueJobType string) {

	datahubServiceClnt := datahub_v1alpha1.NewDatahubServiceClient(dispatcher.datahubGrpcCn)

	if pdUnit == consts.UnitTypeNode {
		res, err := datahubServiceClnt.ListNodes(context.Background(),
			&datahub_resources.ListNodesRequest{})
		if err != nil {
			scope.Errorf(
				"List nodes for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}

		nodes := []*datahub_resources.Node{}
		if queueJobType == "predictionJobSendIntervalSec" {
			for _, no := range res.GetNodes() {
				if (granularity == 30 && !viper.GetBool("hourlyPredict")) && !predictHasVPA {
					continue
				}
				nodes = append(nodes, no)
			}
			scope.Infof(
				"Start sending %v node prediction jobs to queue with granularity %v seconds.",
				len(nodes), granularity)
			dispatcher.predictJobSender.SendNodePredictJobs(nodes, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			for _, no := range res.GetNodes() {
				if (granularity == 30 && !viper.GetBool("hourlyPredict")) && !modelHasVPA {
					continue
				}
				nodes = append(nodes, no)
			}
			scope.Infof(
				"Start sending %v node model jobs to queue with granularity %v seconds.",
				len(nodes), granularity)
			dispatcher.modelJobSender.SendNodeModelJobs(nodes, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof(
			"Sending %v node jobs to queue completely with granularity %v seconds.",
			len(nodes), granularity)

	} else if pdUnit == consts.UnitTypePod {
		res, err := datahubServiceClnt.ListPods(context.Background(),
			&datahub_resources.ListPodsRequest{
				ScalingTool: datahub_resources.ScalingTool_VPA,
			})
		if err != nil {
			scope.Errorf(
				"List pods for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}

		pods := []*datahub_resources.Pod{}
		hasVPA := false
		for _, pod := range res.GetPods() {
			if granularity == 30 && (!viper.GetBool("hourlyPredict") &&
				pod.GetAlamedaPodSpec().GetScalingTool() != datahub_resources.ScalingTool_VPA) {
				continue
			}
			if pod.GetAlamedaPodSpec().GetScalingTool() == datahub_resources.ScalingTool_VPA {
				hasVPA = true
			}
			pods = append(pods, pod)
		}

		if queueJobType == "predictionJobSendIntervalSec" {
			if hasVPA {
				predictHasVPA = true
			}
			scope.Infof(
				"Start sending %v pod prediction jobs to queue with granularity %v seconds.",
				len(pods), granularity)
			dispatcher.predictJobSender.SendPodPredictJobs(pods, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			if hasVPA {
				modelHasVPA = true
			}
			scope.Infof(
				"Start sending %v pod model jobs to queue with granularity %v seconds.",
				len(pods), granularity)
			dispatcher.modelJobSender.SendPodModelJobs(pods, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof(
			"Sending %v pod jobs to queue completely with granularity %v seconds.",
			len(pods), granularity)
	} else if pdUnit == consts.UnitTypeGPU {
		res, err := datahubServiceClnt.ListGpus(context.Background(),
			&datahub_gpu.ListGpusRequest{})
		if err != nil {
			scope.Errorf(
				"List gpus for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}
		gpus := res.GetGpus()
		if queueJobType == "predictionJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v gpu prediction jobs to queue with granularity %v seconds.",
				len(gpus), granularity)
			dispatcher.predictJobSender.SendGPUPredictJobs(gpus, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v gpu model jobs to queue with granularity %v seconds.",
				len(gpus), granularity)
			dispatcher.modelJobSender.SendGPUModelJobs(gpus, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof("Sending %v gpu jobs to queue completely with granularity %v seconds.",
			len(gpus), granularity)
	} else if pdUnit == consts.UnitTypeApplication {
		res, err := datahubServiceClnt.ListApplications(context.Background(),
			&datahub_resources.ListApplicationsRequest{})
		if err != nil {
			scope.Errorf(
				"List applications for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}
		applications := []*datahub_resources.Application{}
		for _, app := range res.GetApplications() {
			if granularity == 30 && (!viper.GetBool("hourlyPredict") &&
				app.GetAlamedaApplicationSpec().GetScalingTool() !=
					datahub_resources.ScalingTool_VPA) {
				continue
			}
			applications = append(applications, app)
		}
		if queueJobType == "predictionJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v application prediction jobs to queue with granularity %v seconds.",
				len(applications), granularity)
			dispatcher.predictJobSender.SendApplicationPredictJobs(applications, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v application model jobs to queue with granularity %v seconds.",
				len(applications), granularity)
			dispatcher.modelJobSender.SendApplicationModelJobs(applications, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof("Sending %v application jobs to queue completely with granularity %v seconds.",
			len(applications), granularity)
	} else if pdUnit == consts.UnitTypeNamespace {
		res, err := datahubServiceClnt.ListNamespaces(context.Background(),
			&datahub_resources.ListNamespacesRequest{})
		if err != nil {
			scope.Errorf(
				"List namespaces for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}

		namespaces := res.GetNamespaces()
		if queueJobType == "predictionJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v namespace prediction jobs to queue with granularity %v seconds.",
				len(namespaces), granularity)
			dispatcher.predictJobSender.SendNamespacePredictJobs(namespaces, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v namespace model jobs to queue with granularity %v seconds.",
				len(namespaces), granularity)
			dispatcher.modelJobSender.SendNamespaceModelJobs(namespaces, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof(
			"Sending %v namespace jobs to queue completely with granularity %v seconds.",
			len(namespaces), granularity)
	} else if pdUnit == consts.UnitTypeCluster {
		res, err := datahubServiceClnt.ListClusters(context.Background(),
			&datahub_resources.ListClustersRequest{})
		if err != nil {
			scope.Errorf(
				"List clusters for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}

		clusters := res.GetClusters()
		if queueJobType == "predictionJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v cluster prediction jobs to queue with granularity %v seconds.",
				len(clusters), granularity)
			dispatcher.predictJobSender.SendClusterPredictJobs(clusters, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v cluster model jobs to queue with granularity %v seconds.",
				len(clusters), granularity)
			dispatcher.modelJobSender.SendClusterModelJobs(clusters, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof(
			"Sending %v cluster jobs to queue completely with granularity %v seconds.",
			len(clusters), granularity)
	} else if pdUnit == consts.UnitTypeController {
		res, err := datahubServiceClnt.ListControllers(context.Background(),
			&datahub_resources.ListControllersRequest{})
		if err != nil {
			scope.Errorf(
				"List controllers for model/predict job failed with granularity %v seconds. %s",
				granularity, err.Error())
			return
		}
		controllers := []*datahub_resources.Controller{}
		for _, ctrl := range res.GetControllers() {
			if granularity == 30 && (!viper.GetBool("hourlyPredict") &&
				ctrl.GetAlamedaControllerSpec().GetScalingTool() !=
					datahub_resources.ScalingTool_VPA) {
				continue
			}
			controllers = append(controllers, ctrl)
		}
		if queueJobType == "predictionJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v controller prediction jobs to queue with granularity %v seconds.",
				len(controllers), granularity)
			dispatcher.predictJobSender.SendControllerPredictJobs(controllers, queueSender, pdUnit, granularity)
		}
		if viper.GetBool("model.enabled") && queueJobType == "modelJobSendIntervalSec" {
			scope.Infof(
				"Start sending %v controller model jobs to queue with granularity %v seconds.",
				len(controllers), granularity)
			dispatcher.modelJobSender.SendControllerModelJobs(controllers, queueSender, pdUnit, granularity,
				predictionStep)
		}
		scope.Infof("Sending %v controller jobs to queue completely with granularity %v seconds.",
			len(controllers), granularity)
	}
}

func (dispatcher *Dispatcher) skipJobSending(pdUnit string, granularitySec int64) bool {
	if pdUnit == consts.UnitTypeGPU && granularitySec != 3600 {
		return true
	}

	return (pdUnit == consts.UnitTypeCluster || pdUnit == consts.UnitTypeNamespace) &&
		(granularitySec == 30 && !viper.GetBool("hourlyPredict"))
}

func (dispatcher *Dispatcher) initExportedMetrics(cfg *config.Config) {
	for _, unit := range dispatcher.cfg.GetUnits() {
		metrics.InitMetric(fmt.Sprintf("%s", unit.Scope), unit.Category, unit.Type, unit.IDKeys)
	}
}
