package entities

import (
	"time"
)

type PredictionKafkaTopicCurrentOffset struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"topic" measurement:"kafka_topic_partition_current_offset" metric:"current_offset" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionKafkaTopicCurrentOffsetUpperBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"topic" measurement:"kafka_topic_partition_current_offset_upper_bound" metric:"current_offset" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionKafkaTopicCurrentOffsetLowerBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"topic" measurement:"kafka_topic_partition_current_offset_lower_bound" metric:"current_offset" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionKafkaConsumerGroupCurrentOffset struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group_current_offset" metric:"current_offset" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	TopicName     string     `json:"topic_name"    required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionKafkaConsumerGroupCurrentOffsetUpperBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group_current_offset_upper_bound" metric:"current_offset" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	TopicName     string     `json:"topic_name"    required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionKafkaConsumerGroupCurrentOffsetLowerBound struct {
	DatahubEntity `scope:"prediction" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group_current_offset_lower_bound" metric:"current_offset" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	TopicName     string     `json:"topic_name"    required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionNginxHttpResponseTotal struct {
	DatahubEntity               `scope:"prediction" category:"nginx" type:"nginx" measurement:"nginx_http_response_total" metric:"number" boundary:"raw" quota:"undefined"`
	Time                        *time.Time `json:"time"                           required:"false" column:"tag"   type:"time"`
	ClusterName                 string     `json:"cluster_name"                   required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceName      string     `json:"resource_k8s_service_name"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceNamespace string     `json:"resource_k8s_service_namespace" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName             string     `json:"resource_k8s_name"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace        string     `json:"resource_k8s_namespace"         required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind             string     `json:"resource_k8s_kind"              required:"true"  column:"tag"   type:"string"`
	Granularity                 string     `json:"granularity"                    required:"true"  column:"tag"   type:"string"`
	ModelId                     string     `json:"model_id"                       required:"true"  column:"field" type:"string"`
	PredictionId                string     `json:"prediction_id"                  required:"true"  column:"field" type:"string"`
	Value                       float64    `json:"value"                          required:"true"  column:"field" type:"float64"`
}

type PredictionNginxHttpResponseTotalUpperBound struct {
	DatahubEntity               `scope:"prediction" category:"nginx" type:"nginx" measurement:"nginx_http_response_total_upper_bound" metric:"number" boundary:"upper_bound" quota:"undefined"`
	Time                        *time.Time `json:"time"                           required:"false" column:"tag"   type:"time"`
	ClusterName                 string     `json:"cluster_name"                   required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceName      string     `json:"resource_k8s_service_name"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceNamespace string     `json:"resource_k8s_service_namespace" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName             string     `json:"resource_k8s_name"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace        string     `json:"resource_k8s_namespace"         required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind             string     `json:"resource_k8s_kind"              required:"true"  column:"tag"   type:"string"`
	Granularity                 string     `json:"granularity"                    required:"true"  column:"tag"   type:"string"`
	ModelId                     string     `json:"model_id"                       required:"true"  column:"field" type:"string"`
	PredictionId                string     `json:"prediction_id"                  required:"true"  column:"field" type:"string"`
	Value                       float64    `json:"value"                          required:"true"  column:"field" type:"float64"`
}

