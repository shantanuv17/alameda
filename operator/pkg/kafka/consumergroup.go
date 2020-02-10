package kafka

type ConsumerGroup struct {
	Name              string
	ExporterNamespace string
	ClusterName       string
	AlamedaScalerName string
	Policy            string
	EnableExecution   bool
	ConsumeTopic      string
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
	Namespace     string
	Name          string
	Kind          string
	ReadyReplicas int32
	SpecReplicas  int32
}

func (m KubernetesMeta) IsEmpty() bool {
	empty := KubernetesMeta{}
	if m == empty {
		return true
	}
	return false
}
