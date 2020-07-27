package entities

import (
	"time"
)

type RecommendationKafkaConsumerGroup struct {
	DatahubEntity         `scope:"recommendation" category:"kafka" type:"consumer_group"`
	Metadata              *Metadata  `measurement:"kafka_consumer_group" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time `json:"time"                    required:"false" column:"tag"   type:"time"`
	Name                  string     `json:"name"                    required:"true"  column:"tag"   type:"string"`
	Namespace             string     `json:"namespace"               required:"true"  column:"tag"   type:"string"`
	ClusterName           string     `json:"cluster_name"            required:"true"  column:"tag"   type:"string"`
	ResourceK8sName       string     `json:"resource_k8s_name"       required:"true"  column:"tag"   type:"string"`
	ResourceK8sNamespace  string     `json:"resource_k8s_namespace"  required:"true"  column:"tag"   type:"string"`
	Kind                  string     `json:"kind"                    required:"true"  column:"tag"   type:"string"`
	CreateTime            int64      `json:"create_time"             required:"true"  column:"field" type:"int64"`
	CurrentReplicas       int32      `json:"current_replicas"        required:"true"  column:"field" type:"int32"`
	DesiredReplicas       int32      `json:"desired_replicas"        required:"true"  column:"field" type:"int32"`
	PerConsumerCapability float64    `json:"per_consumer_capability" required:"false" column:"field" type:"float64"`
}

type RecommendationNginx struct {
	DatahubEntity    `scope:"recommendation" category:"nginx" type:"nginx" measurement:"nginx" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time             *time.Time `json:"time"               required:"false" column:"tag"   type:"time"`
	Name             string     `json:"name"               required:"true"  column:"tag"   type:"string"`
	Namespace        string     `json:"namespace"          required:"true"  column:"tag"   type:"string"`
	ClusterName      string     `json:"cluster_name"       required:"true"  column:"tag"   type:"string"`
	Kind             string     `json:"kind"               required:"true"  column:"tag"   type:"string"`
	CreateTime       int64      `json:"create_time"        required:"true"  column:"field" type:"int64"`
	CurrentReplicas  int32      `json:"current_replicas"   required:"true"  column:"field" type:"int32"`
	DesiredReplicas  int32      `json:"desired_replicas"   required:"true"  column:"field" type:"int32"`
	ReplicaCapacity  float64    `json:"replica_capacity"   required:"true"  column:"field" type:"float64"`
	CurrentNumerator float64    `json:"current_numerator"  required:"true"  column:"field" type:"float64"`
	HttpResponseTime int64      `json:"http_response_time" required:"false" column:"field" type:"int64"`
	Alpha            float64    `json:"alpha"              required:"false" column:"field" type:"float64"`
	Mape0            float64    `json:"mape0"              required:"true"  column:"field" type:"float64"`
	Mape1            float64    `json:"mape1"              required:"true"  column:"field" type:"float64"`
	Mape2            float64    `json:"mape2"              required:"true"  column:"field" type:"float64"`
	Reserve          string     `json:"reserve"            required:"false" column:"field" type:"string"`
}

