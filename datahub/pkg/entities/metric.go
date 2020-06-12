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

type MetricNginxHttpResponseTotal struct {
	DatahubEntity               `scope:"metric" category:"nginx" type:"nginx" measurement:"nginx_http_response_total" metric:"number" boundary:"undefined" quota:"undefined"`
	Time                        *time.Time `json:"time"                           required:"false" column:"tag"   type:"time"`
	ClusterName                 string     `json:"cluster_name"                   required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceName      string     `json:"resource_k8s_service_name"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceNamespace string     `json:"resource_k8s_service_namespace" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName             string     `json:"resource_k8s_name"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace        string     `json:"resource_k8s_namespace"         required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind             string     `json:"resource_k8s_kind"              required:"true"  column:"tag"   type:"string"`
	Value                       float64    `json:"value"                          required:"true"  column:"field" type:"float64"`
}

type MetricNginxHttpResponseLatency struct {
	DatahubEntity               `scope:"metric" category:"nginx" type:"nginx" measurement:"nginx_http_response_latency_ms" metric:"latency" boundary:"undefined" quota:"undefined"`
	Time                        *time.Time `json:"time"                           required:"false" column:"tag"   type:"time"`
	ClusterName                 string     `json:"cluster_name"                   required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceName      string     `json:"resource_k8s_service_name"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceNamespace string     `json:"resource_k8s_service_namespace" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName             string     `json:"resource_k8s_name"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace        string     `json:"resource_k8s_namespace"         required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind             string     `json:"resource_k8s_kind"              required:"true"  column:"tag"   type:"string"`
	Value                       float64    `json:"value"                          required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusApplicationCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"application" measurement:"application_cpu" metric:"cpu_usage_percentage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusApplicationMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"application" measurement:"application_memory" metric:"memory_usage_bytes" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusClusterCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"cluster" measurement:"cluster_cpu" metric:"cpu_usage_percentage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"  required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"  required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"   required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value" required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusClusterMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"cluster" measurement:"cluster_memory" metric:"memory_usage_bytes" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"  required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"  required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"   required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value" required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusContainerCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_cpu" metric:"cpu_usage_percentage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	RateRange     string     `json:"rate_range"    required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusContainerMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"container" measurement:"container_memory" metric:"memory_usage_bytes" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	PodNamespace  string     `json:"pod_namespace" required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	RateRange     string     `json:"rate_range"    required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusControllerCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"controller" measurement:"controller_cpu" metric:"cpu_usage_percentage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Kind          string     `json:"kind"         required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusControllerMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"controller" measurement:"controller_memory" metric:"memory_usage_bytes" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Kind          string     `json:"kind"         required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNamespaceCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"namespace" measurement:"namespace_cpu" metric:"cpu_usage_percentage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"  type:"time"`
	Name          string     `json:"name"         required:"true" column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true" column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true" column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true" column:"field" type:"float64"`
}

type MetricClusterStatusNamespaceMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"namespace" measurement:"namespace_memory" metric:"memory_usage_bytes" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeCPU struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_cpu" metric:"cpu_usage_percentage" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeMemory struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_memory" metric:"memory_usage_bytes" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}

type MetricClusterStatusNodeCPUCores struct {
	DatahubEntity `scope:"metric" category:"cluster_status" type:"node" measurement:"node_cpu_cores" metric:"cpu_cores" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         float64    `json:"value"        required:"true"  column:"field" type:"float64"`
}
