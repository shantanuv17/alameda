package entities

import (
	"time"
)

type ResourceClusterAutoscalerMachinegroup struct {
	DatahubEntity                         `scope:"resource" category:"cluster_autoscaler" type:"machinegroup"`
	Measurement                           *Measurement `name:"cluster_autoscaler_machinegroup" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                                  *time.Time   `json:"time"                                      required:"false" column:"tag"`
	Name                                  string       `json:"name"                                      required:"true"  column:"tag"`
	Namespace                             string       `json:"namespace"                                 required:"true"  column:"tag"`
	ClusterName                           string       `json:"cluster_name"                              required:"true"  column:"tag"`
	AlamedaScalerName                     string       `json:"alameda_scaler_name"                       required:"true"  column:"tag"`
	AlamedaScalerNamespace                string       `json:"alameda_scaler_namespace"                  required:"true"  column:"tag"`
	CPUMetricUtilizationTarget            int32        `json:"cpu_metric_utilization_target"             required:"false" column:"field"`
	CPUMetricScaleUpGap                   int32        `json:"cpu_metric_scaleup_gap"                    required:"false" column:"field"`
	CPUMetricScaleDownGap                 int32        `json:"cpu_metric_scaledown_gap"                  required:"false" column:"field"`
	MemoryMetricUtilizationTarget         int32        `json:"memory_metric_utilization_target"          required:"false" column:"field"`
	MemoryMetricScaleUpGap                int32        `json:"memory_metric_scaleup_gap"                 required:"false" column:"field"`
	MemoryMetricScaleDownGap              int32        `json:"memory_metric_scaledown_gap"               required:"false" column:"field"`
	CPUDurationUpThresholdPercentage      int32        `json:"cpu_duration_up_threshold_percentage"      required:"false" column:"field"`
	CPUDurationDownThresholdPercentage    int32        `json:"cpu_duration_down_threshold_percentage"    required:"false" column:"field"`
	MemoryDurationUpThresholdPercentage   int32        `json:"memory_duration_up_threshold_percentage"   required:"false" column:"field"`
	MemoryDurationDownThresholdPercentage int32        `json:"memory_duration_down_threshold_percentage" required:"false" column:"field"`
}

type ResourceClusterAutoscalerMachineset struct {
	DatahubEntity           `scope:"resource" category:"cluster_autoscaler" type:"machineset"`
	Measurement             *Measurement `name:"cluster_autoscaler_machineset" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                    *time.Time   `json:"time"                       required:"false" column:"tag"`
	Name                    string       `json:"name"                       required:"true"  column:"tag"`
	Namespace               string       `json:"namespace"                  required:"true"  column:"tag"`
	ClusterName             string       `json:"cluster_name"               required:"true"  column:"tag"`
	MachinegroupName        string       `json:"machinegroup_name"          required:"true"  column:"tag"`
	MachinegroupNamespace   string       `json:"machinegroup_namespace"     required:"true"  column:"tag"`
	ResourceK8sSpecReplicas int32        `json:"resource_k8s_spec_replicas" required:"false" column:"field"`
	ResourceK8sReplicas     int32        `json:"resource_k8s_replicas"      required:"false" column:"field"`
	ResourceK8sMinReplicas  int32        `json:"resource_k8s_min_replicas"  required:"false" column:"field"`
	ResourceK8sMaxReplicas  int32        `json:"resource_k8s_max_replicas"  required:"false" column:"field"`
	EnableExecution         bool         `json:"enable_execution"           required:"false" column:"field"`
	ScaleUpDelay            int64        `json:"scale_up_delay"             required:"false" column:"field"`
	ScaleDownDelay          int64        `json:"scale_down_delay"           required:"false" column:"field"`
}

type ResourceClusterStatusApplication struct {
	DatahubEntity `scope:"resource" category:"cluster_status" type:"application"`
	Measurement   *Measurement `name:"application" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	Namespace     string       `json:"namespace"    required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Dummy         string       `json:"dummy"        required:"true"  column:"field"`
}