type RecommendationClusterAutoscalerMachinegroup struct {
	DatahubEntity       `scope:"recommendation" category:"cluster_autoscaler" type:"machinegroup" measurement:"machinegroup" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                *time.Time `json:"time"                 required:"false" column:"tag"   type:"time"`
	Name                string     `json:"name"                 required:"true"  column:"tag"   type:"string"`
	Namespace           string     `json:"namespace"            required:"true"  column:"tag"   type:"string"`
	ClusterName         string     `json:"cluster_name"         required:"true"  column:"tag"   type:"string"`
	MachinesetName      string     `json:"machineset_name"      required:"true"  column:"field" type:"string"`
	MachinesetNamespace string     `json:"machineset_namespace" required:"true"  column:"field" type:"string"`
	CreateTime          int64      `json:"create_time"          required:"true"  column:"field" type:"int64"`
	CurrentReplicas     int32      `json:"current_replicas"     required:"true"  column:"field" type:"int32"`
	DesiredReplicas     int32      `json:"desired_replicas"     required:"true"  column:"field" type:"int32"`
}

type RecommendationClusterStatusApplication struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"application"`
	Metadata              *Metadata  `measurement:"application" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time `json:"time"                 required:"false" column:"tag"   type:"time"`
	Name                  string     `json:"name"                 required:"true"  column:"tag"   type:"string"`
	Namespace             string     `json:"namespace"            required:"true"  column:"tag"   type:"string"`
	ClusterName           string     `json:"cluster_name"         required:"true"  column:"tag"   type:"string"`
	Type                  string     `json:"type"                 required:"true"  column:"tag"   type:"string"`
	Kind                  string     `json:"kind"                 required:"true"  column:"field" type:"string"`
	CurrentReplicas       int32      `json:"current_replicas"     required:"true"  column:"field" type:"int32"`
	DesiredReplicas       int32      `json:"desired_replicas"     required:"true"  column:"field" type:"int32"`
	CreateTime            int64      `json:"create_time"          required:"true"  column:"field" type:"int64"`
	CurrentCPURequests    float64    `json:"current_cpu_requests" required:"false" column:"field" type:"float64"`
	CurrentMemoryRequests float64    `json:"current_mem_requests" required:"false" column:"field" type:"float64"`
	CurrentCPULimits      float64    `json:"current_cpu_limits"   required:"false" column:"field" type:"float64"`
	CurrentMemoryLimits   float64    `json:"current_mem_limits"   required:"false" column:"field" type:"float64"`
	DesiredCPULimits      float64    `json:"desired_cpu_limits"   required:"false" column:"field" type:"float64"`
	DesiredMemoryLimits   float64    `json:"desired_mem_limits"   required:"false" column:"field" type:"float64"`
	TotalCost             float64    `json:"total_cost"           required:"false" column:"field" type:"float64"`
}

type RecommendationClusterStatusCluster struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"cluster"`
	Metadata              *Metadata  `measurement:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time `json:"time"                 required:"false" column:"tag"   type:"time"`
	Name                  string     `json:"name"                 required:"true"  column:"tag"   type:"string"`
	Type                  string     `json:"type"                 required:"true"  column:"tag"   type:"string"`
	Kind                  string     `json:"kind"                 required:"true"  column:"field" type:"string"`
	CurrentReplicas       int32      `json:"current_replicas"     required:"true"  column:"field" type:"int32"`
	DesiredReplicas       int32      `json:"desired_replicas"     required:"true"  column:"field" type:"int32"`
	CreateTime            int64      `json:"create_time"          required:"true"  column:"field" type:"int64"`
	CurrentCPURequests    float64    `json:"current_cpu_requests" required:"false" column:"field" type:"float64"`
	CurrentMemoryRequests float64    `json:"current_mem_requests" required:"false" column:"field" type:"float64"`
	CurrentCPULimits      float64    `json:"current_cpu_limits"   required:"false" column:"field" type:"float64"`
	CurrentMemoryLimits   float64    `json:"current_mem_limits"   required:"false" column:"field" type:"float64"`
	DesiredCPULimits      float64    `json:"desired_cpu_limits"   required:"false" column:"field" type:"float64"`
	DesiredMemoryLimits   float64    `json:"desired_mem_limits"   required:"false" column:"field" type:"float64"`
	TotalCost             float64    `json:"total_cost"           required:"false" column:"field" type:"float64"`
}

type RecommendationClusterStatusContainerLimit struct {
	DatahubEntity              `scope:"recommendation" category:"cluster_status" type:"container"`
	Metadata                   *Metadata  `measurement:"container" metric:"undefined" boundary:"undefined" quota:"limit" ts:"true"`
	Time                       *time.Time `json:"time"                          required:"false" column:"tag"   type:"time"`
	Name                       string     `json:"name"                          required:"true"  column:"tag"   type:"string"`
	Namespace                  string     `json:"namespace"                     required:"true"  column:"tag"   type:"string"`
	ClusterName                string     `json:"cluster_name"                  required:"true"  column:"tag"   type:"string"`
	PodName                    string     `json:"pod_name"                      required:"true"  column:"tag"   type:"string"`
	Granularity                string     `json:"granularity"                   required:"true"  column:"tag"   type:"string"`
	TopControllerName          string     `json:"top_controller_name"           required:"true"  column:"field" type:"string"`
	TopControllerKind          string     `json:"top_controller_kind"           required:"true"  column:"field" type:"string"`
	Policy                     string     `json:"policy"                        required:"true"  column:"field" type:"string"`
	PolicyTime                 int64      `json:"policy_time"                   required:"true"  column:"field" type:"int64"`
	PodTotalCost               float64    `json:"pod_total_cost"                required:"true"  column:"field" type:"float64"`
	StartTime                  int64      `json:"start_time"                    required:"true"  column:"field" type:"int64"`
	EndTime                    int64      `json:"end_time"                      required:"true"  column:"field" type:"int64"`
	ResourceLimitCPU           float64    `json:"resource_limit_cpu"            required:"true"  column:"field" type:"float64"`
	ResourceLimitMemory        float64    `json:"resource_limit_memory"         required:"true"  column:"field" type:"float64"`
	InitialResourceLimitCPU    float64    `json:"initial_resource_limit_cpu"    required:"true"  column:"field" type:"float64"`
	InitialResourceLimitMemory float64    `json:"initial_resource_limit_memory" required:"true"  column:"field" type:"float64"`
}

type RecommendationClusterStatusContainerRequest struct {
	DatahubEntity                `scope:"recommendation" category:"cluster_status" type:"container"`
	Metadata                     *Metadata  `measurement:"container" metric:"undefined" boundary:"undefined" quota:"request" ts:"true"`
	Time                         *time.Time `json:"time"                            required:"false" column:"tag"   type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	Namespace                    string     `json:"namespace"                       required:"true"  column:"tag"   type:"string"`
	ClusterName                  string     `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	PodName                      string     `json:"pod_name"                        required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	TopControllerName            string     `json:"top_controller_name"             required:"true"  column:"field" type:"string"`
	TopControllerKind            string     `json:"top_controller_kind"             required:"true"  column:"field" type:"string"`
	Policy                       string     `json:"policy"                          required:"true"  column:"field" type:"string"`
	PolicyTime                   int64      `json:"policy_time"                     required:"true"  column:"field" type:"int64"`
	PodTotalCost                 float64    `json:"pod_total_cost"                  required:"true"  column:"field" type:"float64"`
	StartTime                    int64      `json:"start_time"                      required:"true"  column:"field" type:"int64"`
	EndTime                      int64      `json:"end_time"                        required:"true"  column:"field" type:"int64"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"true"  column:"field" type:"float64"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"true"  column:"field" type:"float64"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"true"  column:"field" type:"float64"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"true"  column:"field" type:"float64"`
}

