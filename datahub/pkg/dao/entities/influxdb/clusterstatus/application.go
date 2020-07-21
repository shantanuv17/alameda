package clusterstatus

import (
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"time"
)

const (
	ApplicationName        influxdb.Tag   = "name"
	ApplicationNamespace   influxdb.Tag   = "namespace"
	ApplicationClusterName influxdb.Tag   = "cluster_name"
	ApplicationUid         influxdb.Tag   = "uid"
	ApplicationDummy       influxdb.Field = "dummy"
)

var (
	ApplicationTags = []influxdb.Tag{
		ApplicationName,
		ApplicationNamespace,
		ApplicationClusterName,
		ApplicationUid,
	}

	ApplicationFields = []influxdb.Field{
		ApplicationDummy,
	}

	ApplicationColumns = []string{
		string(ApplicationName),
		string(ApplicationNamespace),
		string(ApplicationClusterName),
		string(ApplicationUid),
		string(ApplicationDummy),
	}
)

type ApplicationEntity struct {
	// InfluxDB tags
	Time        time.Time
	Name        string
	Namespace   string
	ClusterName string
	Uid         string

	// InfluxDB fields
	Dummy string
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

	// InfluxDB fields
	entity.Dummy = ""

	return &entity
}

func (p *ApplicationEntity) BuildInfluxPoint(measurement string) (*InfluxClient.Point, error) {
	// Pack influx tags
	tags := map[string]string{
		string(ApplicationName):        p.Name,
		string(ApplicationNamespace):   p.Namespace,
		string(ApplicationClusterName): p.ClusterName,
		string(ApplicationUid):         p.Uid,
	}

	// Pack influx fields
	fields := map[string]interface{}{
		string(ApplicationDummy): p.Dummy,
	}

	return InfluxClient.NewPoint(measurement, tags, fields, p.Time)
}
