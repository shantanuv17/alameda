package entities

import (
	"time"
)

type MeteringFederatoraiCapacity struct {
	DatahubEntity     `scope:"metering" category:"federatorai" type:"capacity"`
	Measurement       *Measurement `name:"federatorai_capacity" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time              *time.Time   `json:"time"               required:"false" column:"tag"`
	Name              string       `json:"name"               required:"true"  column:"tag"`
	NumberCluster     int32        `json:"number_cluster"     required:"true"  column:"field"`
	NumberNode        int32        `json:"number_node"        required:"true"  column:"field"`
	NumberNamespace   int32        `json:"number_namespace"   required:"true"  column:"field"`
	NumberApplication int32        `json:"number_application" required:"true"  column:"field"`
	NumberController  int32        `json:"number_controller"  required:"true"  column:"field"`
	CPUCores          int64        `json:"cpu_cores"          required:"true"  column:"field"`
	MemoryBytes       int64        `json:"memory_bytes"       required:"true"  column:"field"`
}