type PredictionNginxHttpResponseTotalLowerBound struct {
	DatahubEntity               `scope:"prediction" category:"nginx" type:"nginx" measurement:"nginx_http_response_total_lower_bound" metric:"number" boundary:"lower_bound" quota:"undefined"`
	Time                        *time.Time `json:"time"                           required:"false" column:"tag"   type:"time"`
	ClusterName                 string     `json:"cluster_name"                   required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceName      string     `json:"resource_k8s_service_name"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceNamespace string     `json:"resource_k8s_service_namespace" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName             string     `json:"resource_k8s_name"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace        string     `json:"resource_k8s_namespace"         required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind             string     `json:"resource_k8s_kind"              required:"true"  column:"tag"   type:"string"`
	Granularity                 string     `json:"granularity"                    required:"true"  column:"tag"   type:"string"`
	ModelId                     string     `json:"model_id"                       required:"true"  column:"field" type:"string"`
	PredictionId                string     `json:"prediction_id"                  required:"true"  column:"field" type:"string"`
	Value                       float64    `json:"value"                          required:"true"  column:"field" type:"float64"`
}

type PredictionClusterAutoscalerMachinegroupCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup_cpu" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterAutoscalerMachinegroupCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup_cpu_upper_bound" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterAutoscalerMachinegroupCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup_cpu_lower_bound" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterAutoscalerMachinegroupMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup_memory" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterAutoscalerMachinegroupMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup_memory_upper_bound" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterAutoscalerMachinegroupMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup_memory_lower_bound" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusApplicationCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application" measurement:"application" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusApplicationCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application" measurement:"application" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusApplicationCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application" measurement:"application" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusApplicationMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application" measurement:"application" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusApplicationMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application" measurement:"application" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusApplicationMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"application" measurement:"application" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusClusterCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusClusterCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusClusterCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusClusterMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusClusterMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusClusterMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusContainerCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container" measurement:"container" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusContainerCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container" measurement:"container" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusContainerCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container" measurement:"container" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusContainerMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container" measurement:"container" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusContainerMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container" measurement:"container" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusContainerMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"container" measurement:"container" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	PodName       string     `json:"pod_name"      required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"     required:"true"  column:"tag"   type:"string"`
	NodeName      string     `json:"node_name"     required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusControllerCPU struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller" measurement:"controller" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time           *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	Name           string     `json:"name"            required:"true"  column:"tag"   type:"string"`
	Namespace      string     `json:"namespace"       required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"cluster_name"    required:"true"  column:"tag"   type:"string"`
	Metric         string     `json:"metric"          required:"false" column:"tag"   type:"string"`
	Kind           string     `json:"kind"            required:"false" column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	ControllerKind string     `json:"controller_kind" required:"true"  column:"tag"   type:"string"`
	ModelId        string     `json:"model_id"        required:"true"  column:"field" type:"string"`
	PredictionId   string     `json:"prediction_id"   required:"true"  column:"field" type:"string"`
	Value          float64    `json:"value"           required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusControllerCPUUpperBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller" measurement:"controller" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time           *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	Name           string     `json:"name"            required:"true"  column:"tag"   type:"string"`
	Namespace      string     `json:"namespace"       required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"cluster_name"    required:"true"  column:"tag"   type:"string"`
	Metric         string     `json:"metric"          required:"false" column:"tag"   type:"string"`
	Kind           string     `json:"kind"            required:"false" column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	ControllerKind string     `json:"controller_kind" required:"true"  column:"tag"   type:"string"`
	ModelId        string     `json:"model_id"        required:"true"  column:"field" type:"string"`
	PredictionId   string     `json:"prediction_id"   required:"true"  column:"field" type:"string"`
	Value          float64    `json:"value"           required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusControllerCPULowerBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller" measurement:"controller" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time           *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	Name           string     `json:"name"            required:"true"  column:"tag"   type:"string"`
	Namespace      string     `json:"namespace"       required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"cluster_name"    required:"true"  column:"tag"   type:"string"`
	Metric         string     `json:"metric"          required:"false" column:"tag"   type:"string"`
	Kind           string     `json:"kind"            required:"false" column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	ControllerKind string     `json:"controller_kind" required:"true"  column:"tag"   type:"string"`
	ModelId        string     `json:"model_id"        required:"true"  column:"field" type:"string"`
	PredictionId   string     `json:"prediction_id"   required:"true"  column:"field" type:"string"`
	Value          float64    `json:"value"           required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusControllerMemory struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller" measurement:"controller" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time           *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	Name           string     `json:"name"            required:"true"  column:"tag"   type:"string"`
	Namespace      string     `json:"namespace"       required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"cluster_name"    required:"true"  column:"tag"   type:"string"`
	Metric         string     `json:"metric"          required:"false" column:"tag"   type:"string"`
	Kind           string     `json:"kind"            required:"false" column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	ControllerKind string     `json:"controller_kind" required:"true"  column:"tag"   type:"string"`
	ModelId        string     `json:"model_id"        required:"true"  column:"field" type:"string"`
	PredictionId   string     `json:"prediction_id"   required:"true"  column:"field" type:"string"`
	Value          float64    `json:"value"           required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusControllerMemoryUpperBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller" measurement:"controller" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time           *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	Name           string     `json:"name"            required:"true"  column:"tag"   type:"string"`
	Namespace      string     `json:"namespace"       required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"cluster_name"    required:"true"  column:"tag"   type:"string"`
	Metric         string     `json:"metric"          required:"false" column:"tag"   type:"string"`
	Kind           string     `json:"kind"            required:"false" column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	ControllerKind string     `json:"controller_kind" required:"true"  column:"tag"   type:"string"`
	ModelId        string     `json:"model_id"        required:"true"  column:"field" type:"string"`
	PredictionId   string     `json:"prediction_id"   required:"true"  column:"field" type:"string"`
	Value          float64    `json:"value"           required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusControllerMemoryLowerBound struct {
	DatahubEntity  `scope:"prediction" category:"cluster_status" type:"controller" measurement:"controller" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time           *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	Name           string     `json:"name"            required:"true"  column:"tag"   type:"string"`
	Namespace      string     `json:"namespace"       required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"cluster_name"    required:"true"  column:"tag"   type:"string"`
	Metric         string     `json:"metric"          required:"false" column:"tag"   type:"string"`
	Kind           string     `json:"kind"            required:"false" column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	ControllerKind string     `json:"controller_kind" required:"true"  column:"tag"   type:"string"`
	ModelId        string     `json:"model_id"        required:"true"  column:"field" type:"string"`
	PredictionId   string     `json:"prediction_id"   required:"true"  column:"field" type:"string"`
	Value          float64    `json:"value"           required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNamespaceCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNamespaceCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNamespaceCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNamespaceMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNamespaceMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNamespaceMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNodeCPU struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node" measurement:"node" metric:"cpu_usage_percentage" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	IsScheduled   string     `json:"is_scheduled"  required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNodeCPUUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node" measurement:"node" metric:"cpu_usage_percentage" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	IsScheduled   string     `json:"is_scheduled"  required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNodeCPULowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node" measurement:"node" metric:"cpu_usage_percentage" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	IsScheduled   string     `json:"is_scheduled"  required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNodeMemory struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node" measurement:"node" metric:"memory_usage_bytes" boundary:"raw" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	IsScheduled   string     `json:"is_scheduled"  required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNodeMemoryUpperBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node" measurement:"node" metric:"memory_usage_bytes" boundary:"upper_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	IsScheduled   string     `json:"is_scheduled"  required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}

type PredictionClusterStatusNodeMemoryLowerBound struct {
	DatahubEntity `scope:"prediction" category:"cluster_status" type:"node" measurement:"node" metric:"memory_usage_bytes" boundary:"lower_bound" quota:"undefined"`
	Time          *time.Time `json:"time"          required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"          required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name"  required:"true"  column:"tag"   type:"string"`
	Metric        string     `json:"metric"        required:"false" column:"tag"   type:"string"`
	Kind          string     `json:"kind"          required:"false" column:"tag"   type:"string"`
	Granularity   string     `json:"granularity"   required:"true"  column:"tag"   type:"string"`
	IsScheduled   string     `json:"is_scheduled"  required:"true"  column:"tag"   type:"string"`
	ModelId       string     `json:"model_id"      required:"true"  column:"field" type:"string"`
	PredictionId  string     `json:"prediction_id" required:"true"  column:"field" type:"string"`
	Value         float64    `json:"value"         required:"true"  column:"field" type:"float64"`
}
