package entities

import (
	"time"
)

type PlanningClusterStatusApplication struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"application"`
	Metadata                     *Metadata  `measurement:"application" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time `json:"time"                            required:"false" column:"tag"   type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	Namespace                    string     `json:"namespace"                       required:"true"  column:"tag"   type:"string"`
	ClusterName                  string     `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	PlanningId                   string     `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string     `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"false" column:"field" type:"int32"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"false" column:"field" type:"int32"`
	ResourceLimitCPU             float64    `json:"resource_limit_cpu"              required:"false" column:"field" type:"int32"`
	ResourceLimitMemory          float64    `json:"resource_limit_memory"           required:"false" column:"field" type:"int32"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"int32"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"false" column:"field" type:"int32"`
	InitialResourceLimitCPU      float64    `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"int32"`
	InitialResourceLimitMemory   float64    `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"int32"`
	StartTime                    int64      `json:"start_time"                      required:"false" column:"field" type:"int32"`
	EndTime                      int64      `json:"end_time"                        required:"false" column:"field" type:"int32"`
	TotalCost                    float64    `json:"total_cost"                      required:"false" column:"field" type:"int32"`
	ApplyPlanningNow             bool       `json:"apply_planning_now"              required:"false" column:"field" type:"int32"`
}

type PlanningClusterStatusCluster struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"cluster"`
	Metadata                     *Metadata  `measurement:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	PlanningId                   string     `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string     `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"false" column:"field" type:"int32"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"false" column:"field" type:"int32"`
	ResourceLimitCPU             float64    `json:"resource_limit_cpu"              required:"false" column:"field" type:"int32"`
	ResourceLimitMemory          float64    `json:"resource_limit_memory"           required:"false" column:"field" type:"int32"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"int32"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"false" column:"field" type:"int32"`
	InitialResourceLimitCPU      float64    `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"int32"`
	InitialResourceLimitMemory   float64    `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"int32"`
	StartTime                    int64      `json:"start_time"                      required:"false" column:"field" type:"int32"`
	EndTime                      int64      `json:"end_time"                        required:"false" column:"field" type:"int32"`
	TotalCost                    float64    `json:"total_cost"                      required:"false" column:"field" type:"int32"`
	ApplyPlanningNow             bool       `json:"apply_planning_now"              required:"false" column:"field" type:"int32"`
}

type PlanningClusterStatusContainer struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"container"`
	Metadata                     *Metadata  `measurement:"container" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	PodName                      string     `json:"pod_name"                        required:"true"  column:"tag"   type:"string"`
	Namespace                    string     `json:"namespace"                       required:"true"  column:"tag"   type:"string"`
	ClusterName                  string     `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	PlanningId                   string     `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string     `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	TopControllerName            string     `json:"top_controller_name"             required:"false" column:"field" type:"string"`
	TopControllerKind            string     `json:"top_controller_kind"             required:"false" column:"field" type:"string"`
	Policy                       string     `json:"policy"                          required:"false" column:"field" type:"string"`
	PolicyTime                   int64      `json:"policy_time"                     required:"false" column:"field" type:"int64"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"false" column:"field" type:"float64"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"false" column:"field" type:"float64"`
	ResourceLimitCPU             float64    `json:"resource_limit_cpu"              required:"false" column:"field" type:"float64"`
	ResourceLimitMemory          float64    `json:"resource_limit_memory"           required:"false" column:"field" type:"float64"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"float64"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"false" column:"field" type:"float64"`
	InitialResourceLimitCPU      float64    `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"float64"`
	InitialResourceLimitMemory   float64    `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"float64"`
	StartTime                    int64      `json:"start_time"                      required:"false" column:"field" type:"int64"`
	EndTime                      int64      `json:"end_time"                        required:"false" column:"field" type:"int64"`
	PodTotalCost                 float64    `json:"pod_total_cost"                  required:"false" column:"field" type:"float64"`
}

type PlanningClusterStatusController struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"controller"`
	Metadata                     *Metadata  `measurement:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	Namespace                    string     `json:"namespace"                       required:"true"  column:"tag"   type:"string"`
	ClusterName                  string     `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	PlanningId                   string     `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string     `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Kind                         string     `json:"kind"                            required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"false" column:"field" type:"float64"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"false" column:"field" type:"float64"`
	ResourceLimitCPU             float64    `json:"resource_limit_cpu"              required:"false" column:"field" type:"float64"`
	ResourceLimitMemory          float64    `json:"resource_limit_memory"           required:"false" column:"field" type:"float64"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"float64"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"false" column:"field" type:"float64"`
	InitialResourceLimitCPU      float64    `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"float64"`
	InitialResourceLimitMemory   float64    `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"float64"`
	StartTime                    int64      `json:"start_time"                      required:"false" column:"field" type:"int64"`
	EndTime                      int64      `json:"end_time"                        required:"false" column:"field" type:"int64"`
	TotalCost                    float64    `json:"total_cost"                      required:"false" column:"field" type:"float64"`
	ApplyPlanningNow             bool       `json:"apply_planning_now"              required:"false" column:"field" type:"bool"`
}

