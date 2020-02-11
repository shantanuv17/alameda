package entity

import (
	"github.com/containers-ai/alameda/operator/pkg/kafka"
)

type KafkaTopic struct {
	Name              string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	ClusterName       string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName string `datahubcolumntype:"field" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
}

func NewKafkaTopic(topic kafka.Topic) KafkaTopic {
	return KafkaTopic{
		Name:              topic.Name,
		ExporterNamespace: topic.ExporterNamespace,
		ClusterName:       topic.ClusterName,
		AlamedaScalerName: topic.AlamedaScalerName,
	}
}

type KafkaConsumerGroup struct {
	Name                      string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace         string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	ClusterName               string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName         string `datahubcolumntype:"field" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	Policy                    string `datahubcolumntype:"field" datahubcolumn:"policy" datahubdatatype:"DATATYPE_STRING"`
	EnableExecution           bool   `datahubcolumntype:"field" datahubcolumn:"enable_execution" datahubdatatype:"DATATYPE_BOOL"`
	ConsumeTopic              string `datahubcolumntype:"tag" datahubcolumn:"topic_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceCustomName        string `datahubcolumntype:"field" datahubcolumn:"resource_custom_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SNamespace      string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_namespace" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SName           string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SKind           string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_kind" datahubdatatype:"DATATYPE_STRING"`
	ReadyReplicas             int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_replicas" datahubdatatype:"DATATYPE_INT32"`
	SpecReplicas              int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_spec_replicas" datahubdatatype:"DATATYPE_INT32"`
	MinReplicas               int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_min_replicas" datahubdatatype:"DATATYPE_INT32"`
	MaxReplicas               int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_max_replicas" datahubdatatype:"DATATYPE_INT32"`
	ResourceK8SCPULimit       string `datahubcolumntype:"field" datahubcolumn:"resource_cpu_limit" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SCPURequest     string `datahubcolumntype:"field" datahubcolumn:"resource_cpu_request" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SMemoryLimit    string `datahubcolumntype:"field" datahubcolumn:"resource_memory_limit" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SMemoryRequest  string `datahubcolumntype:"field" datahubcolumn:"resource_memory_request" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SVolumesSize    string `datahubcolumntype:"field" datahubcolumn:"volumes_size" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SVolumesPVCSize string `datahubcolumntype:"field" datahubcolumn:"volumes_pvc_size" datahubdatatype:"DATATYPE_STRING"`
}

func NewKafkaConsumerGroup(consumerGroup kafka.ConsumerGroup) KafkaConsumerGroup {
	return KafkaConsumerGroup{
		Name:                      consumerGroup.Name,
		ExporterNamespace:         consumerGroup.ExporterNamespace,
		ClusterName:               consumerGroup.ClusterName,
		AlamedaScalerName:         consumerGroup.AlamedaScalerName,
		Policy:                    consumerGroup.Policy,
		EnableExecution:           consumerGroup.EnableExecution,
		ConsumeTopic:              consumerGroup.ConsumeTopic,
		ResourceCustomName:        consumerGroup.ResourceMeta.CustomName,
		ResourceK8SNamespace:      consumerGroup.ResourceMeta.Namespace,
		ResourceK8SName:           consumerGroup.ResourceMeta.Name,
		ResourceK8SKind:           consumerGroup.ResourceMeta.Kind,
		ReadyReplicas:             consumerGroup.ResourceMeta.ReadyReplicas,
		SpecReplicas:              consumerGroup.ResourceMeta.SpecReplicas,
		MinReplicas:               consumerGroup.MinReplicas,
		MaxReplicas:               consumerGroup.MaxReplicas,
		ResourceK8SCPULimit:       consumerGroup.ResourceMeta.KubernetesMeta.CPULimit,
		ResourceK8SCPURequest:     consumerGroup.ResourceMeta.KubernetesMeta.CPURequest,
		ResourceK8SMemoryLimit:    consumerGroup.ResourceMeta.KubernetesMeta.MemoryLimit,
		ResourceK8SMemoryRequest:  consumerGroup.ResourceMeta.KubernetesMeta.MemoryRequest,
		ResourceK8SVolumesSize:    consumerGroup.ResourceMeta.KubernetesMeta.VolumesSize,
		ResourceK8SVolumesPVCSize: consumerGroup.ResourceMeta.KubernetesMeta.VolumesPVCSize,
	}
}
