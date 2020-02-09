package kafka

type ListTopicsOption struct {
	ClusterName       string `datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace string `datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
}

type ListConsumerGroupsOption struct {
	ClusterName       string `datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace string `datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
}
