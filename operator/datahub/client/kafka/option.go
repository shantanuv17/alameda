package kafka

type ListTopicsOption struct {
	ClusterName       string `datahubcolumn:"cluster_name"`
	ExporterNamespace string `datahubcolumn:"namespace"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name"`
}

type ListConsumerGroupsOption struct {
	ClusterName       string `datahubcolumn:"cluster_name"`
	ExporterNamespace string `datahubcolumn:"namespace"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name"`
}
