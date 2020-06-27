package entities

import (
	"time"
)

type ResourceClusterStatusApplication struct {
	DatahubEntity `scope:"resource" category:"cluster_status" type:"application" measurement:"application" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	Namespace     string     `json:"namespace"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	ScalingTool   string     `json:"scaling_tool" required:"true"  column:"tag"   type:"string"`
	Dummy         string     `json:"dummy"        required:"true"  column:"field" type:"string"`
}

type ResourceClusterStatusCluster struct {
	DatahubEntity     `scope:"resource" category:"cluster_status" type:"cluster" measurement:"cluster" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time              *time.Time `json:"time"                required:"false" column:"tag"   type:"time"`
	Name              string     `json:"name"                required:"true"  column:"tag"   type:"string"`
	Uid               string     `json:"uid"                 required:"true"  column:"tag"   type:"string"`
	Value             string     `json:"value"               required:"true"  column:"field" type:"string"`
	CrashingPods      int64      `json:"crashing_pods"       required:"true"  column:"field" type:"int64"`
	CrashingPlanePods int64      `json:"crashing_plane_pods" required:"false" column:"field" type:"int64"`
}

type ResourceClusterStatusContainer struct {
	DatahubEntity                             `scope:"resource" category:"cluster_status" type:"container" measurement:"container" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time                                      *time.Time `json:"time"                                           required:"false" column:"tag"   type:"time"`
	Name                                      string     `json:"name"                                           required:"true"  column:"tag"   type:"string"`
	Namespace                                 string     `json:"namespace"                                      required:"true"  column:"tag"   type:"string"`
	NodeName                                  string     `json:"node_name"                                      required:"true"  column:"tag"   type:"string"`
	ClusterName                               string     `json:"cluster_name"                                   required:"true"  column:"tag"   type:"string"`
	Uid                                       string     `json:"uid"                                            required:"true"  column:"tag"   type:"string"`
	PodName                                   string     `json:"pod_name"                                       required:"true"  column:"tag"   type:"string"`
	TopControllerName                         string     `json:"top_controller_name"                            required:"true"  column:"tag"   type:"string"`
	TopControllerKind                         string     `json:"top_controller_kind"                            required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName                         string     `json:"alameda_scaler_name"                            required:"true"  column:"tag"   type:"string"`
	AlamedaScalerScalingTool                  string     `json:"alameda_scaler_scaling_tool"                    required:"true"  column:"tag"   type:"string"`
	ResourceRequestCPU                        string     `json:"resource_request_cpu"                           required:"false" column:"field" type:"string"`
	ResourceRequestMemory                     string     `json:"resource_request_memory"                        required:"false" column:"field" type:"string"`
	ResourceLimitCpu                          string     `json:"resource_limit_cpu"                             required:"false" column:"field" type:"string"`
	ResourceLimitMemory                       string     `json:"resource_limit_memory"                          required:"false" column:"field" type:"string"`
	StatusWaitingReason                       string     `json:"status_waiting_reason"                          required:"false" column:"field" type:"string"`
	StatusWaitingMessage                      string     `json:"status_waiting_message"                         required:"false" column:"field" type:"string"`
	StatusRunningStartTt                      int64      `json:"status_running_start_at"                        required:"false" column:"field" type:"int64"`
	StatusTerminatedExitCode                  int32      `json:"status_terminated_exit_code"                    required:"false" column:"field" type:"int32"`
	StatusTerminatedReason                    string     `json:"status_terminated_reason"                       required:"false" column:"field" type:"string"`
	StatusTerminatedMessage                   string     `json:"status_terminated_message"                      required:"false" column:"field" type:"string"`
	StatusTerminatedStartedAt                 int64      `json:"status_terminated_started_at"                   required:"false" column:"field" type:"int64"`
	StatusTerminatedFinishedAt                int64      `json:"status_terminated_finished_at"                  required:"false" column:"field" type:"int64"`
	LastTerminationStatusWaitingReason        string     `json:"last_termination_status_waiting_reason"         required:"false" column:"field" type:"string"`
	LastTerminationStatusWaitingMessage       string     `json:"last_termination_status_waiting_message"        required:"false" column:"field" type:"string"`
	LastTerminationStatusRunningStartAt       int64      `json:"last_termination_status_running_start_at"       required:"false" column:"field" type:"int64"`
	LastTerminationStatusTerminatedExitCode   int32      `json:"last_termination_status_terminated_exit_code"   required:"false" column:"field" type:"int32"`
	LastTerminationStatusTerminatedReason     string     `json:"last_termination_status_terminated_reason"      required:"false" column:"field" type:"string"`
	LastTerminationStatusTerminatedMessage    string     `json:"last_termination_status_terminated_message"     required:"false" column:"field" type:"string"`
	LastTerminationStatusTerminatedStartedAt  int64      `json:"last_termination_status_terminated_started_at"  required:"false" column:"field" type:"int64"`
	LastTerminationStatusTerminatedFinishedAt int64      `json:"last_termination_status_terminated_finished_at" required:"false" column:"field" type:"int64"`
	RestartCount                              int32      `json:"restart_count"                                  required:"false" column:"field" type:"int32"`
}

type ResourceClusterStatusController struct {
	DatahubEntity            `scope:"resource" category:"cluster_status" type:"controller" measurement:"controller" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time                     *time.Time `json:"time"                        required:"false" column:"tag"   type:"time"`
	Name                     string     `json:"name"                        required:"true"  column:"tag"   type:"string"`
	Namespace                string     `json:"namespace"                   required:"true"  column:"tag"   type:"string"`
	ClusterName              string     `json:"cluster_name"                required:"true"  column:"tag"   type:"string"`
	Uid                      string     `json:"uid"                         required:"true"  column:"tag"   type:"string"`
	Kind                     string     `json:"kind"                        required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName        string     `json:"alameda_scaler_name"         required:"true"  column:"tag"   type:"string"`
	AlamedaScalerScalingTool string     `json:"alameda_scaler_scaling_tool" required:"true"  column:"tag"   type:"string"`
	Replicas                 int32      `json:"replicas"                    required:"false" column:"field" type:"int32"`
	SpecReplicas             int32      `json:"spec_replicas"               required:"false" column:"field" type:"int32"`
	Policy                   string     `json:"policy"                      required:"false" column:"field" type:"string"`
	EnableExecution          bool       `json:"enable_execution"            required:"false" column:"field" type:"bool"`
}

type ResourceClusterStatusNamespace struct {
	DatahubEntity `scope:"resource" category:"cluster_status" type:"namespace" measurement:"namespace" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time          *time.Time `json:"time"         required:"false" column:"tag"   type:"time"`
	Name          string     `json:"name"         required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"cluster_name" required:"true"  column:"tag"   type:"string"`
	Uid           string     `json:"uid"          required:"true"  column:"tag"   type:"string"`
	Value         string     `json:"value"        required:"true"  column:"field" type:"string"`
}

