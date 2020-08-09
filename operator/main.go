/*
Copyright 2020 The Alameda Authors.

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
	"os"
	"strings"
	"time"

	datahub_client_cluster "github.com/containers-ai/alameda/operator/datahub/client/cluster"
	datahub_client_kafka "github.com/containers-ai/alameda/operator/datahub/client/kafka"
	datahub_client_nginx "github.com/containers-ai/alameda/operator/datahub/client/nginx"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	alamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	routeapi_v1 "github.com/openshift/api/route/v1"
	openshift_machineapi_v1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	kafkaclient "github.com/containers-ai/alameda/internal/pkg/message-queue/kafka/client"

	"github.com/containers-ai/alameda/pkg/provider"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	"github.com/containers-ai/alameda/operator/controllers"
	datahub_client_application "github.com/containers-ai/alameda/operator/datahub/client/application"
	datahub_client_controller "github.com/containers-ai/alameda/operator/datahub/client/controller"
	datahub_client_namespace "github.com/containers-ai/alameda/operator/datahub/client/namespace"
	datahub_client_node "github.com/containers-ai/alameda/operator/datahub/client/node"
	datahub_client_pod "github.com/containers-ai/alameda/operator/datahub/client/pod"
	"github.com/containers-ai/alameda/operator/pkg/probe"
	"github.com/containers-ai/alameda/operator/pkg/utils"

	osappsapi "github.com/openshift/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
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
	operatorConf                       Config
	scope                              *logUtil.Scope
	alamedaScalerKafkaControllerLogger *logUtil.Scope
	alamedaScalerNginxControllerLogger *logUtil.Scope
	datahubClientLogger                *logUtil.Scope

	clusterUID string

	// Third party clients
	k8sClient        client.Client
	datahubConn      *grpc.ClientConn
	datahubClient    *datahubpkg.Client
	kafkaClient      kafka.Client
	prometheusClient prometheus.Prometheus
)

func init() {
	flag.BoolVar(&showVer, "version", false, "show version")
	flag.BoolVar(&readinessProbeFlag, "readiness-probe", false, "probe for readiness")
	flag.BoolVar(&livenessProbeFlag, "liveness-probe", false, "probe for liveness")
	flag.StringVar(&operatorConfigFile, "config", "/etc/alameda/operator/operator.toml",
		"File path to operator coniguration")
	flag.StringVar(&crdLocation, "crd-location", "/etc/alameda/operator/crds", "CRD location")
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", true,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")

	scope = logUtil.RegisterScope("main", "operator entry point", 0)
	alamedaScalerKafkaControllerLogger = logUtil.RegisterScope("alameda_scaler_kafka_controller", "AlamedaScaler Kafka Controller", 0)
	alamedaScalerNginxControllerLogger = logUtil.RegisterScope("alameda_scaler_nginx_controller", "AlamedaScaler Nginx Controller", 0)
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

	operatorConf = NewConfigWithoutMgr()
	if mgr != nil {
		operatorConf = NewConfig(*mgr)
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	datahubConn, err = grpc.DialContext(ctx, operatorConf.Datahub.Address, grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(3)))))
	if err != nil {
		return errors.Wrap(err, "new connection to datahub failed")
	}
	datahubClient = datahubpkg.NewClient(operatorConf.Datahub.Address)

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

func setupManager() (manager.Manager, error) {
	return ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		MetricsBindAddress:      metricsAddr,
		LeaderElection:          enableLeaderElection,
		LeaderElectionNamespace: alamedaUtils.GetRunningNamespace(),
		Port:                    9443,
		LeaderElectionID:        "alameda-operator-leader-election",
		SyncPeriod:              &syncPriod,
	})
}

func addNecessaryAPIToScheme(scheme *runtime.Scheme) error {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = openshift_machineapi_v1beta1.AddToScheme(scheme)

	if hasOpenShiftAPIAppsv1 {
		if err := osappsapi.AddToScheme(scheme); err != nil {
			return err
		}
		if err := routeapi_v1.AddToScheme(scheme); err != nil {
			return err
		}
	}
	_ = autoscalingv1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
	return nil
}

func addControllersToManager(mgr manager.Manager) error {
	datahubControllerRepo := datahub_client_controller.NewControllerRepository(
		datahubConn, clusterUID)
	datahubPodRepo := datahub_client_pod.NewPodRepository(datahubConn, clusterUID)
	var err error

	if err = (&controllers.AlamedaScalerReconciler{
		Client:                 mgr.GetClient(),
		Scheme:                 mgr.GetScheme(),
		ClusterUID:             clusterUID,
		DatahubControllerRepo:  datahubControllerRepo,
		DatahubPodRepo:         datahubPodRepo,
		DatahubClient:          datahubClient,
		ReconcileTimeout:       3 * time.Second,
		ForceReconcileInterval: 1 * time.Minute,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if err = (&controllers.AlamedaRecommendationReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		ClusterUID:    clusterUID,
		DatahubClient: datahubClient,
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

		if err = (&controllers.MachineSetReconciler{
			Client:           mgr.GetClient(),
			Scheme:           mgr.GetScheme(),
			ClusterUID:       clusterUID,
			ReconcileTimeout: 3 * time.Second,
			DatahubClient:    datahubClient,
		}).SetupWithManager(mgr); err != nil {
			return err
		}

		if err = (&controllers.MachineReconciler{
			Client:           mgr.GetClient(),
			Scheme:           mgr.GetScheme(),
			ClusterUID:       clusterUID,
			ReconcileTimeout: 3 * time.Second,
			DatahubClient:    datahubClient,
		}).SetupWithManager(mgr); err != nil {
			return err
		}
	}

	if err = (&controllers.NamespaceReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		DatahubClient: datahubClient,
		ClusterUID:    clusterUID,
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
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		ClusterUID:    clusterUID,
		Cloudprovider: cloudprovider,
		RegionName:    regionName,
		DatahubClient: datahubClient,
		DatahubNodeRepo: *datahub_client_node.NewNodeRepository(
			datahubClient, clusterUID),
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
		DatahubClient:         datahubClient,
		ClusterUID:            clusterUID,
		HasOpenShiftAPIAppsv1: hasOpenShiftAPIAppsv1,
		K8SClient:             mgr.GetClient(),
		Scheme:                mgr.GetScheme(),
		KafkaClient:           kafkaClient,
		PrometheusClient:      prometheusClient,
		ReconcileTimeout:      3 * time.Second,
		Logger:                alamedaScalerKafkaControllerLogger,
		NeededMetrics:         operatorConf.Prometheus.RequiredMetrics,
	}).SetupWithManager(mgr); err != nil {
		return err
	}
	if err = (&controllers.AlamedaScalerNginxReconciler{
		ClusterUID:            clusterUID,
		Client:                mgr.GetClient(),
		Scheme:                mgr.GetScheme(),
		ReconcileTimeout:      3 * time.Second,
		DatahubClient:         datahubClient,
		HasOpenShiftAPIAppsv1: hasOpenShiftAPIAppsv1,
		Logger:                alamedaScalerNginxControllerLogger,
	}).SetupWithManager(mgr); err != nil {
		return err
	}

	if err = (&controllers.AlamedaMachineGroupScalerReconciler{
		ClusterUID: clusterUID,
		Client:     mgr.GetClient(),
		Log: ctrl.Log.WithName("controllers").WithName(
			"AlamedaMachineGroupScaler"),
		Scheme:           mgr.GetScheme(),
		DatahubClient:    datahubClient,
		ReconcileTimeout: 3 * time.Second,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err,
			"unable to create controller", "controller", "AlamedaMachineGroupScaler")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder
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
func printSoftwareInfo() {
	scope.Infof(fmt.Sprintf("Alameda Version: %s", VERSION))
	scope.Infof(fmt.Sprintf("Alameda Build Time: %s", BUILD_TIME))
	scope.Infof(fmt.Sprintf("Alameda GO Version: %s", GO_VERSION))
}

func syncResourcesWithDatahub(client client.Client, datahubConn *grpc.ClientConn) {
	for {
		clusterUID, err := k8sutils.GetClusterUID(client)
		if err == nil {
			scope.Infof(
				"Get cluster UID %s successfully, and then try synchronzing resources with datahub.",
				clusterUID)
			break
		} else {
			scope.Infof("Sync resources with datahub failed. %s", err.Error())
		}
		time.Sleep(time.Duration(1) * time.Second)
	}

	go func() {
		if err := datahub_client_namespace.SyncWithDatahub(
			client, datahubClient); err != nil {
			scope.Errorf("sync namespace failed at start due to %s", err.Error())
		}
	}()
	go func() {
		if err := datahub_client_node.SyncWithDatahub(client, datahubClient); err != nil {
			scope.Errorf("sync node failed at start due to %s", err.Error())
		}
	}()

	go func() {
		if err := datahub_client_application.SyncWithDatahub(
			client, datahubClient); err != nil {
			scope.Errorf("sync application failed at start due to %s", err.Error())
		}
	}()

	go func() {
		if err := datahub_client_cluster.SyncWithDatahub(client,
			datahubConn); err != nil {
			scope.Errorf("sync cluster failed at start due to %s", err.Error())
		}
	}()
	go func() {
		if err := datahub_client_controller.SyncWithDatahub(client,
			datahubConn); err != nil {
			scope.Errorf("sync controller failed at start due to %s", err.Error())
		}
	}()
	go func() {
		if err := datahub_client_pod.SyncWithDatahub(client,
			datahubConn); err != nil {
			scope.Errorf("sync pod failed at start due to %s", err.Error())
		}
	}()
	go func() {
		if err := datahub_client_kafka.SyncWithDatahub(client, datahubClient); err != nil {
			scope.Errorf("sync kafka failed at start due to %s", err.Error())
		}
	}()
	go func() {
		if err := datahub_client_nginx.SyncWithDatahub(client, datahubClient); err != nil {
			scope.Errorf("sync nginx failed at start due to %s", err.Error())
		}
	}()

}
func livenessProbe(cfg *probe.LivenessProbeConfig) {
	// probe.LivenessProbe(cfg)
	os.Exit(0)
}

func readinessProbe(cfg *probe.ReadinessProbeConfig) {
	// probe.ReadinessProbe(cfg)
	os.Exit(0)
}
