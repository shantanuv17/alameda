package machinegroup

type MachineGroup struct {
	ClusterName string
	ResourceMeta
	AlamedaScalerNamespace string
	AlamedaScalerName      string
}

type ResourceMeta struct {
	KubernetesMeta
}

type KubernetesMeta struct {
	Namespace string
	Name      string
}

func (m KubernetesMeta) IsEmpty() bool {
	empty := KubernetesMeta{}
	if m == empty {
		return true
	}
	return false
}