type ResourceClusterStatusNode struct {
	DatahubEntity               `scope:"resource" category:"cluster_status" type:"node" measurement:"node" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time                        *time.Time `json:"time"                          required:"false" column:"tag"   type:"time"`
	Name                        string     `json:"name"                          required:"true"  column:"tag"   type:"string"`
	ClusterName                 string     `json:"cluster_name"                  required:"true"  column:"tag"   type:"string"`
	Uid                         string     `json:"uid"                           required:"true"  column:"tag"   type:"string"`
	CreateTime                  int64      `json:"create_time"                   required:"false" column:"field" type:"int64"`
	NodeCPUCores                int64      `json:"node_cpu_cores"                required:"false" column:"field" type:"int64"`  // NodeCPUCores is the amount of cores in node
	NodeMemoryBytes             int64      `json:"node_memory_bytes"             required:"false" column:"field" type:"int64"`  // NodeMemoryBytes is the amount of memory bytes in node
	NodeNetworkMbps             int64      `json:"node_network_mbps"             required:"false" column:"field" type:"int64"`  // NodeNetworkMbps is mega bits per second
	IOProvider                  string     `json:"io_provider"                   required:"false" column:"field" type:"string"` // Cloud service provider
	IOInstanceType              string     `json:"io_instance_type"              required:"false" column:"field" type:"string"`
	IORegion                    string     `json:"io_region"                     required:"false" column:"field" type:"string"`
	IOZone                      string     `json:"io_zone"                       required:"false" column:"field" type:"string"`
	IOOs                        string     `json:"io_os"                         required:"false" column:"field" type:"string"`
	IORole                      string     `json:"io_role"                       required:"false" column:"field" type:"string"`
	IOInstanceId                string     `json:"io_instance_id"                required:"false" column:"field" type:"string"`
	IOStorageSize               int64      `json:"io_storage_size"               required:"false" column:"field" type:"int64"`
	ConditionReady              bool       `json:"condition_ready"               required:"false" column:"field" type:"bool"`
	ConditionDiskPressure       bool       `json:"condition_disk_pressure"       required:"false" column:"field" type:"bool"`
	ConditionMemoryPressure     bool       `json:"condition_memory_pressure"     required:"false" column:"field" type:"bool"`
	ConditionPIDPressure        bool       `json:"condition_pid_pressure"        required:"false" column:"field" type:"bool"`
	ConditionNetworkUnavailable bool       `json:"condition_network_unavailable" required:"false" column:"field" type:"bool"`
}

type ResourceClusterStatusPod struct {
	DatahubEntity                      `scope:"resource" category:"cluster_status" type:"pod" measurement:"pod" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time                               *time.Time `json:"time"                                   required:"false" column:"tag"   type:"time"`
	Name                               string     `json:"name"                                   required:"true"  column:"tag"   type:"string"`
	Namespace                          string     `json:"namespace"                              required:"true"  column:"tag"   type:"string"`
	NodeName                           string     `json:"node_name"                              required:"true"  column:"tag"   type:"string"`
	ClusterName                        string     `json:"cluster_name"                           required:"true"  column:"tag"   type:"string"`
	Uid                                string     `json:"uid"                                    required:"true"  column:"tag"   type:"string"`
	TopControllerName                  string     `json:"top_controller_name"                    required:"true"  column:"tag"   type:"string"`
	TopControllerKind                  string     `json:"top_controller_kind"                    required:"true"  column:"tag"   type:"string"`
	AlamedaScalerName                  string     `json:"alameda_scaler_name"                    required:"true"  column:"tag"   type:"string"`
	AlamedaScalerScalingTool           string     `json:"alameda_scaler_scaling_tool"            required:"true"  column:"tag"   type:"string"`
	AppName                            string     `json:"app_name"                               required:"true"  column:"tag"   type:"string"`
	AppPartOf                          string     `json:"app_part_of"                            required:"true"  column:"tag"   type:"string"`
	PodCreateTime                      int64      `json:"pod_create_time"                        required:"false" column:"field" type:"int64"`
	ResourceLink                       string     `json:"resource_link"                          required:"false" column:"field" type:"string"`
	TopControllerReplicas              int32      `json:"top_controller_replicas"                required:"false" column:"field" type:"int32"`
	PodPhase                           string     `json:"pod_phase"                              required:"false" column:"field" type:"string"`
	PodMessage                         string     `json:"pod_message"                            required:"false" column:"field" type:"string"`
	PodReason                          string     `json:"pod_reason"                             required:"false" column:"field" type:"string"`
	Policy                             string     `json:"policy"                                 required:"false" column:"field" type:"string"`
	UsedRecommendationDd               string     `json:"used_recommendation_id"                 required:"false" column:"field" type:"string"`
	AlamedaScalerResourceLimitCPU      string     `json:"alameda_scaler_resource_limit_cpu"      required:"false" column:"field" type:"string"`
	AlamedaScalerResourceLimitMemory   string     `json:"alameda_scaler_resource_limit_memory"   required:"false" column:"field" type:"string"`
	AlamedaScalerResourceRequestCPU    string     `json:"alameda_scaler_resource_request_cpu"    required:"false" column:"field" type:"string"`
	AlamedaScalerResourceRequestMemory string     `json:"alameda_scaler_resource_request_memory" required:"false" column:"field" type:"string"`
}