type ResourceClusterStatusCluster struct {
	DatahubEntity     `scope:"resource" category:"cluster_status" type:"cluster"`
	Measurement       *Measurement `name:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time              *time.Time   `json:"time"                required:"false" column:"tag"`
	Name              string       `json:"name"                required:"true"  column:"tag"`
	Uid               string       `json:"uid"                 required:"true"  column:"tag"`
	Value             string       `json:"value"               required:"true"  column:"field"`
	CrashingPods      int64        `json:"crashing_pods"       required:"true"  column:"field"`
	CrashingPlanePods int64        `json:"crashing_plane_pods" required:"false" column:"field"`
}

type ResourceClusterStatusContainer struct {
	DatahubEntity                             `scope:"resource" category:"cluster_status" type:"container"`
	Measurement                               *Measurement `name:"container" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                                      *time.Time   `json:"time"                                           required:"false" column:"tag"`
	Name                                      string       `json:"name"                                           required:"true"  column:"tag"`
	Namespace                                 string       `json:"namespace"                                      required:"true"  column:"tag"`
	NodeName                                  string       `json:"node_name"                                      required:"true"  column:"tag"`
	ClusterName                               string       `json:"cluster_name"                                   required:"true"  column:"tag"`
	Uid                                       string       `json:"uid"                                            required:"true"  column:"tag"`
	PodName                                   string       `json:"pod_name"                                       required:"true"  column:"tag"`
	TopControllerName                         string       `json:"top_controller_name"                            required:"true"  column:"tag"`
	TopControllerKind                         Kind         `json:"top_controller_kind"                            required:"true"  column:"tag"`
	AlamedaScalerName                         string       `json:"alameda_scaler_name"                            required:"true"  column:"tag"`
	AlamedaScalerNamespace                    string       `json:"alameda_scaler_namespace"                       required:"true"  column:"tag"`
	AlamedaScalerScalingTool                  ScalingTool  `json:"alameda_scaler_scaling_tool"                    required:"true"  column:"tag"`
	ResourceRequestCPU                        string       `json:"resource_request_cpu"                           required:"true"  column:"field"`
	ResourceRequestMemory                     string       `json:"resource_request_memory"                        required:"false" column:"field"`
	ResourceLimitCpu                          string       `json:"resource_limit_cpu"                             required:"false" column:"field"`
	ResourceLimitMemory                       string       `json:"resource_limit_memory"                          required:"false" column:"field"`
	StatusWaitingReason                       string       `json:"status_waiting_reason"                          required:"false" column:"field"`
	StatusWaitingMessage                      string       `json:"status_waiting_message"                         required:"false" column:"field"`
	StatusRunningStartTt                      int64        `json:"status_running_start_at"                        required:"false" column:"field"`
	StatusTerminatedExitCode                  int32        `json:"status_terminated_exit_code"                    required:"false" column:"field"`
	StatusTerminatedReason                    string       `json:"status_terminated_reason"                       required:"false" column:"field"`
	StatusTerminatedMessage                   string       `json:"status_terminated_message"                      required:"false" column:"field"`
	StatusTerminatedStartedAt                 int64        `json:"status_terminated_started_at"                   required:"false" column:"field"`
	StatusTerminatedFinishedAt                int64        `json:"status_terminated_finished_at"                  required:"false" column:"field"`
	LastTerminationStatusWaitingReason        string       `json:"last_termination_status_waiting_reason"         required:"false" column:"field"`
	LastTerminationStatusWaitingMessage       string       `json:"last_termination_status_waiting_message"        required:"false" column:"field"`
	LastTerminationStatusRunningStartAt       int64        `json:"last_termination_status_running_start_at"       required:"false" column:"field"`
	LastTerminationStatusTerminatedExitCode   int32        `json:"last_termination_status_terminated_exit_code"   required:"false" column:"field"`
	LastTerminationStatusTerminatedReason     string       `json:"last_termination_status_terminated_reason"      required:"false" column:"field"`
	LastTerminationStatusTerminatedMessage    string       `json:"last_termination_status_terminated_message"     required:"false" column:"field"`
	LastTerminationStatusTerminatedStartedAt  int64        `json:"last_termination_status_terminated_started_at"  required:"false" column:"field"`
	LastTerminationStatusTerminatedFinishedAt int64        `json:"last_termination_status_terminated_finished_at" required:"false" column:"field"`
	RestartCount                              int32        `json:"restart_count"                                  required:"false" column:"field"`
}

