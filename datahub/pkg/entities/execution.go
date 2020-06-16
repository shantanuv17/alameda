package entities

import (
	"time"
)

type ExecutionClusterAutoscalerMachineset struct {
	DatahubEntity    `scope:"execution" category:"cluster_autoscaler" type:"machineset" measurement:"cluster_autoscaler_machineset" metric:"undefined" boundary:"undefined" quota:"undefined"`
	Time             *time.Time `json:"time"              required:"false" column:"tag"   type:"time"`
	Name             string     `json:"name"              required:"true"  column:"tag"   type:"string"`
	Namespace        string     `json:"namespace"         required:"true"  column:"tag"   type:"string"`
	ClusterName      string     `json:"cluster_name"      required:"true"  column:"tag"   type:"string"`
	MachinegroupName string     `json:"machinegroup_name" required:"true"  column:"tag"   type:"string"`
	NodeName         string     `json:"node_name"         required:"false" column:"field" type:"string"`
	ReplicasFrom     int32      `json:"replicas_from"     required:"false" column:"field" type:"int32"`
	ReplicasTo       int32      `json:"replicas_to"       required:"false" column:"field" type:"int32"`
	DeltaUpTime      int64      `json:"delta_up_time"     required:"false" column:"field" type:"int64"`
	DeltaDownTime    int64      `json:"delta_down_time"   required:"false" column:"field" type:"int64"`
}
