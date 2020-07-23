package clusterstatus

import (
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
	"time"
)

const (
	NodeName                influxdb.Tag = "name" // NodeName is the name of node
	NodeClusterName         influxdb.Tag = "cluster_name"
	NodeUid                 influxdb.Tag = "uid"
	NodeMachinesetName      influxdb.Tag = "machineset_name"
	NodeMachinesetNamespace influxdb.Tag = "machineset_namespace"

	NodeCreateTime        influxdb.Field = "create_time"
	NodeMachineCreateTime influxdb.Field = "machine_create_time"
	NodeCPUCores          influxdb.Field = "node_cpu_cores"    // NodeCPUCores is the amount of cores in node
	NodeMemoryBytes       influxdb.Field = "node_memory_bytes" // NodeMemoryBytes is the amount of memory bytes in node
	NodeNetworkMbps       influxdb.Field = "node_network_mbps" // NodeNetworkMbps is mega bits per second
	NodeIOProvider        influxdb.Field = "io_provider"       // Cloud service provider
	NodeIOInstanceType    influxdb.Field = "io_instance_type"
	NodeIORegion          influxdb.Field = "io_region"
	NodeIOZone            influxdb.Field = "io_zone"
	NodeIOOS              influxdb.Field = "io_os"
	NodeIORole            influxdb.Field = "io_role"
	NodeIOInstanceID      influxdb.Field = "io_instance_id"
	NodeIOStorageSize     influxdb.Field = "io_storage_size"
	NodeRoleMaster        influxdb.Field = "role_master"
	NodeRoleWorker        influxdb.Field = "role_worker"
	NodeRoleInfra         influxdb.Field = "role_infra"
)

var (
	// NodeTags list tags of node measurement
	NodeTags = []influxdb.Tag{
		NodeName,
		NodeClusterName,
		NodeUid,
		NodeMachinesetName,
		NodeMachinesetNamespace,
	}

	// NodeFields list fields of node measurement
	NodeFields = []influxdb.Field{
		NodeCreateTime,
		NodeMachineCreateTime,
		NodeCPUCores,
		NodeMemoryBytes,
		NodeNetworkMbps,
		NodeIOProvider,
		NodeIOInstanceType,
		NodeIORegion,
		NodeIOZone,
		NodeIOOS,
		NodeIORole,
		NodeIOInstanceID,
		NodeIOStorageSize,
		NodeRoleMaster,
		NodeRoleWorker,
		NodeRoleInfra,
	}
)

// NodeEntity is entity in database
type NodeEntity struct {
	Time                time.Time
	Name                string
	ClusterName         string
	Uid                 string
	MachinesetName      string
	MachinesetNamespace string

	CreateTime        int64
	MachineCreateTime int64
	CPUCores          int64
	MemoryBytes       int64
	NetworkMbps       int64
	IOProvider        string
	IOInstanceType    string
	IORegion          string
	IOZone            string
	IOOS              string
	IORole            string
	IOInstanceID      string
	IOStorageSize     int64
	RoleMaster        bool
	RoleWorker        bool
	RoleInfra         bool
}

// NewNodeEntityFromMap Build entity from map
func NewNodeEntity(data map[string]string) *NodeEntity {
	entity := NodeEntity{}

	tempTimestamp, _ := utils.ParseTime(data["time"])
	entity.Time = tempTimestamp

	// InfluxDB tags
	if value, exist := data[string(NodeName)]; exist {
		entity.Name = value
	}
	if value, exist := data[string(NodeClusterName)]; exist {
		entity.ClusterName = value
	}
	if value, exist := data[string(NodeUid)]; exist {
		entity.Uid = value
	}
	if value, exist := data[string(NodeMachinesetName)]; exist {
		entity.MachinesetName = value
	}
	if value, exist := data[string(NodeMachinesetNamespace)]; exist {
		entity.MachinesetNamespace = value
	}

	// InfluxDB fields
	if value, exist := data[string(NodeCreateTime)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.CreateTime = valueInt64
	}
	if value, exist := data[string(NodeMachineCreateTime)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.MachineCreateTime = valueInt64
	}
	if value, exist := data[string(NodeCPUCores)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.CPUCores = valueInt64
	}
	if value, exist := data[string(NodeMemoryBytes)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.MemoryBytes = valueInt64
	}
	if value, exist := data[string(NodeNetworkMbps)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.NetworkMbps = valueInt64
	}
	if value, exist := data[string(NodeIOProvider)]; exist {
		entity.IOProvider = value
	}
	if value, exist := data[string(NodeIOInstanceType)]; exist {
		entity.IOInstanceType = value
	}
	if value, exist := data[string(NodeIORegion)]; exist {
		entity.IORegion = value
	}
	if value, exist := data[string(NodeIOZone)]; exist {
		entity.IOZone = value
	}
	if value, exist := data[string(NodeIOOS)]; exist {
		entity.IOOS = value
	}
	if value, exist := data[string(NodeIORole)]; exist {
		entity.IORole = value
	}
	if value, exist := data[string(NodeIOInstanceID)]; exist {
		entity.IOInstanceID = value
	}
	if value, exist := data[string(NodeIOStorageSize)]; exist {
		valueInt64, _ := strconv.ParseInt(value, 10, 64)
		entity.IOStorageSize = valueInt64
	}
	if value, exist := data[string(NodeRoleMaster)]; exist {
		valueBool, _ := strconv.ParseBool(value)
		entity.RoleMaster = valueBool
	}
	if value, exist := data[string(NodeRoleWorker)]; exist {
		valueBool, _ := strconv.ParseBool(value)
		entity.RoleWorker = valueBool
	}
	if value, exist := data[string(NodeRoleInfra)]; exist {
		valueBool, _ := strconv.ParseBool(value)
		entity.RoleInfra = valueBool
	}

	return &entity
}

func (p *NodeEntity) BuildInfluxPoint(measurement string) (*InfluxClient.Point, error) {
	// Pack influx tags
	tags := map[string]string{
		string(NodeName):                p.Name,
		string(NodeClusterName):         p.ClusterName,
		string(NodeUid):                 p.Uid,
		string(NodeMachinesetName):      p.MachinesetName,
		string(NodeMachinesetNamespace): p.MachinesetNamespace,
	}

	// Pack influx fields
	fields := map[string]interface{}{
		string(NodeCreateTime):        p.CreateTime,
		string(NodeMachineCreateTime): p.MachineCreateTime,
		string(NodeCPUCores):          p.CPUCores,
		string(NodeMemoryBytes):       p.MemoryBytes,
		string(NodeNetworkMbps):       p.NetworkMbps,
		string(NodeIOProvider):        p.IOProvider,
		string(NodeIOInstanceType):    p.IOInstanceType,
		string(NodeIORegion):          p.IORegion,
		string(NodeIOZone):            p.IOZone,
		string(NodeIOOS):              p.IOOS,
		string(NodeIORole):            p.IORole,
		string(NodeIOInstanceID):      p.IOInstanceID,
		string(NodeIOStorageSize):     p.IOStorageSize,
		string(NodeRoleMaster):        p.RoleMaster,
		string(NodeRoleWorker):        p.RoleWorker,
		string(NodeRoleInfra):         p.RoleInfra,
	}

	return InfluxClient.NewPoint(measurement, tags, fields, p.Time)
}
