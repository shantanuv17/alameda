package kafka

import (
	"strconv"

	corev1 "k8s.io/api/core/v1"
)

type ConsumerGroup struct {
	Name                   string
	ExporterNamespace      string
	ClusterName            string
	AlamedaScalerName      string
	AlamedaScalerNamespace string
	Policy                 string
	EnableExecution        bool
	ConsumeTopic           string
	ResourceMeta
	// Since value of min/max replicas are initialized from counts of topic's partition, define these fields in here rather than KubernetesMeta.
	MinReplicas int32
	MaxReplicas int32
}

type ResourceMeta struct {
	KubernetesMeta
	CustomName string
}

type KubernetesMeta struct {
	Namespace      string
	Name           string
	Kind           string
	ReadyReplicas  int32
	SpecReplicas   int32
	CPULimit       string
	CPURequest     string
	MemoryLimit    string
	MemoryRequest  string
	VolumesSize    string
	VolumesPVCSize string
}

func (m *KubernetesMeta) SetResourceRequirements(r corev1.ResourceRequirements) {
	m.CPULimit = strconv.FormatInt(r.Limits.Cpu().MilliValue(), 10)
	m.CPURequest = strconv.FormatInt(r.Requests.Cpu().MilliValue(), 10)
	m.MemoryLimit = strconv.FormatInt(r.Limits.Memory().Value(), 10)
	m.MemoryRequest = strconv.FormatInt(r.Requests.Memory().Value(), 10)
}

func (m KubernetesMeta) IsEmpty() bool {
	empty := KubernetesMeta{}
	if m == empty {
		return true
	}
	return false
}