type RecommendationClusterStatusController struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"controller"`
	Metadata              *Metadata  `measurement:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time `json:"time"                 required:"false" column:"tag"   type:"time"`
	Name                  string     `json:"name"                 required:"true"  column:"tag"   type:"string"`
	Namespace             string     `json:"namespace"            required:"true"  column:"tag"   type:"string"`
	ClusterName           string     `json:"cluster_name"         required:"true"  column:"tag"   type:"string"`
	Type                  string     `json:"type"                 required:"true"  column:"tag"   type:"string"`
	Kind                  string     `json:"kind"                 required:"true"  column:"field" type:"string"`
	CurrentReplicas       int32      `json:"current_replicas"     required:"true"  column:"field" type:"int32"`
	DesiredReplicas       int32      `json:"desired_replicas"     required:"true"  column:"field" type:"int32"`
	CreateTime            int64      `json:"create_time"          required:"true"  column:"field" type:"int64"`
	CurrentCPURequests    float64    `json:"current_cpu_requests" required:"false" column:"field" type:"float64"`
	CurrentMemoryRequests float64    `json:"current_mem_requests" required:"false" column:"field" type:"float64"`
	CurrentCPULimits      float64    `json:"current_cpu_limits"   required:"false" column:"field" type:"float64"`
	CurrentMemoryLimits   float64    `json:"current_mem_limits"   required:"false" column:"field" type:"float64"`
	DesiredCPULimits      float64    `json:"desired_cpu_limits"   required:"false" column:"field" type:"float64"`
	DesiredMemoryLimits   float64    `json:"desired_mem_limits"   required:"false" column:"field" type:"float64"`
	TotalCost             float64    `json:"total_cost"           required:"false" column:"field" type:"float64"`
}

