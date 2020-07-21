package entities

import (
	"time"
)

type ApplicationKafkaTopic struct {
	DatahubEntity          `scope:"application" category:"kafka" type:"topic" measurement:"kafka_topic" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                   *time.Time `json:"time"                     required:"false" column:"tag"   type:"time"`
	Name                   string     `json:"name"                     required:"true"  column:"tag"   type:"string"`
	Namespace              string     `json:"namespace"                required:"true"  column:"tag"   type:"string"`
	ClusterName            string     `json:"cluster_name"             required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName      string     `json:"alameda_scaler_name"      required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace string     `json:"alameda_scaler_namespace" required:"true"  column:"tag"   type:"string"`
	Dummy                  string     `json:"dummy"                    required:"true"  column:"field" type:"string"`
}

type ApplicationKafkaConsumerGroup struct {
	DatahubEntity            `scope:"application" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time `json:"time"                        required:"false" column:"tag"   type:"time"`
	Name                     string     `json:"name"                        required:"true"  column:"tag"   type:"string"`
	Namespace                string     `json:"namespace"                   required:"true"  column:"tag"   type:"string"`
	ClusterName              string     `json:"cluster_name"                required:"true"  column:"tag"   type:"string"`
	TopicName                string     `json:"topic_name"                  required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName        string     `json:"alameda_scaler_name"         required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace   string     `json:"alameda_scaler_namespace"    required:"true"  column:"tag"   type:"string"`
	AlamedaScalerScalingTool string     `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName          string     `json:"resource_k8s_name"           required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace     string     `json:"resource_k8s_namespace"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind          string     `json:"resource_k8s_kind"           required:"true"  column:"tag"   type:"string"`
	ResourceK8sSpecReplicas  int32      `json:"resource_k8s_spec_replicas"  required:"true"  column:"field" type:"int32"`
	ResourceK8sReplicas      int32      `json:"resource_k8s_replicas"       required:"false" column:"field" type:"int32"`
	ResourceK8sMinReplicas   int32      `json:"resource_k8s_min_replicas"   required:"false" column:"field" type:"int32"`
	ResourceK8sMaxReplicas   int32      `json:"resource_k8s_max_replicas"   required:"false" column:"field" type:"int32"`
	ResourceCustomName       string     `json:"resource_custom_name"        required:"false" column:"field" type:"string"`
	Policy                   string     `json:"policy"                      required:"false" column:"field" type:"string"`
	EnableExecution          bool       `json:"enable_execution"            required:"false" column:"field" type:"bool"`
	ResourceCPULimit         string     `json:"resource_cpu_limit"          required:"false" column:"field" type:"string"`
	ResourceCPURequest       string     `json:"resource_cpu_request"        required:"false" column:"field" type:"string"`
	ResourceMemoryLimit      string     `json:"resource_memory_limit"       required:"false" column:"field" type:"string"`
	ResourceMemoryRequest    string     `json:"resource_memory_request"     required:"false" column:"field" type:"string"`
	VolumesSize              string     `json:"volumes_size"                required:"false" column:"field" type:"string"`
	VolumesPvcSize           string     `json:"volumes_pvc_size"            required:"false" column:"field" type:"string"`
}

type ApplicationNginx struct {
	DatahubEntity               `scope:"application" category:"nginx" type:"nginx" measurement:"nginx" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                        *time.Time `json:"time"                           required:"false" column:"tag"   type:"time"`
	Namespace                   string     `json:"namespace"                      required:"true"  column:"tag"   type:"string"`
	ClusterName                 string     `json:"cluster_name"                   required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName           string     `json:"alameda_scaler_name"            required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace      string     `json:"alameda_scaler_namespace"       required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceName      string     `json:"resource_k8s_service_name"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sServiceNamespace string     `json:"resource_k8s_service_namespace" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName             string     `json:"resource_k8s_name"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace        string     `json:"resource_k8s_namespace"         required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind             string     `json:"resource_k8s_kind"              required:"true"  column:"tag"   type:"string"`
	ResourceK8sSpecReplicas     int32      `json:"resource_k8s_spec_replicas"     required:"true"  column:"field" type:"int32"`
	ResourceK8sReplicas         int32      `json:"resource_k8s_replicas"          required:"false" column:"field" type:"int32"`
	ResourceK8sMinReplicas      int32      `json:"resource_k8s_min_replicas"      required:"false" column:"field" type:"int32"`
	ResourceK8sMaxReplicas      int32      `json:"resource_k8s_max_replicas"      required:"false" column:"field" type:"int32"`
	ResourceK8sRouteName        string     `json:"resource_k8s_route_name"        required:"false" column:"field" type:"string"`
	ResourceK8sRouteNamespace   string     `json:"resource_k8s_route_namespace"   required:"false" column:"field" type:"string"`
	ExporterPods                string     `json:"exporter_pods"                  required:"false" column:"field" type:"string"`
	ExporterNamespace           string     `json:"exporter_namespace"             required:"false" column:"field" type:"string"`
	Policy                      string     `json:"policy"                         required:"false" column:"field" type:"string"`
	EnableExecution             bool       `json:"enable_execution"               required:"false" column:"field" type:"bool"`
	ResourceCPULimit            string     `json:"resource_cpu_limit"             required:"false" column:"field" type:"string"`
	ResourceCPURequest          string     `json:"resource_cpu_request"           required:"false" column:"field" type:"string"`
	ResourceMemoryLimit         string     `json:"resource_memory_limit"          required:"false" column:"field" type:"string"`
	ResourceMemoryRequest       string     `json:"resource_memory_request"        required:"false" column:"field" type:"string"`
	ReplicaMarginPercentage     int32      `json:"replica_margin_percentage"      required:"false" column:"field" type:"int32"`
	HttpResponseTime            int64      `json:"http_response_time"             required:"false" column:"field" type:"int64"`
}
