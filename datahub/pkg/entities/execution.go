package entities

type ExecutionMachineset struct {
	DatahubEntity    `scope:"SCOPE_EXECUTION"  category:"cluster_autoscaler"   type:"machineset"`
	Name             string `json:"name"              required:"true"  column:"tag"   type:"DATATYPE_STRING"`
	Namespace        string `json:"namespace"         required:"true"  column:"tag"   type:"DATATYPE_STRING"`
	ClusterName      string `json:"cluster_name"      required:"true"  column:"tag"   type:"DATATYPE_STRING"`
	MachinegroupName string `json:"machinegroup_name" required:"true"  column:"tag"   type:"DATATYPE_STRING"`
	ExecutionTime    string `json:"execution_time"    required:"false" column:"field" type:"DATATYPE_STRING"`
	ReplicasFrom     int32  `json:"replicas_from"     required:"false" column:"field" type:"DATATYPE_INT32"`
	ReplicasTo       int32  `json:"replicas_to"       required:"false" column:"field" type:"DATATYPE_INT32"`
	DeltaUpTime      int64  `json:"delta_up_time​"     required:"false" column:"field" type:"DATATYPE_INT64"`
	DeltaDownTime    int64  `json:"delta_down_time"   required:"false" column:"field" type:"DATATYPE_INT64"`
}
