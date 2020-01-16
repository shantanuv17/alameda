package kafka

type Topic struct {
	Name              string `datahubcolumn:"name"`
	ExporterNamespace string `datahubcolumn:"namespace"`
	ClusterName       string `datahubcolumn:"cluster_name"`
	AlamedaScalerName string `datahubcolumn:"alameda_scaler_name"`
}
