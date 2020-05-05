package machineset

type MachineSet struct {
	ClusterName           string
	MachineGroupName      string
	MachineGroupNamespace string
	ResourceMeta
	MinReplicas int32
	MaxReplicas int32
}

type ResourceMeta struct {
	KubernetesMeta
}

type KubernetesMeta struct {
	Namespace     string
	Name          string
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
