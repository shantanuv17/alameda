package dispatcher

import (
	"encoding/json"
	"time"

	"github.com/spf13/viper"
	"prophetstor.com/alameda/ai-dispatcher/consts"
	"prophetstor.com/alameda/ai-dispatcher/pkg/metrics"
	"prophetstor.com/alameda/ai-dispatcher/pkg/queue"
)

type modelCompleteMsg struct {
	UnitType        string `json:"unit_type"`
	DataGranularity string `json:"data_granularity"`
	JobCreateTime   int64  `json:"job_create_time"`
	ClusterName     string `json:"cluster_name"`
	MetricTypeStr   string `json:"metric_type_str"`
	ContainerName   string `json:"container_name"`

	Unit  unit   `json:"unit"`
	JobID string `json:"job_id"`
}

type unit struct {
	Name           string         `json:"name"`
	Host           string         `json:"host"`
	MinorNumber    string         `json:"minor_number"`
	Kind           string         `json:"kind"`
	NamespacedName namespacedName `json:"namespaced_name"`
}

type namespacedName struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func ModelCompleteNotification(modelMapper *ModelMapper, metricExporter *metrics.Exporter) {

	reconnectInterval := viper.GetInt64("queue.consumer.reconnectInterval")
	queueConnRetryItvMS := viper.GetInt64("queue.retry.connectIntervalMs")

	modelCompleteQueue := "model_complete"
	queueURL := viper.GetString("queue.url")
	for {
		queueConn := queue.GetQueueConn(queueURL, queueConnRetryItvMS)
		queueConsumer := queue.NewRabbitMQConsumer(queueConn)
		msgCH, err := queueConsumer.ConsumeJsonString(modelCompleteQueue)
		if err != nil {
			if queueConn != nil {
				queueConn.Close()
			}
			scope.Warnf("Consume message from model complete queue error: %s", err.Error())
			time.Sleep(time.Duration(reconnectInterval) * time.Second)
			continue
		}
		for cmsg := range msgCH {
			msg := string(cmsg.Body)
			scope.Debugf("receive complete model body is %s", msg)
			var msgMap modelCompleteMsg
			msgByte := []byte(msg)
			if err := json.Unmarshal(msgByte, &msgMap); err != nil {
				scope.Errorf("decode model complete job from queue failed: %s", err.Error())
				break
			}

			unitType := msgMap.UnitType
			dataGranularity := msgMap.DataGranularity
			jobCreateTime := msgMap.JobCreateTime

			if unitType == consts.UnitTypeNode {
				nodeName := msgMap.Unit.Name
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity, metricType, map[string]string{
					"name": nodeName,
				})

				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[NODE][%s][%s] Export metric %s model time value %v",
					dataGranularity, nodeName, metricType, mt)
				metricExporter.ExportNodeMetricModelTime(clusterID, nodeName, dataGranularity, metricType,
					time.Now().Unix(), float64(mt))
			} else if unitType == consts.UnitTypePod {
				podNS := msgMap.Unit.NamespacedName.Namespace
				podName := msgMap.Unit.NamespacedName.Name
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				ctName := msgMap.ContainerName
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity, metricType, map[string]string{
					"namespace":     podNS,
					"name":          podName,
					"containerName": ctName,
				})

				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[POD][%s][%s/%s/%s] Export metric %s model time with value %v",
					dataGranularity, podNS, podName, ctName, metricType, mt)
				metricExporter.ExportContainerMetricModelTime(clusterID, podNS, podName, ctName, dataGranularity, metricType,
					time.Now().Unix(), float64(mt))
			} else if unitType == consts.UnitTypeGPU {
				gpuHost := msgMap.Unit.Host
				gpuMinorNumber := msgMap.Unit.MinorNumber
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity,
					metricType, map[string]string{
						"host":        gpuHost,
						"minorNumber": gpuMinorNumber,
					})
				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[GPU][%s][%s/%s] Export metric %s model time value %v",
					dataGranularity, gpuHost, gpuMinorNumber, metricType, mt)
				metricExporter.ExportGPUMetricModelTime(clusterID, gpuHost, gpuMinorNumber,
					dataGranularity, metricType, time.Now().Unix(), float64(mt))
			} else if unitType == consts.UnitTypeApplication {
				appNS := msgMap.Unit.NamespacedName.Namespace
				appName := msgMap.Unit.NamespacedName.Name
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity,
					metricType, map[string]string{
						"namespace": appNS,
						"name":      appName,
					})

				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[APPLICATION][%s][%s/%s] Export metric %s model time value %v",
					dataGranularity, appNS, appName, metricType, mt)
				metricExporter.ExportApplicationMetricModelTime(clusterID, appNS, appName,
					dataGranularity, metricType, time.Now().Unix(), float64(mt))
			} else if unitType == consts.UnitTypeNamespace {
				namespaceName := msgMap.Unit.Name
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity, metricType, map[string]string{
					"name": namespaceName,
				})

				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[NAMESPACE][%s][%s] Export metric %s model time value %v",
					dataGranularity, namespaceName, metricType, mt)
				metricExporter.ExportNamespaceMetricModelTime(clusterID, namespaceName, dataGranularity, metricType,
					time.Now().Unix(), float64(mt))
			} else if unitType == consts.UnitTypeCluster {
				clusterName := msgMap.Unit.Name
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity, metricType, map[string]string{
					"name": clusterName,
				})

				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[CLUSTER][%s][%s] Export metric %s model time value %v",
					dataGranularity, clusterName, metricType, mt)
				metricExporter.ExportClusterMetricModelTime(clusterName, dataGranularity, metricType,
					time.Now().Unix(), float64(mt))
			} else if unitType == consts.UnitTypeController {
				controllerNS := msgMap.Unit.NamespacedName.Namespace
				controllerName := msgMap.Unit.NamespacedName.Name
				kind := msgMap.Unit.Kind
				clusterID := msgMap.ClusterName
				metricType := msgMap.MetricTypeStr
				modelMapper.RemoveModelInfo(clusterID, unitType, dataGranularity,
					metricType, map[string]string{
						"namespace": controllerNS,
						"name":      controllerName,
						"kind":      kind,
					})

				mt := time.Now().Unix() - jobCreateTime
				scope.Infof("[CONTROLLER][%s][%s][%s/%s] Export metric %s model time value %v",
					kind, dataGranularity, controllerNS, controllerName, metricType, mt)
				metricExporter.ExportControllerMetricModelTime(clusterID, controllerNS, controllerName,
					kind, dataGranularity, metricType, time.Now().Unix(), float64(mt))
			} else if msgMap.JobID != "" {
				modelMapper.RemoveModelInfoV2(msgMap.JobID)
				mt := time.Now().Unix() - jobCreateTime
				metrics.SetMetricModelTime(msgMap.JobID, float64(mt))
				metrics.AddMetricModelTime(msgMap.JobID, float64(mt))
				scope.Infof("[%s] Export metric model time %v seconds", msgMap.JobID, mt)
			}
		}
		scope.Warnf("Retry construct consume model complete channel")
		if queueConn != nil {
			queueConn.Close()
		}
		time.Sleep(time.Duration(reconnectInterval) * time.Second)
	}
}
