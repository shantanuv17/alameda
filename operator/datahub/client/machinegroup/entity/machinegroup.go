package entity

import (
	"github.com/containers-ai/alameda/operator/pkg/machinegroup"
)

type MachineGroup struct {
	Namespace                     string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	Name                          string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	ClusterName                   string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName             string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace        string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
	CPUMetricUtilizationTarget    int32  `datahubcolumntype:"field" datahubcolumn:"cpu_metric_utilization_target" datahubdatatype:"DATATYPE_INT32"`
	CPUMetricScaleupGap           int32  `datahubcolumntype:"field" datahubcolumn:"cpu_metric_scaleup_gap" datahubdatatype:"DATATYPE_INT32"`
	CPUMetricScaledownGap         int32  `datahubcolumntype:"field" datahubcolumn:"cpu_metric_scaledown_gap" datahubdatatype:"DATATYPE_INT32"`
	MemoryMetricUtilizationTarget int32  `datahubcolumntype:"field" datahubcolumn:"memory_metric_utilization_target" datahubdatatype:"DATATYPE_INT32"`
	MemoryMetricScaleupGap        int32  `datahubcolumntype:"field" datahubcolumn:"memory_metric_scaleup_gap" datahubdatatype:"DATATYPE_INT32"`
	MemoryMetricScaledownGap      int32  `datahubcolumntype:"field" datahubcolumn:"memory_metric_scaledown_gap" datahubdatatype:"DATATYPE_INT32"`
}

func NewMachineGroup(machineGroup machinegroup.MachineGroup) MachineGroup {
	return MachineGroup{
		Name:                          machineGroup.Name,
		Namespace:                     machineGroup.Namespace,
		AlamedaScalerNamespace:        machineGroup.AlamedaScalerNamespace,
		AlamedaScalerName:             machineGroup.AlamedaScalerName,
		ClusterName:                   machineGroup.ClusterName,
		CPUMetricUtilizationTarget:    machineGroup.CPUMetricUtilizationTarget,
		CPUMetricScaleupGap:           machineGroup.CPUMetricScaleupGap,
		CPUMetricScaledownGap:         machineGroup.CPUMetricScaledownGap,
		MemoryMetricUtilizationTarget: machineGroup.MemoryMetricUtilizationTarget,
		MemoryMetricScaleupGap:        machineGroup.MemoryMetricScaleupGap,
		MemoryMetricScaledownGap:      machineGroup.MemoryMetricScaledownGap,
	}
}
