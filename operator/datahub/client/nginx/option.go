package nginx

type ListNginxsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}

type DeleteNginxsOption struct {
	ClusterName            string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
}
