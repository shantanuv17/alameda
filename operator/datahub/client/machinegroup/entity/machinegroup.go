package entity

import (
	"github.com/containers-ai/alameda/operator/pkg/machinegroup"
)

type MachineGroup struct {
	Namespace              string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	Name                   string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
	Dummy                  string `datahubcolumntype:"field" datahubcolumn:"dummy" datahubdatatype:"DATATYPE_STRING"`
}

func NewMachineGroup(machineGroup machinegroup.MachineGroup) MachineGroup {
	return MachineGroup{
		Name:                   machineGroup.Name,
		Namespace:              machineGroup.Namespace,
		AlamedaScalerNamespace: machineGroup.AlamedaScalerNamespace,
		AlamedaScalerName:      machineGroup.AlamedaScalerName,
		ClusterName:            machineGroup.ClusterName,
	}
}