type ResourceClusterStatusController struct {
	DatahubEntity            `scope:"resource" category:"cluster_status" type:"controller"`
	Measurement              *Measurement `name:"controller" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                     *time.Time   `json:"time"                        required:"false" column:"tag"`
	Name                     string       `json:"name"                        required:"true"  column:"tag"`
	Namespace                string       `json:"namespace"                   required:"true"  column:"tag"`
	ClusterName              string       `json:"cluster_name"                required:"true"  column:"tag"`
	Uid                      string       `json:"uid"                         required:"true"  column:"tag"`
	Kind                     Kind         `json:"kind"                        required:"true"  column:"tag"`
	AlamedaScalerName        string       `json:"alameda_scaler_name"         required:"true"  column:"tag"`
	AlamedaScalerNamespace   string       `json:"alameda_scaler_namespace"    required:"true"  column:"tag"`
	AlamedaScalerScalingTool ScalingTool  `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"`
	Replicas                 int32        `json:"replicas"                    required:"true" column:"field"`
	SpecReplicas             int32        `json:"spec_replicas"               required:"false" column:"field"`
	ResourceK8sMinReplicas   int32        `json:"resource_k8s_min_replicas"   required:"false" column:"field"`
	ResourceK8sMaxReplicas   int32        `json:"resource_k8s_max_replicas"   required:"false" column:"field"`
	Policy                   Policy       `json:"policy"                      required:"false" column:"field"`
	EnableExecution          bool         `json:"enable_execution"            required:"false" column:"field"`
}

type ResourceClusterStatusNamespace struct {
	DatahubEntity `scope:"resource" category:"cluster_status" type:"namespace"`
	Measurement   *Measurement `name:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time          *time.Time   `json:"time"         required:"false" column:"tag"`
	Name          string       `json:"name"         required:"true"  column:"tag"`
	ClusterName   string       `json:"cluster_name" required:"true"  column:"tag"`
	Uid           string       `json:"uid"          required:"true"  column:"tag"`
	Value         string       `json:"value"        required:"true"  column:"field"`
}

type ResourceClusterStatusNode struct {
	DatahubEntity               `scope:"resource" category:"cluster_status" type:"node"`
	Measurement                 *Measurement `name:"node" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                        *time.Time   `json:"time"                          required:"false" column:"tag"`
	Name                        string       `json:"name"                          required:"true"  column:"tag"`
	ClusterName                 string       `json:"cluster_name"                  required:"true"  column:"tag"`
	Uid                         string       `json:"uid"                           required:"true"  column:"tag"`
	MachinesetName              string       `json:"machineset_name"               required:"true"  column:"tag"`
	MachinesetNamespace         string       `json:"machineset_namespace"          required:"true"  column:"tag"`
	RoleMaster                  bool         `json:"role_master"                   required:"false" column:"field"`
	RoleWorker                  bool         `json:"role_worker"                   required:"false" column:"field"`
	RoleInfra                   bool         `json:"role_infra"                    required:"false" column:"field"`
	MachineCreateTime           int64        `json:"machine_create_time"           required:"false" column:"field"`
	CreateTime                  int64        `json:"create_time"                   required:"true"  column:"field"`
	NodeCPUCores                int64        `json:"node_cpu_cores"                required:"false" column:"field"` // NodeCPUCores is the amount of cores in node
	NodeMemoryBytes             int64        `json:"node_memory_bytes"             required:"false" column:"field"` // NodeMemoryBytes is the amount of memory bytes in node
	NodeNetworkMbps             int64        `json:"node_network_mbps"             required:"false" column:"field"` // NodeNetworkMbps is mega bits per second
	IOProvider                  string       `json:"io_provider"                   required:"false" column:"field"` // Cloud service provider
	IOInstanceType              string       `json:"io_instance_type"              required:"false" column:"field"`
	IORegion                    string       `json:"io_region"                     required:"false" column:"field"`
	IOZone                      string       `json:"io_zone"                       required:"false" column:"field"`
	IOOs                        string       `json:"io_os"                         required:"false" column:"field"`
	IORole                      string       `json:"io_role"                       required:"false" column:"field"`
	IOInstanceId                string       `json:"io_instance_id"                required:"false" column:"field"`
	IOStorageSize               int64        `json:"io_storage_size"               required:"false" column:"field"`
	ConditionReady              bool         `json:"condition_ready"               required:"false" column:"field"`
	ConditionDiskPressure       bool         `json:"condition_disk_pressure"       required:"false" column:"field"`
	ConditionMemoryPressure     bool         `json:"condition_memory_pressure"     required:"false" column:"field"`
	ConditionPIDPressure        bool         `json:"condition_pid_pressure"        required:"false" column:"field"`
	ConditionNetworkUnavailable bool         `json:"condition_network_unavailable" required:"false" column:"field"`
}

