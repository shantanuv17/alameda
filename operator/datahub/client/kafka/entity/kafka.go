package entity

import (
	"github.com/containers-ai/alameda/operator/pkg/kafka"
)

type KafkaTopic struct {
	Name              string `datahubcolumn:"name"`
	ExporterNamespace string `datahubcolumn:"namespace"`
	ClusterName       string `datahubcolumn:"cluster_name"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name"`
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
	Name                 string `datahubcolumn:"name"`
	ExporterNamespace    string `datahubcolumn:"namespace"`
	ClusterName          string `datahubcolumn:"cluster_name"`
	AlamedaScalerName    string `datahubcolumn:"alameda_scaler_name"`
	Policy               string `datahubcolumn:"policy"`
	EnableExecution      bool   `datahubcolumn:"enable_execution"`
	ConsumeTopic         string `datahubcolumn:"topic_name"`
	ResourceCustomName   string `datahubcolumn:"resource_custom_name"`
	ResourceK8SNamespace string `datahubcolumn:"resource_k8s_namespace"`
	ResourceK8SName      string `datahubcolumn:"resource_k8s_name"`
	ResourceK8SKind      string `datahubcolumn:"resource_k8s_kind"`
	ReadyReplicas        int32  `datahubcolumn:"resource_k8s_replicas"`
	SpecReplicas         int32  `datahubcolumn:"resource_k8s_spec_replicas"`
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
	}
}
