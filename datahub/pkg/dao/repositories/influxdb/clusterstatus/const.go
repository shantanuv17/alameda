package clusterstatus

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
)

const (
	Container   influxdb.Measurement = "container"
	Pod         influxdb.Measurement = "pod"
	Controller  influxdb.Measurement = "controller"
	Application influxdb.Measurement = "application"
	Namespace   influxdb.Measurement = "namespace"
	Node        influxdb.Measurement = "node"
	Cluster     influxdb.Measurement = "cluster"
	Keycode     influxdb.Measurement = "keycode"
)