type ResourceClusterStatusPod struct {
	DatahubEntity                      `scope:"resource" category:"cluster_status" type:"pod"`
	Measurement                        *Measurement `name:"pod" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                               *time.Time   `json:"time"                                   required:"false" column:"tag"`
	Name                               string       `json:"name"                                   required:"true"  column:"tag"`
	Namespace                          string       `json:"namespace"                              required:"true"  column:"tag"`
	NodeName                           string       `json:"node_name"                              required:"true"  column:"tag"`
	ClusterName                        string       `json:"cluster_name"                           required:"true"  column:"tag"`
	Uid                                string       `json:"uid"                                    required:"true"  column:"tag"`
	TopControllerName                  string       `json:"top_controller_name"                    required:"true"  column:"tag"`
	TopControllerKind                  Kind         `json:"top_controller_kind"                    required:"true"  column:"tag"`
	AlamedaScalerName                  string       `json:"alameda_scaler_name"                    required:"true"  column:"tag"`
	AlamedaScalerNamespace             string       `json:"alameda_scaler_namespace"               required:"true"  column:"tag"`
	AlamedaScalerScalingTool           ScalingTool  `json:"alameda_scaler_scaling_tool"            required:"true"  column:"tag"`
	AppName                            string       `json:"app_name"                               required:"true"  column:"tag"`
	AppPartOf                          string       `json:"app_part_of"                            required:"true"  column:"tag"`
	PodCreateTime                      int64        `json:"pod_create_time"                        required:"true"  column:"field"`
	ResourceLink                       string       `json:"resource_link"                          required:"false" column:"field"`
	TopControllerReplicas              int32        `json:"top_controller_replicas"                required:"false" column:"field"`
	PodPhase                           PodPhase     `json:"pod_phase"                              required:"false" column:"field"`
	PodMessage                         string       `json:"pod_message"                            required:"false" column:"field"`
	PodReason                          string       `json:"pod_reason"                             required:"false" column:"field"`
	Policy                             Policy       `json:"policy"                                 required:"false" column:"field"`
	UsedRecommendationDd               string       `json:"used_recommendation_id"                 required:"false" column:"field"`
	AlamedaScalerResourceLimitCPU      string       `json:"alameda_scaler_resource_limit_cpu"      required:"false" column:"field"`
	AlamedaScalerResourceLimitMemory   string       `json:"alameda_scaler_resource_limit_memory"   required:"false" column:"field"`
	AlamedaScalerResourceRequestCPU    string       `json:"alameda_scaler_resource_request_cpu"    required:"false" column:"field"`
	AlamedaScalerResourceRequestMemory string       `json:"alameda_scaler_resource_request_memory" required:"false" column:"field"`
}
