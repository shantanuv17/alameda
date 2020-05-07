package machinegroup

type MachineGroup struct {
	ClusterName string
	ResourceMeta
	AlamedaScalerNamespace        string
	AlamedaScalerName             string
	CPUMetricUtilizationTarget    int32
	CPUMetricScaleupGap           int32
	CPUMetricScaledownGap         int32
	MemoryMetricUtilizationTarget int32
	MemoryMetricScaleupGap        int32
	MemoryMetricScaledownGap      int32
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
