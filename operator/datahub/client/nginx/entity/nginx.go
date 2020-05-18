package entity

import "github.com/containers-ai/alameda/operator/pkg/nginx"

type Nginx struct {
	Namespace                   string `datahubcolumntype:"tag" datahubcolumn:"namespace" datahubdatatype:"DATATYPE_STRING"`
	ClusterName                 string `datahubcolumntype:"tag" datahubcolumn:"cluster_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerName           string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_name" datahubdatatype:"DATATYPE_STRING"`
	AlamedaScalerNamespace      string `datahubcolumntype:"tag" datahubcolumn:"alameda_scaler_namespace" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SServiceName      string `datahubcolumntype:"tag" datahubcolumn:"resource_k8s_service_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SServiceNamespace string `datahubcolumntype:"tag" datahubcolumn:"resource_k8s_service_namespace" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SKind             string `datahubcolumntype:"tag" datahubcolumn:"resource_k8s_kind" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SRouteName        string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_route_name" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SRouteNamespace   string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_route_namespace" datahubdatatype:"DATATYPE_STRING"`
	ExporterPods                string `datahubcolumntype:"field" datahubcolumn:"exporter_pods" datahubdatatype:"DATATYPE_STRING"`
	ExporterNamespace           string `datahubcolumntype:"field" datahubcolumn:"exporter_namespace" datahubdatatype:"DATATYPE_STRING"`
	Policy                      string `datahubcolumntype:"field" datahubcolumn:"policy" datahubdatatype:"DATATYPE_STRING"`
	EnableExecution             bool   `datahubcolumntype:"field" datahubcolumn:"enable_execution" datahubdatatype:"DATATYPE_BOOL"`
	ResourceK8SNamespace        string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_namespace" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SName             string `datahubcolumntype:"field" datahubcolumn:"resource_k8s_name" datahubdatatype:"DATATYPE_STRING"`
	ReadyReplicas               int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_replicas" datahubdatatype:"DATATYPE_INT32"`
	SpecReplicas                int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_spec_replicas" datahubdatatype:"DATATYPE_INT32"`
	MinReplicas                 int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_min_replicas" datahubdatatype:"DATATYPE_INT32"`
	MaxReplicas                 int32  `datahubcolumntype:"field" datahubcolumn:"resource_k8s_max_replicas" datahubdatatype:"DATATYPE_INT32"`
	ResourceK8SCPULimit         string `datahubcolumntype:"field" datahubcolumn:"resource_cpu_limit" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SCPURequest       string `datahubcolumntype:"field" datahubcolumn:"resource_cpu_request" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SMemoryLimit      string `datahubcolumntype:"field" datahubcolumn:"resource_memory_limit" datahubdatatype:"DATATYPE_STRING"`
	ResourceK8SMemoryRequest    string `datahubcolumntype:"field" datahubcolumn:"resource_memory_request" datahubdatatype:"DATATYPE_STRING"`
	ReplicaMarginPercentage     int32  `datahubcolumntype:"field" datahubcolumn:"replica_margin_percentage" datahubdatatype:"DATATYPE_INT32"`
}

func NewNginx(nginx nginx.Nginx) Nginx {
	return Nginx{
		Namespace:                   nginx.ExporterNamespace,
		ClusterName:                 nginx.ClusterName,
		AlamedaScalerName:           nginx.AlamedaScalerName,
		AlamedaScalerNamespace:      nginx.AlamedaScalerNamespace,
		ExporterNamespace:           nginx.ExporterNamespace,
		ExporterPods:                nginx.ExporterPods,
		ResourceK8SRouteNamespace:   nginx.RouteNamespace,
		ResourceK8SRouteName:        nginx.RouteName,
		Policy:                      nginx.Policy,
		EnableExecution:             nginx.EnableExecution,
		ResourceK8SServiceName:      nginx.ResourceMeta.KubernetesMeta.ServiceName,
		ResourceK8SServiceNamespace: nginx.ResourceMeta.KubernetesMeta.ServiceNamespace,
		ResourceK8SNamespace:        nginx.ResourceMeta.Namespace,
		ResourceK8SName:             nginx.ResourceMeta.Name,
		ResourceK8SKind:             nginx.ResourceMeta.Kind,
		ReadyReplicas:               nginx.ResourceMeta.ReadyReplicas,
		SpecReplicas:                nginx.ResourceMeta.SpecReplicas,
		MinReplicas:                 nginx.MinReplicas,
		MaxReplicas:                 nginx.MaxReplicas,
		ResourceK8SCPULimit:         nginx.ResourceMeta.KubernetesMeta.CPULimit,
		ResourceK8SCPURequest:       nginx.ResourceMeta.KubernetesMeta.CPURequest,
		ResourceK8SMemoryLimit:      nginx.ResourceMeta.KubernetesMeta.MemoryLimit,
		ResourceK8SMemoryRequest:    nginx.ResourceMeta.KubernetesMeta.MemoryRequest,
		ReplicaMarginPercentage:     nginx.ReplicaMarginPercentage,
	}
}
