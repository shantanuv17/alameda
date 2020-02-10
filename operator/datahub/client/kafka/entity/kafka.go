package entity

import (
	"github.com/containers-ai/alameda/operator/pkg/kafka"
)

type KafkaTopic struct {
	Name              string `datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace string `datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	ClusterName       string `datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
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
	Name                 string `datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace    string `datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	ClusterName          string `datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName    string `datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	Policy               string `datahubcolumn:"policy" datahubdatatype:"DATATYPE_STRING"`
	EnableExecution      bool   `datahubcolumn:"enable_execution" datahubdatatype:"DATATYPE_BOOL"`
	ConsumeTopic         string `datahubcolumn:"topic_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceCustomName   string `datahubcolumn:"resource_custom_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SNamespace string `datahubcolumn:"resource_k8s_namespace" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SName      string `datahubcolumn:"resource_k8s_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SKind      string `datahubcolumn:"resource_k8s_kind" datahubdatatype:"DATATYPE_STRING"`
	ReadyReplicas        int32  `datahubcolumn:"resource_k8s_replicas" datahubdatatype:"DATATYPE_INT32"`
	SpecReplicas         int32  `datahubcolumn:"resource_k8s_spec_replicas" datahubdatatype:"DATATYPE_INT32"`
	MinReplicas          int32  `datahubcolumn:"resource_k8s_min_replicas" datahubdatatype:"DATATYPE_INT32"`
	MaxReplicas          int32  `datahubcolumn:"resource_k8s_max_replicas" datahubdatatype:"DATATYPE_INT32"`
}

func NewKafkaConsumerGroup(consumerGroup kafka.ConsumerGroup) KafkaConsumerGroup {
	return KafkaConsumerGroup{
		Name:                 consumerGroup.Name,
		ExporterNamespace:    consumerGroup.ExporterNamespace,
		ClusterName:          consumerGroup.ClusterName,
		AlamedaScalerName:    consumerGroup.AlamedaScalerName,
		Policy:               consumerGroup.Policy,
		EnableExecution:      consumerGroup.EnableExecution,
		ConsumeTopic:         consumerGroup.ConsumeTopic,
		ResourceCustomName:   consumerGroup.ResourceMeta.CustomName,
		ResourceK8SNamespace: consumerGroup.ResourceMeta.Namespace,
		ResourceK8SName:      consumerGroup.ResourceMeta.Name,
		ResourceK8SKind:      consumerGroup.ResourceMeta.Kind,
		ReadyReplicas:        consumerGroup.ResourceMeta.ReadyReplicas,
		SpecReplicas:         consumerGroup.ResourceMeta.SpecReplicas,
		MinReplicas:          consumerGroup.MinReplicas,
		MaxReplicas:          consumerGroup.MaxReplicas,
	}
}
