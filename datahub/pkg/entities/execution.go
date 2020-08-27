package entities

import (
	"time"
)

type ExecutionClusterAutoscalerMachineset struct {
	DatahubEntity         `scope:"execution" category:"cluster_autoscaler" type:"machineset"`
	Measurement           *Measurement `name:"cluster_autoscaler_machineset" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time                  *time.Time   `json:"time"                   required:"false" column:"tag"`
	Name                  string       `json:"name"                   required:"true"  column:"tag"`
	Namespace             string       `json:"namespace"              required:"true"  column:"tag"`
	ClusterName           string       `json:"cluster_name"           required:"true"  column:"tag"`
	MachinegroupName      string       `json:"machinegroup_name"      required:"true"  column:"tag"`
	MachinegroupNamespace string       `json:"machinegroup_namespace" required:"true"  column:"tag"`
	NodeName              string       `json:"node_name"              required:"false" column:"field"`
	ReplicasFrom          int32        `json:"replicas_from"          required:"false" column:"field"`
	ReplicasTo            int32        `json:"replicas_to"            required:"false" column:"field"`
	DeltaUpTime           int64        `json:"delta_up_time"          required:"false" column:"field"`
	DeltaDownTime         int64        `json:"delta_down_time"        required:"false" column:"field"`
}
