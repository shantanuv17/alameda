package entities

import (
	"time"
)

type TargetClusterStatusCluster struct {
	DatahubEntity          `scope:"target" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                   *time.Time `json:"time"                     required:"false" column:"tag"   type:"time"`
	Name                   string     `json:"name"                     required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName      string     `json:"alameda_scaler_name"      required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace string     `json:"alameda_scaler_namespace" required:"true"  column:"tag"   type:"string"`
	RawSpec                string     `json:"raw_spec"                 required:"true"  column:"field" type:"string"`
}

type TargetClusterStatusController struct {
	DatahubEntity            `scope:"target" category:"cluster_status" type:"controller" measurement:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time `json:"time"                        required:"false" column:"tag"   type:"time"`
	Name                     string     `json:"name"                        required:"true"  column:"tag"   type:"string"`
	Namespace                string     `json:"namespace"                   required:"true"  column:"tag"   type:"string"`
	ClusterName              string     `json:"cluster_name"                required:"true"  column:"tag"   type:"string"`
	Kind                     string     `json:"kind"                        required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName        string     `json:"alameda_scaler_name"         required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace   string     `json:"alameda_scaler_namespace"    required:"true"  column:"tag"   type:"string"`
	AlamedaScalerScalingTool string     `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"   type:"string"`
	MinReplicas              int32      `json:"resource_k8s_min_replicas"   required:"false" column:"field" type:"int32"`
	MaxReplicas              int32      `json:"resource_k8s_max_replicas"   required:"false" column:"field" type:"int32"`
	Policy                   string     `json:"policy"                      required:"false" column:"field" type:"string"`
	EnableExecution          bool       `json:"enable_execution"            required:"false" column:"field" type:"bool"`
	RawSpec                  string     `json:"raw_spec"                    required:"false" column:"field" type:"string"`
}

type TargetKafkaTopic struct {
	DatahubEntity          `scope:"target" category:"kafka" type:"topic" measurement:"kafka_topic" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                   *time.Time `json:"time"                     required:"false" column:"tag"   type:"time"`
	Name                   string     `json:"name"                     required:"true"  column:"tag"   type:"string"`
	ClusterName            string     `json:"cluster_name"             required:"true"  column:"tag"   type:"string"`
	ExporterNamespace      string     `json:"exporter_namespace"       required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName      string     `json:"alameda_scaler_name"      required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace string     `json:"alameda_scaler_namespace" required:"true"  column:"tag"   type:"string"`
	RawSpec                string     `json:"raw_spec"                 required:"true"  column:"field" type:"string"`
}

type TargetKafkaConsumerGroup struct {
	DatahubEntity            `scope:"target" category:"kafka" type:"consumer_group" measurement:"kafka_consumer_group" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time `json:"time"                        required:"false" column:"tag"   type:"time"`
	Name                     string     `json:"name"                        required:"true"  column:"tag"   type:"string"`
	ClusterName              string     `json:"cluster_name"                required:"true"  column:"tag"   type:"string"`
	ExporterNamespace        string     `json:"exporter_namespace"          required:"true"  column:"tag"   type:"string"`
	TopicName                string     `json:"topic_name"                  required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName        string     `json:"alameda_scaler_name"         required:"true"  column:"tag"   type:"string"`
	AlamedaScalerNamespace   string     `json:"alameda_scaler_namespace"    required:"true"  column:"tag"   type:"string"`
	AlamedaScalerScalingTool string     `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"   type:"string"`
	ResourceK8sName          string     `json:"resource_k8s_name"           required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace     string     `json:"resource_k8s_namespace"      required:"true"  column:"tag"   type:"string"`
	ResourceK8sKind          string     `json:"resource_k8s_kind"           required:"true"  column:"tag"   type:"string"`
	ResourceK8sMinReplicas   int32      `json:"resource_k8s_min_replicas"   required:"false" column:"field" type:"int32"`
	ResourceK8sMaxReplicas   int32      `json:"resource_k8s_max_replicas"   required:"false" column:"field" type:"int32"`
	Policy                   string     `json:"policy"                      required:"false" column:"field" type:"string"`
	GroupId                  string     `json:"group_id"                    required:"true"  column:"field" type:"string"`
	EnableExecution          bool       `json:"enable_execution"            required:"false" column:"field" type:"bool"`
	RawSpec                  string     `json:"raw_spec"                    required:"false" column:"field" type:"string"`
}
