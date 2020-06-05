package entities

type PlanningClusterStatusApplication struct {
	DatahubEntity                `scope:"planning" category:"cluster_status" type:"application" measurement:"application" metric:"undefined" boundary:"undefined" quota:"undefined"`
	PlanningId                   string `json:"planning_id"                     required:"true"  column:"tag"   type:"string"`
	PlanningType                 string `json:"planning_type"                   required:"true"  column:"tag"   type:"string"`
	Time                         string `json:"time"                            required:"true"  column:"tag"   type:"string"`
	ClusterName                  string `json:"cluster_name"                    required:"true"  column:"tag"   type:"string"`
	Namespace                    string `json:"namespace"                       required:"true"  column:"tag"   type:"string"`
	Name                         string `json:"name"                            required:"true"  column:"tag"   type:"string"`
	Granularity                  string `json:"granularity"                     required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU           int32  `json:"resource_request_cpu"            required:"false" column:"field" type:"int32"`
	ResourceRequestMemory        int32  `json:"resource_request_memory"         required:"false" column:"field" type:"int32"`
	ResourceLimitCPU             int32  `json:"resource_limit_cpu"              required:"false" column:"field" type:"int32"`
	ResourceLimitMemory          int32  `json:"resource_limit_memory"           required:"false" column:"field" type:"int32"`
	InitialResourceRequestCPU    int32  `json:"initial_resource_request_cpu"    required:"false" column:"field" type:"int32"`
	InitialResourceRequestMemory int32  `json:"initial_resource_request_memory" required:"false" column:"field" type:"int32"`
	InitialResourceLimitCPU      int32  `json:"initial_resource_limit_cpu"      required:"false" column:"field" type:"int32"`
	InitialResourceLimitMemory   int32  `json:"initial_resource_limit_memory"   required:"false" column:"field" type:"int32"`
	StartTime                    int32  `json:"start_time"                      required:"false" column:"field" type:"int32"`
	EndTime                      int32  `json:"end_time"                        required:"false" column:"field" type:"int32"`
	TotalCost                    int32  `json:"total_cost"                      required:"false" column:"field" type:"int32"`
	ApplyPlanningNow             int32  `json:"apply_planning_now"              required:"false" column:"field" type:"int32"`
}

type PlanningClusterStatusCluster struct {
	DatahubEntity `scope:"planning" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined"`
}

type PlanningClusterStatusContainer struct {
	DatahubEntity `scope:"planning" category:"cluster_status" type:"container" measurement:"container" metric:"undefined" boundary:"undefined" quota:"undefined"`
}

type PlanningClusterStatusController struct {
	DatahubEntity `scope:"planning" category:"cluster_status" type:"controller" measurement:"controller" metric:"undefined" boundary:"undefined" quota:"undefined"`
}

type PlanningClusterStatusNamespace struct {
	DatahubEntity `scope:"planning" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined"`
}

type PlanningClusterStatusNode struct {
	DatahubEntity `scope:"planning" category:"cluster_status" type:"node" measurement:"node" metric:"undefined" boundary:"undefined" quota:"undefined"`
}
