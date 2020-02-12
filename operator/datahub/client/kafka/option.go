package kafka

type ListTopicsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace      string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}

type ListConsumerGroupsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace      string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}

type DeleteTopicsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace      string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}

type DeleteConsumerGroupsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace      string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}
