package entities

import (
	"time"
)

type MetricKafkaTopicCurrentOffset struct {
	DatahubEntity `scope:"metric" category:"kafka" type:"topic" measurement:"kafka_topic_partition_current_offset" metric:"current_offset" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricKafkaConsumerGroupCurrentOffset struct {
	DatahubEntity `scope:"metric" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group_current_offset" metric:"current_offset" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	TopicName     string     `json:"topic_name"   required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricKafkaConsumerGroupLag struct {
	DatahubEntity `scope:"metric" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group_lag" metric:"lag" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	TopicName     string     `json:"topic_name"   required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}
type MetricClusterStatusApplicationCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"application" measurement:"application_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusApplicationMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"application" measurement:"application_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusClusterCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"cluster" measurement:"cluster_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"  required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"  required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"   required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value" required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusClusterMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"cluster" measurement:"cluster_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"  required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"  required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"   required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value" required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusContainerCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	RateRange     string     `json:"rate_range"    required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusContainerMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	RateRange     string     `json:"rate_range"    required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

// Resource requests cpu millicores
type MetricClusterStatusContainerRscReqCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_resource_requests_cpu_millicores" metric:"cpu_millicores_total" boundary:"undefined" quota:"request"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

// Resource limits cpu millicores
type MetricClusterStatusContainerRscLimitCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_resource_limits_cpu_millicores" metric:"cpu_millicores_total" boundary:"undefined" quota:"limit"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

// Resource requests memory bytes
type MetricClusterStatusContainerRscReqMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_resource_requests_memory_bytes" metric:"memory_bytes_total" boundary:"undefined" quota:"request"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

// Resource limits memory bytes
type MetricClusterStatusContainerRscLimitMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_resource_limits_memory_bytes" metric:"memory_bytes_total" boundary:"undefined" quota:"limit"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusContainerRestartsTotal struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_restarts_total" metric:"restarts_total" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusControllerCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"controller" measurement:"controller_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Kind          string     `json:"kind"         required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusControllerMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"controller" measurement:"controller_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Kind          string     `json:"kind"         required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNamespaceCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"namespace" measurement:"namespace_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"  type:"time"`
	Name          string     `json:"name"         required:"true" column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true" column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true" column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true" column:"field" type:"float64"`
}

type MetricClusterStatusNamespaceMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"namespace" measurement:"namespace_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeCPUAllocatable struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_cpu_cores_allocatable" metric:"cpu_cores_alloc" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeCPUTotal struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_cpu_millicores_total" metric:"cpu_millicores_total" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeCPUUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_cpu" metric:"cpu_millicores_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeMemoryTotal struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_memory_bytes_total" metric:"memory_bytes_total" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_memory" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

// Filesystem bytes usage percentage
type MetricClusterStatusNodeFSPCT struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_fs_bytes_usage_pct" metric:"fs_bytes_usage_pct" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Device        string     `json:"device"       required:"false" column:"field" type:"string"`
	FSType        string     `json:"fs_type"      required:"false" column:"field" type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

// Disk I/O utilization
type MetricClusterStatusNodeDiskIOUtil struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_disk_io_util" metric:"disk_io_util" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodePodPhaseCount struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_pod_phase_count" metric:"number" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Pending       int64      `json:"pending"      required:"true"  column:"field" type:"int64"`
	Running       int64      `json:"running"      required:"true"  column:"field" type:"int64"`
	Succeeded     int64      `json:"succeeded"    required:"true"  column:"field" type:"int64"`
	Failed        int64      `json:"failed"       required:"true"  column:"field" type:"int64"`
	Unknown       int64      `json:"unknown"      required:"true"  column:"field" type:"int64"`
}

type MetricClusterStatusNodeUnschedulable struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_unschedulable" metric:"unschedulable" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusServiceHealth struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"service" measurement:"service_health" metric:"health" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Job           string     `json:"job"           required:"true"  column:"tag"   type:"string"`
	Value         int32      `json:"value"         required:"true"  column:"field" type:"int32"` // 1: healthy, 0: failed
}

type MetricTopContainerCPUUsagePCT struct {
	DatahubEntity `scope:"metric" category:"top" type:"container" measurement:"top_container_cpu_millicores_usage_pct" metric:"cpu_millicores_usage_pct" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type MetricTopContainerMemoryUsage struct {
	DatahubEntity `scope:"metric" category:"top" type:"container" measurement:"top_container_memory_bytes_usage" metric:"memory_bytes_usage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}
