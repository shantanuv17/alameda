package entity

import (
	"github.com/containers-ai/alameda/operator/pkg/machineset"
)

type MachineSet struct {
	Namespace             string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	Name                  string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ClusterName           string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	MachineGroupName      string `datahubcolumntype:"tag" datahubcolumn:"machinegroup_name" datahubdatatype:"DATATYPE_STRING"`
	MachineGroupNamespace string `datahubcolumntype:"tag" datahubcolumn:"machinegroup_namespace" datahubdatatype:"DATATYPE_STRING"`
	EnableExecution       bool   `datahubcolumntype:"field" datahubcolumn:"enable_execution" datahubdatatype:"DATATYPE_BOOL"`
	ReadyReplicas         int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_replicas" datahubdatatype:"DATATYPE_INT32"`
	SpecReplicas          int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_spec_replicas" datahubdatatype:"DATATYPE_INT32"`
	MinReplicas           int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_min_replicas" datahubdatatype:"DATATYPE_INT32"`
	MaxReplicas           int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_max_replicas" datahubdatatype:"DATATYPE_INT32"`
}

func NewMachineSet(machineSet machineset.MachineSet) MachineSet {
	return MachineSet{
		Name:                  machineSet.Name,
		Namespace:             machineSet.Namespace,
		ClusterName:           machineSet.ClusterName,
		MachineGroupName:      machineSet.MachineGroupName,
		MachineGroupNamespace: machineSet.MachineGroupNamespace,
		EnableExecution:       machineSet.EnableExecution,
		ReadyReplicas:         machineSet.ReadyReplicas,
		SpecReplicas:          machineSet.SpecReplicas,
		MinReplicas:           machineSet.MinReplicas,
		MaxReplicas:           machineSet.MaxReplicas,
	}
}
