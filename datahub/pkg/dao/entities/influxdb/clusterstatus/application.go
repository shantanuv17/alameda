package clusterstatus

import (
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
	"time"
)

const (
	ApplicationName        influxdb.Tag   = "name"
	ApplicationNamespace   influxdb.Tag   = "namespace"
	ApplicationClusterName influxdb.Tag   = "cluster_name"
	ApplicationUid         influxdb.Tag   = "uid"
	ApplicationScalingTool influxdb.Tag   = "scaling_tool"
	ApplicationMinReplicas influxdb.Field = "resource_k8s_min_replicas"
	ApplicationMaxReplicas influxdb.Field = "resource_k8s_max_replicas"
)

var (
	ApplicationTags = []influxdb.Tag{
		ApplicationName,
		ApplicationNamespace,
		ApplicationClusterName,
		ApplicationUid,
		ApplicationScalingTool,
	}

	ApplicationFields = []influxdb.Field{
		ApplicationMinReplicas,
		ApplicationMaxReplicas,
	}

	ApplicationColumns = []string{
		string(ApplicationName),
		string(ApplicationNamespace),
		string(ApplicationClusterName),
		string(ApplicationUid),
		string(ApplicationScalingTool),
		string(ApplicationMinReplicas),
		string(ApplicationMaxReplicas),
	}
)

type ApplicationEntity struct {
	// InfluxDB tags
	Time        time.Time
	Name        string
	Namespace   string
	ClusterName string
	Uid         string
	ScalingTool string

	// InfluxDB fields
	MinReplicas int32
	MaxReplicas int32
}

func NewApplicationEntity(data map[string]string) *ApplicationEntity {
	entity := ApplicationEntity{}

	tempTimestamp, _ := utils.ParseTime(data["time"])
	entity.Time = tempTimestamp

	// InfluxDB tags
	if value, exist := data[string(ApplicationName)]; exist {
		entity.Name = value
	}
	if value, exist := data[string(ApplicationNamespace)]; exist {
		entity.Namespace = value
	}
	if value, exist := data[string(ApplicationClusterName)]; exist {
		entity.ClusterName = value
	}
	if value, exist := data[string(ApplicationUid)]; exist {
		entity.Uid = value
	}
	if value, exist := data[string(ApplicationScalingTool)]; exist {
		entity.ScalingTool = value
	}

	// InfluxDB fields
	if value, exist := data[string(ApplicationMinReplicas)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.MinReplicas = int32(valueInt64)
	}
	if value, exist := data[string(ApplicationMaxReplicas)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.MaxReplicas = int32(valueInt64)
	}

	return &entity
}

func (p *ApplicationEntity) BuildInfluxPoint(measurement string) (*InfluxClient.Point, error) {
	// Pack influx tags
	tags := map[string]string{
		string(ApplicationName):        p.Name,
		string(ApplicationNamespace):   p.Namespace,
		string(ApplicationClusterName): p.ClusterName,
		string(ApplicationUid):         p.Uid,
		string(ApplicationScalingTool): p.ScalingTool,
	}

	// Pack influx fields
	fields := map[string]interface{}{
		string(ApplicationMinReplicas): p.MinReplicas,
		string(ApplicationMaxReplicas): p.MaxReplicas,
	}

	return InfluxClient.NewPoint(measurement, tags, fields, p.Time)
}
