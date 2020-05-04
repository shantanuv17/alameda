package machineset

type ListMachineSetsOption struct {
	ClusterName      string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	Name             string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	Namespace        string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	MachineGroupName string `datahubcolumntype:"tag" datahubcolumn:"machinegroup_name" datahubdatatype:"DATATYPE_STRING"`
}

type DeleteMachineSetsOption struct {
	ClusterName      string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	Name             string `datahubcolumntype:"tag" datahubcolumn:"name" datahubdatatype:"DATATYPE_STRING"`
	Namespace        string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	MachineGroupName string `datahubcolumntype:"tag" datahubcolumn:"machinegroup_name" datahubdatatype:"DATATYPE_STRING"`
}
