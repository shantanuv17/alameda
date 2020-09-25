package metrics

import (
	"prophetstor.com/alameda/pkg/consts"
	"prophetstor.com/alameda/pkg/database/influxdb"
)

const (
	ContainerCpu      influxdb.Measurement = "container_cpu"
	ContainerMemory   influxdb.Measurement = "container_memory"
	NodeCpu           influxdb.Measurement = "node_cpu"
	NodeMemory        influxdb.Measurement = "node_memory"
	ApplicationCpu    influxdb.Measurement = "application_cpu"
	ApplicationMemory influxdb.Measurement = "application_memory"
	ClusterCpu        influxdb.Measurement = "cluster_cpu"
	ClusterMemory     influxdb.Measurement = "cluster_memory"
	NamespaceCpu      influxdb.Measurement = "namespace_cpu"
	NamespaceMemory   influxdb.Measurement = "namespace_memory"
	ControllerCpu     influxdb.Measurement = "controller_cpu"
	ControllerMemory  influxdb.Measurement = "controller_memory"
)

var PodNameRegularExpression = map[string]string{
	"DEPLOYMENT":       consts.DeploymentPodFormat,
	"STATEFULSET":      consts.StatefulSetPodFormat,
	"DEPLOYMENTCONFIG": consts.DeploymentConfigPodFormat,
}
