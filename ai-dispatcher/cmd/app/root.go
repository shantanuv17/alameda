package app

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/dispatcher"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/metrics"
	alameda_app "github.com/containers-ai/alameda/cmd/app"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	cfgFile             string
	logRotateOutputFile string

	scope *log.Scope
)

func launchMetricServer() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}

const (
	envVarPrefix = "ALAMEDA_AI_DISPATCHER"

	defaultRotationMaxSizeMegabytes = 100
	defaultRotationMaxBackups       = 7
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(alameda_app.VersionCmd)
	rootCmd.AddCommand(ProbeCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config",
		"/etc/alameda/ai-dispatcher/ai-dispatcher.toml",
		"The path to ai-dispatcher configuration file.")
	rootCmd.PersistentFlags().StringVar(&logRotateOutputFile, "log-output-file",
		"/var/log/alameda/alameda-ai-dispatcher.log", "The path of log file.")
}

var rootCmd = &cobra.Command{
	Use:   "ai-dispatcher",
	Short: "AI dispatcher sends predicted jobs to queue",
	Long: `AI dispatcher send predicted jobs to queue
			including nodes and pods`,
	Run: func(cmd *cobra.Command, args []string) {
		initLogger()
		setLoggerScopesWithConfig()

		cfg := &config.Config{}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			panic(err)
		}
		cfg.Init()

		retryIntervalSec := 1
		if viper.IsSet("datahub.query.retryInterval") {
			retryIntervalSec = viper.GetInt("datahub.query.retryInterval")
		}
		datahubAddr := viper.GetString("datahub.address")
		if datahubAddr == "" {
			scope.Errorf("No configuration of datahub address.")
			return
		}
		datahubClient := datahubpkg.NewClient(datahubAddr)
		if datahubClient == nil {
			scope.Errorf("New datahub client failed.")
			return
		}
		queueURL := viper.GetString("queue.url")
		if queueURL == "" {
			scope.Errorf("No configuration of queue url.")
			return
		}

		for {
			amqConn, err := amqp.Dial(queueURL)
			if err == nil {
				amqConn.Close()
				break
			} else {
				scope.Errorf("connect queue failed on init: %s", err.Error())
			}
			time.Sleep(time.Duration(retryIntervalSec) * time.Second)
		}

		for {
			if checkResourceIsExist(datahubClient) {
				break
			}
			time.Sleep(time.Duration(retryIntervalSec) * time.Second)
		}

		metricExporter := metrics.NewExporter()
		go launchMetricServer()
		granularities := viper.GetStringSlice("serviceSetting.granularities")
		predictUnits := viper.GetStringSlice("serviceSetting.predictUnits")
		modelMapper := dispatcher.NewModelMapper()
		if viper.GetBool("model.enabled") {
			go dispatcher.ModelCompleteNotification(modelMapper, metricExporter)
		}
		dp := dispatcher.NewDispatcher(datahubClient, granularities, predictUnits,
			modelMapper, metricExporter, cfg)
		dp.Start()
	},
}

func checkResourceIsExist(datahubClient *datahubpkg.Client) bool {
	nodeResult, err := datahubClient.ListNodes(&datahub_resources.ListNodesRequest{})
	nodeCount := len(nodeResult.GetNodes())
	if err != nil || nodeCount <= 0 {
		if err != nil {
			scope.Errorf("ListNodes failed on init: %s", err.Error())
		}
		if nodeCount <= 0 {
			scope.Errorf("ListNodes is empty on init")
		}
		return false
	}
	clusterResult, err := datahubClient.ListClusters(&datahub_resources.ListClustersRequest{})
	clusterCount := len(clusterResult.GetClusters())
	if err != nil || clusterCount <= 0 {
		if err != nil {
			scope.Errorf("ListClusters failed on init: %s", err.Error())
		}
		if clusterCount <= 0 {
			scope.Errorf("ListClusters is empty on init")
		}
		return false
	}
	return true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		scope.Errorf("%s", err.Error())
		os.Exit(1)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix(envVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initLogger() {

	opt := log.DefaultOptions()
	opt.RotationMaxSize = defaultRotationMaxSizeMegabytes
	opt.RotationMaxBackups = defaultRotationMaxBackups
	opt.RotateOutputPath = logRotateOutputFile

	err := log.Configure(opt)
	if err != nil {
		panic(err)
	}

	scope = log.RegisterScope("app", "ai-dispatcher app", 0)
}

func setLoggerScopesWithConfig() {
	setLogCallers := true
	if viper.IsSet("log.setLogcallers") {
		setLogCallers = viper.GetBool("log.setLogcallers")
	}

	outputLevel := "none"
	if viper.IsSet("log.outputLevel") {
		outputLevel = viper.GetString("log.outputLevel")
	}

	stackTraceLvl := "none"
	if viper.IsSet("log.stacktraceLevel") {
		stackTraceLvl = viper.GetString("log.stacktraceLevel")
	}
	for _, scope := range log.Scopes() {
		scope.SetLogCallers(setLogCallers)
		if outputLvl, ok := log.StringToLevel(outputLevel); ok {
			scope.SetOutputLevel(outputLvl)
		}
		if stacktraceLevel, ok := log.StringToLevel(stackTraceLvl); ok {
			scope.SetStackTraceLevel(stacktraceLevel)
		}
	}
}