type PlanningClusterStatusNamespace struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"namespace"`
	Metadata                     *Metadata  `measurement:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	ClusterName                  string     `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	PlanningId                   string     `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string     `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"false" column:"field" type:"int32"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"false" column:"field" type:"int32"`
	ResourceLimitCPU             float64    `json:"resource_limit_cpu"              required:"false" column:"field" type:"int32"`
	ResourceLimitMemory          float64    `json:"resource_limit_memory"           required:"false" column:"field" type:"int32"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"int32"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"false" column:"field" type:"int32"`
	InitialResourceLimitCPU      float64    `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"int32"`
	InitialResourceLimitMemory   float64    `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"int32"`
	StartTime                    int64      `json:"start_time"                      required:"false" column:"field" type:"int32"`
	EndTime                      int64      `json:"end_time"                        required:"false" column:"field" type:"int32"`
	TotalCost                    float64    `json:"total_cost"                      required:"false" column:"field" type:"int32"`
	ApplyPlanningNow             bool       `json:"apply_planning_now"              required:"false" column:"field" type:"int32"`
}

type PlanningClusterStatusNode struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"node"`
	Metadata                     *Metadata  `measurement:"node" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                         *time.Time `json:"time" required:"false" column:"tag" type:"time"`
	Name                         string     `json:"name"                            required:"true"  column:"tag"   type:"string"`
	ClusterName                  string     `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	PlanningId                   string     `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string     `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Granularity                  string     `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU           float64    `json:"resource_request_cpu"            required:"false" column:"field" type:"int32"`
	ResourceRequestMemory        float64    `json:"resource_request_memory"         required:"false" column:"field" type:"int32"`
	ResourceLimitCPU             float64    `json:"resource_limit_cpu"              required:"false" column:"field" type:"int32"`
	ResourceLimitMemory          float64    `json:"resource_limit_memory"           required:"false" column:"field" type:"int32"`
	InitialResourceRequestCPU    float64    `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"int32"`
	InitialResourceRequestMemory float64    `json:"initial_resource_request_memory" required:"false" column:"field" type:"int32"`
	InitialResourceLimitCPU      float64    `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"int32"`
	InitialResourceLimitMemory   float64    `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"int32"`
	StartTime                    int64      `json:"start_time"                      required:"false" column:"field" type:"int32"`
	EndTime                      int64      `json:"end_time"                        required:"false" column:"field" type:"int32"`
	TotalCost                    float64    `json:"total_cost"                      required:"false" column:"field" type:"int32"`
	ApplyPlanningNow             bool       `json:"apply_planning_now"              required:"false" column:"field" type:"int32"`
}