type RecommendationClusterStatusNamespace struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"namespace"`
	Metadata              *Metadata  `measurement:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time `json:"time"                 required:"false" column:"tag"   type:"time"`
	Name                  string     `json:"name"                 required:"true"  column:"tag"   type:"string"`
	ClusterName           string     `json:"cluster_name"         required:"true"  column:"tag"   type:"string"`
	Type                  string     `json:"type"                 required:"true"  column:"tag"   type:"string"`
	Kind                  string     `json:"kind"                 required:"true"  column:"field" type:"string"`
	CurrentReplicas       int32      `json:"current_replicas"     required:"true"  column:"field" type:"int32"`
	DesiredReplicas       int32      `json:"desired_replicas"     required:"true"  column:"field" type:"int32"`
	CreateTime            int64      `json:"create_time"          required:"true"  column:"field" type:"int64"`
	CurrentCPURequests    float64    `json:"current_cpu_requests" required:"false" column:"field" type:"float64"`
	CurrentMemoryRequests float64    `json:"current_mem_requests" required:"false" column:"field" type:"float64"`
	CurrentCPULimits      float64    `json:"current_cpu_limits"   required:"false" column:"field" type:"float64"`
	CurrentMemoryLimits   float64    `json:"current_mem_limits"   required:"false" column:"field" type:"float64"`
	DesiredCPULimits      float64    `json:"desired_cpu_limits"   required:"false" column:"field" type:"float64"`
	DesiredMemoryLimits   float64    `json:"desired_mem_limits"   required:"false" column:"field" type:"float64"`
	TotalCost             float64    `json:"total_cost"           required:"false" column:"field" type:"float64"`
}

type RecommendationClusterStatusNode struct {
	DatahubEntity         `scope:"recommendation" category:"cluster_status" type:"node"`
	Metadata              *Metadata  `measurement:"node" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time `json:"time"                 required:"false" column:"tag"   type:"time"`
	Name                  string     `json:"name"                 required:"true"  column:"tag"   type:"string"`
	ClusterName           string     `json:"cluster_name"         required:"true"  column:"tag"   type:"string"`
	Type                  string     `json:"type"                 required:"true"  column:"tag"   type:"string"`
	Kind                  string     `json:"kind"                 required:"true"  column:"field" type:"string"`
	CurrentReplicas       int32      `json:"current_replicas"     required:"true"  column:"field" type:"int32"`
	DesiredReplicas       int32      `json:"desired_replicas"     required:"true"  column:"field" type:"int32"`
	CreateTime            int64      `json:"create_time"          required:"true"  column:"field" type:"int64"`
	CurrentCPURequests    float64    `json:"current_cpu_requests" required:"false" column:"field" type:"float64"`
	CurrentMemoryRequests float64    `json:"current_mem_requests" required:"false" column:"field" type:"float64"`
	CurrentCPULimits      float64    `json:"current_cpu_limits"   required:"false" column:"field" type:"float64"`
	CurrentMemoryLimits   float64    `json:"current_mem_limits"   required:"false" column:"field" type:"float64"`
	DesiredCPULimits      float64    `json:"desired_cpu_limits"   required:"false" column:"field" type:"float64"`
	DesiredMemoryLimits   float64    `json:"desired_mem_limits"   required:"false" column:"field" type:"float64"`
	TotalCost             float64    `json:"total_cost"           required:"false" column:"field" type:"float64"`
}
