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
