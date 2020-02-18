/*
Copyright 2019 The Alameda Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	kafkaclient "github.com/containers-ai/alameda/internal/pkg/message-queue/kafka/client"
	"github.com/containers-ai/alameda/operator"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/autoscaling/v1alpha1"
	"github.com/containers-ai/alameda/operator/controllers"
	datahubclient "github.com/containers-ai/alameda/operator/datahub/client"
	datahub_client_application "github.com/containers-ai/alameda/operator/datahub/client/application"
	datahub_client_controller "github.com/containers-ai/alameda/operator/datahub/client/controller"
	datahub_client_kafka "github.com/containers-ai/alameda/operator/datahub/client/kafka"
	datahub_client_namespace "github.com/containers-ai/alameda/operator/datahub/client/namespace"
	datahub_client_node "github.com/containers-ai/alameda/operator/datahub/client/node"
	datahub_client_pod "github.com/containers-ai/alameda/operator/datahub/client/pod"
	internaldatahubschema "github.com/containers-ai/alameda/operator/datahub/schema"
	"github.com/containers-ai/alameda/operator/pkg/probe"
	"github.com/containers-ai/alameda/operator/pkg/utils"
	"github.com/containers-ai/alameda/pkg/provider"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	datahubv1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahubschemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"

	osappsapi "github.com/openshift/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	JSONIndent = "  "

	envVarPrefix = "ALAMEDA_OPERATOR"

	defaultRotationMaxSizeMegabytes = 100
	defaultRotationMaxBackups       = 7
	defaultLogRotateOutputFile      = "/var/log/alameda/alameda-operator.log"
)

var (
	// VERSION is sofeware version
	VERSION string
	// BUILD_TIME is build time
	BUILD_TIME string
	// GO_VERSION is go version
	GO_VERSION string

	// Variables for flags
	showVer              bool
	operatorConfigFile   string
	crdLocation          string
	readinessProbeFlag   bool
	livenessProbeFlag    bool
	metricsAddr          string
	enableLeaderElection bool

	// Global variables
	syncPriod                          = time.Duration(1 * time.Minute)
	hasOpenShiftAPIAppsv1              bool
	operatorConf                       operator.Config
	scope                              *logUtil.Scope
	alamedaScalerKafkaControllerLogger *logUtil.Scope
	datahubClientLogger                *logUtil.Scope

	clusterUID     string
	datahubSchemas = map[string]datahubschemas.Schema{
		"kafkaTopic":         datahubschemas.Schema{},
		"kafkaConsumerGroup": datahubschemas.Schema{},
	}

	// Third party clients
	k8sClient        client.Client
	datahubConn      *grpc.ClientConn
	datahubClient    datahubv1alpha1.DatahubServiceClient
	kafkaClient      kafka.Client
	prometheusClient prometheus.Prometheus

	// Resource repositories
	datahubKafkaRepo datahub_client_kafka.KafkaRepository
)

func init() {
	flag.BoolVar(&showVer, "version", false, "show version")
	flag.BoolVar(&readinessProbeFlag, "readiness-probe", false, "probe for readiness")
	flag.BoolVar(&livenessProbeFlag, "liveness-probe", false, "probe for liveness")
	flag.StringVar(&operatorConfigFile, "config", "/etc/alameda/operator/operator.toml",
		"File path to operator coniguration")
	flag.StringVar(&crdLocation, "crd-location", "/etc/alameda/operator/crds", "CRD location")
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")

	scope = logUtil.RegisterScope("manager", "operator entry point", 0)
	alamedaScalerKafkaControllerLogger = logUtil.RegisterScope("alameda_scaler_kafka_controller", "AlamedaScaler Kafka Controller", 0)
	datahubClientLogger = logUtil.RegisterScope("datahub_client", "AlamedaScaler Kafka Controller", 0)

	ok, err := utils.ServerHasOpenshiftAPIAppsV1()
	if err != nil {
		panic(errors.Wrap(err, "check if cluster has openshift api appsv1 failed"))
	}
	hasOpenShiftAPIAppsv1 = ok
}

func initLogger() error {

	opt := logUtil.DefaultOptions()
	opt.RotationMaxSize = defaultRotationMaxSizeMegabytes
	logFilePath := viper.GetString("log.filePath")
	if logFilePath == "" {
		logFilePath = defaultLogRotateOutputFile
	}
	opt.RotationMaxBackups = defaultRotationMaxBackups
	opt.RotateOutputPath = logFilePath
	if err := logUtil.Configure(opt); err != nil {
		return errors.Wrap(err, "configure log util failed")
	}

	scope.Infof("Log output level is %s.", operatorConf.Log.OutputLevel)
	scope.Infof("Log stacktrace level is %s.", operatorConf.Log.StackTraceLevel)
	for _, scope := range logUtil.Scopes() {
		scope.SetLogCallers(operatorConf.Log.SetLogCallers == true)
		if outputLvl, ok := logUtil.StringToLevel(operatorConf.Log.OutputLevel); ok {
			scope.SetOutputLevel(outputLvl)
		}
		if stacktraceLevel, ok :=
			logUtil.StringToLevel(operatorConf.Log.StackTraceLevel); ok {
			scope.SetStackTraceLevel(stacktraceLevel)
		}
	}

	return nil
}

func initServerConfig(mgr *manager.Manager) error {

	operatorConf = operator.NewConfigWithoutMgr()
	if mgr != nil {
		operatorConf = operator.NewConfig(*mgr)
	}

	viper.SetEnvPrefix(envVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// TODO: This config need default value. And it should check the file exists befor SetConfigFile.
	viper.SetConfigFile(operatorConfigFile)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read configuration failed")
	}
	if err := viper.Unmarshal(&operatorConf); err != nil {
		return errors.Wrap(err, "unmarshal config failed")
	}

	if operatorConfBin, err :=
		json.MarshalIndent(operatorConf, "", JSONIndent); err == nil {
		scope.Infof(fmt.Sprintf("Operator configuration: %s",
			string(operatorConfBin)))
	}
	return nil
}

func initThirdPartyClient() error {
	cli, err := client.New(ctrl.GetConfigOrDie(), client.Options{})
	if err != nil {
		return errors.Wrap(err, "new Kubernetes client failed")
	}
	k8sClient = cli

	datahubConn, err = grpc.Dial(operatorConf.Datahub.Address,
		grpc.WithBlock(),
		grpc.WithTimeout(30*time.Second),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(3)))),
	)
	if err != nil {
		return errors.Wrap(err, "new connection to datahub failed")
	}
	datahubClient = datahubv1alpha1.NewDatahubServiceClient(datahubConn)

	if cli, err := kafkaclient.NewClient(*operatorConf.Kafka); err != nil {
		return errors.Wrap(err, "new Kafka client failed")
	} else {
		kafkaClient = cli
	}

	if cli, err := prometheus.NewClient(&operatorConf.Prometheus.Config); err != nil {
		return errors.Wrap(err, "new Prometheus client failed")
	} else {
		prometheusClient = *cli
	}

	return nil
}

func initClusterUID() error {
	uid, err := k8sutils.GetClusterUID(k8sClient)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	} else if uid == "" {
		return errors.New("get empty cluster uid")
	}
	clusterUID = uid
	return nil
}

func initDatahubSchemas(ctx context.Context) error {
	// Get Schemas
	kafkaTopicSchema, err := internaldatahubschema.GetKafkaTopicSchema()
	if err != nil {
		return errors.Wrap(err, "get kafka topic schema failed")
	}
	datahubSchemas["kafkaTopic"] = kafkaTopicSchema
	kafkaConsumerGroupSchema, err := internaldatahubschema.GetKafkaConsumerGroupSchema()
	if err != nil {
		return errors.Wrap(err, "get kafka consumergroup schema failed")
	}
	datahubSchemas["kafkaConsumerGroup"] = kafkaConsumerGroupSchema

	// // Create schemas to Datahub
	// req := datahubschemas.CreateSchemasRequest{
	// 	Schemas: []*datahubschemas.Schema{&kafkaTopicSchema, &kafkaConsumerGroupSchema},
	// }
	// resp, err := datahubClient.CreateSchemas(ctx, &req)
	// if err != nil {
	// 	return errors.Wrap(err, "create schemas failed")
	// } else if resp == nil {
	// 	return errors.New("create schemas failed: receive nil status")
	// } else if resp.Code != int32(code.Code_OK) {
	// 	return errors.Errorf("create schemas failed: status: %d, message: %s", resp.Code, resp.Message)
	// }

	// List schemas from Datahub
	listSchemaReq := datahubschemas.ListSchemasRequest{}
	listSchemaResp, err := datahubClient.ListSchemas(ctx, &listSchemaReq)
	if err != nil {
		return errors.Wrap(err, "list schemas failed")
	} else if listSchemaResp == nil {
		return errors.New("list schemas failed: receive nil response")
	} else if ok, err := datahubclient.IsResponseStatusOK(listSchemaResp.Status); !ok || err != nil {
		return errors.Wrap(err, "list schemas failed")
	}

	return nil
}

func initDatahubResourceRepsitories() {
	datahubKafkaRepo = datahub_client_kafka.NewKafkaRepository(datahubClient, datahubClientLogger)
}

func setupManager() (manager.Manager, error) {
	return ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		Port:               9443,
		SyncPeriod:         &syncPriod,
	})
}

func addNecessaryAPIToScheme(scheme *runtime.Scheme) error {
	if err := autoscalingv1alpha1.AddToScheme(scheme); err != nil {
		return err
	}
	if hasOpenShiftAPIAppsv1 {
		if err := osappsapi.AddToScheme(scheme); err != nil {
			return err
		}
	}
	return nil
}

func addControllersToManager(mgr manager.Manager) error {
	datahubControllerRepo := datahub_client_controller.NewControllerRepository(datahubConn, clusterUID)
	datahubPodRepo := datahub_client_pod.NewPodRepository(datahubConn, clusterUID)
	datahubNamespaceRepo := datahub_client_namespace.NewNamespaceRepository(datahubConn, clusterUID)

	var err error

	if err = (&controllers.AlamedaScalerReconciler{
		Client:                 mgr.GetClient(),
		Scheme:                 mgr.GetScheme(),
		ClusterUID:             clusterUID,
		DatahubApplicationRepo: datahub_client_application.NewApplicationRepository(datahubConn, clusterUID),
		DatahubControllerRepo:  datahubControllerRepo,
		DatahubNamespaceRepo:   datahubNamespaceRepo,
		DatahubPodRepo:         datahubPodRepo,
		ReconcileTimeout:       3 * time.Second,
		ForceReconcileInterval: 1 * time.Minute,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if err = (&controllers.AlamedaRecommendationReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		ClusterUID:    clusterUID,
		DatahubClient: datahubv1alpha1.NewDatahubServiceClient(datahubConn),
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if err = (&controllers.DeploymentReconciler{
		Client:     mgr.GetClient(),
		Scheme:     mgr.GetScheme(),
		ClusterUID: clusterUID,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if hasOpenShiftAPIAppsv1 {
		if err = (&controllers.DeploymentConfigReconciler{
			Client:     mgr.GetClient(),
			Scheme:     mgr.GetScheme(),
			ClusterUID: clusterUID,
		}).SetupWithManager(mgr); err != nil {
			return err
		}
	}

	if err = (&controllers.NamespaceReconciler{
		Client:               mgr.GetClient(),
		Scheme:               mgr.GetScheme(),
		ClusterUID:           clusterUID,
		DatahubNamespaceRepo: datahubNamespaceRepo,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	cloudprovider := ""
	if provider.OnGCE() {
		cloudprovider = provider.GCP
	} else if provider.OnEC2() {
		cloudprovider = provider.AWS
	}
	regionName := ""
	switch cloudprovider {
	case provider.AWS:
		regionName = provider.AWSRegionMap[provider.GetEC2Region()]
	}
	if err = (&controllers.NodeReconciler{
		Client:          mgr.GetClient(),
		Scheme:          mgr.GetScheme(),
		ClusterUID:      clusterUID,
		Cloudprovider:   cloudprovider,
		RegionName:      regionName,
		DatahubNodeRepo: *datahub_client_node.NewNodeRepository(datahubConn, clusterUID),
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if err = (&controllers.StatefulSetReconciler{
		Client:     mgr.GetClient(),
		Scheme:     mgr.GetScheme(),
		ClusterUID: clusterUID,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if err = (&controllers.AlamedaScalerKafkaReconciler{
		ClusterUID:            clusterUID,
		HasOpenShiftAPIAppsv1: hasOpenShiftAPIAppsv1,

		K8SClient: mgr.GetClient(),
		Scheme:    mgr.GetScheme(),

		KafkaRepository:                                 datahubKafkaRepo,
		DatahubApplicationKafkaTopicSchema:              datahubSchemas["kafkaTopic"],
		DatahubApplicationKafkaTopicMeasurement:         *datahubSchemas["kafkaTopic"].Measurements[0],
		DatahubApplicationKafkaConsumerGroupSchema:      datahubSchemas["kafkaConsumerGroup"],
		DatahubApplicationKafkaConsumerGroupMeasurement: *datahubSchemas["kafkaConsumerGroup"].Measurements[0],

		KafkaClient:      kafkaClient,
		PrometheusClient: prometheusClient,

		ReconcileTimeout: 3 * time.Second,

		Logger: alamedaScalerKafkaControllerLogger,

		NeededMetrics: operatorConf.Prometheus.RequiredMetrics,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	printSoftwareInfo()
	if showVer {
		return
	}

	if readinessProbeFlag && livenessProbeFlag {
		scope.Error("Cannot run readiness probe and liveness probe at the same time")
		return
	} else if readinessProbeFlag {
		initServerConfig(nil)
		opWHSrvPort := viper.GetInt32("k8sWebhookServer.port")
		readinessProbe(&probe.ReadinessProbeConfig{
			WHSrvPort:   opWHSrvPort,
			DatahubAddr: operatorConf.Datahub.Address,
		})
		return
	} else if livenessProbeFlag {
		initServerConfig(nil)
		opWHSrvName := viper.GetString("k8sWebhookServer.service.name")
		opWHSrvNamespace := viper.GetString("k8sWebhookServer.service.namespace")
		opWHSrvPort := viper.GetInt32("k8sWebhookServer.service.port")
		livenessProbe(&probe.LivenessProbeConfig{
			ValidationSvc: &probe.ValidationSvc{
				SvcName: opWHSrvName,
				SvcNS:   opWHSrvNamespace,
				SvcPort: opWHSrvPort,
			},
		})
		return
	}

	mgr, err := setupManager()
	if err != nil {
		panic(errors.Wrap(err, "setup manager failed"))
	}
	if err = addNecessaryAPIToScheme(mgr.GetScheme()); err != nil {
		panic(errors.Wrap(err, "add necessary api to scheme failed"))
	}

	// TODO: There are config dependency, this manager should have it's config.
	if err = initServerConfig(&mgr); err != nil {
		panic(errors.Wrap(err, "init server config failed"))
	}
	if err = initLogger(); err != nil {
		panic(errors.Wrap(err, "init logger failed"))
	}
	if err = initThirdPartyClient(); err != nil {
		panic(errors.Wrap(err, "init third party client failed"))
	}
	if err = initClusterUID(); err != nil {
		panic(errors.Wrap(err, "init cluster uid failed"))
	}
	if err = initDatahubSchemas(context.TODO()); err != nil {
		panic(errors.Wrap(err, "init Datahub schemas failed"))
	}
	initDatahubResourceRepsitories()

	scope.Info("Adding controllers to manager...")
	if err := addControllersToManager(mgr); err != nil {
		panic(errors.Wrap(err, "add necessary controllers to manager failed"))
	}

	// Start components
	wg, ctx := errgroup.WithContext(context.Background())
	wg.Go(
		func() error {
			scope.Info("Starting the Cmd.")
			return mgr.Start(ctrl.SetupSignalHandler())
		})
	wg.Go(
		func() error {
			// To use instance from return value of function mgr.GetClient(),
			// block till the cache is synchronized, or the cache will be empty and get/list nothing.
			ok := mgr.GetCache().WaitForCacheSync(ctx.Done())
			if !ok {
				scope.Error("Wait for cache synchronization failed")
			} else {
				go syncResourcesWithDatahub(mgr.GetClient(),
					datahubConn)
			}
			return nil
		})
	if err := wg.Wait(); err != nil {
		panic(err)
	}
	return
}
