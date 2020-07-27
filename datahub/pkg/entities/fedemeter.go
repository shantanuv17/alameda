package entities

import (
	"time"
)

type FedemeterCalculationInstance struct {
	DatahubEntity   `scope:"fedemeter" category:"calculation" type:"price"`
	Metadata        *Metadata  `measurement:"calculation_price_instance" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time            *time.Time `json:"time"            required:"false" column:"tag"   type:"time"`
	NodeName        string     `json:"nodename"        required:"true"  column:"tag"   type:"string"`
	ClusterName     string     `json:"clustername"     required:"true"  column:"tag"   type:"string"`
	Provider        string     `json:"provider"        required:"true"  column:"tag"   type:"string"`
	Region          string     `json:"region"          required:"true"  column:"tag"   type:"string"`
	Unit            string     `json:"unit"            required:"true"  column:"tag"   type:"string"`
	Granularity     string     `json:"granularity"     required:"true"  column:"tag"   type:"string"`
	CPU             float64    `json:"cpu"             required:"false" column:"field" type:"float64"`
	Memory          float64    `json:"memory"          required:"false" column:"field" type:"float64"`
	Description     string     `json:"description"     required:"false" column:"field" type:"string"`
	DisplayName     string     `json:"displayname"     required:"false" column:"field" type:"string"`
	InstanceNum     int32      `json:"instancenum"     required:"false" column:"field" type:"int32"`
	InstanceType    string     `json:"instancetype"    required:"false" column:"field" type:"string"`
	NodeType        string     `json:"nodetype"        required:"false" column:"field" type:"string"`
	OperatingSystem string     `json:"operatingsystem" required:"false" column:"field" type:"string"`
	Period          int32      `json:"period"          required:"false" column:"field" type:"int32"`
	PreinstalledSW  string     `json:"preinstalledsw"  required:"false" column:"field" type:"string"`
	StartTime       int64      `json:"starttime"       required:"false" column:"field" type:"int64"`
	Cost            float64    `json:"cost"            required:"false" column:"field" type:"float64"`
	TotalCost       float64    `json:"totalcost"       required:"false" column:"field" type:"float64"`
}

type FedemeterCalculationStorage struct {
	DatahubEntity `scope:"fedemeter" category:"calculation" type:"price"`
	Metadata      *Metadata  `measurement:"calculation_price_storage" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time          *time.Time `json:"time"        required:"false" column:"tag"   type:"time"`
	NodeName      string     `json:"nodename"    required:"true"  column:"tag"   type:"string"`
	ClusterName   string     `json:"clustername" required:"true"  column:"tag"   type:"string"`
	Provider      string     `json:"provider"    required:"true"  column:"tag"   type:"string"`
	Unit          string     `json:"unit"        required:"true"  column:"tag"   type:"string"`
	Granularity   string     `json:"granularity" required:"true"  column:"tag"   type:"string"`
	Description   string     `json:"description" required:"false" column:"field" type:"string"`
	DisplayName   string     `json:"displayname" required:"false" column:"field" type:"string"`
	Period        int32      `json:"period"      required:"false" column:"field" type:"int32"`
	StorageNum    int32      `json:"storagenum"  required:"false" column:"field" type:"int32"`
	StorageSize   int64      `json:"storagesize" required:"false" column:"field" type:"int64"`
	VolumeType    string     `json:"volumetype"  required:"false" column:"field" type:"string"`
	StartTime     int64      `json:"starttime"   required:"false" column:"field" type:"int64"`
	Cost          float64    `json:"cost"        required:"false" column:"field" type:"float64"`
}

type FedemeterRecommendationJERI struct {
	DatahubEntity     `scope:"fedemeter" category:"recommendation" type:"jeri"`
	Metadata          *Metadata  `measurement:"recommendation_jeri" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time              *time.Time `json:"time"                required:"false" column:"tag"   type:"time"`
	ClusterName       string     `json:"clustername"         required:"true"  column:"tag"   type:"string"`
	Country           string     `json:"country"             required:"true"  column:"tag"   type:"string"`
	InstanceType      string     `json:"instancetype"        required:"true"  column:"tag"   type:"string"`
	Provider          string     `json:"provider"            required:"true"  column:"tag"   type:"string"`
	Rank              string     `json:"rank"                required:"true"  column:"tag"   type:"string"`
	Region            string     `json:"region"              required:"true"  column:"tag"   type:"string"`
	ReservedInstances string     `json:"reservedinstances"   required:"true"  column:"tag"   type:"string"`
	Granularity       string     `json:"granularity"         required:"true"  column:"tag"   type:"string"`
	MasterNum         int32      `json:"master_num"          required:"false" column:"field" type:"int32"`
	WorkerNum         int32      `json:"worker_num"          required:"false" column:"field" type:"int32"`
	MasterStorageSize float64    `json:"master_storage_size" required:"false" column:"field" type:"float64"`
	WorkerStorageSize float64    `json:"worker_storage_size" required:"false" column:"field" type:"float64"`
	AccCost           float64    `json:"acc_cost"            required:"false" column:"field" type:"float64"`
	DisplayName       string     `json:"display_name"        required:"false" column:"field" type:"string"`
	ResourceName      string     `json:"resource_name"       required:"false" column:"field" type:"string"`
	StartTime         int64      `json:"start_time"          required:"false" column:"field" type:"int64"`
	OndemandNum       int32      `json:"ondemand_num"        required:"false" column:"field" type:"int32"`
	MasterRiNum       int32      `json:"master_ri_num"       required:"false" column:"field" type:"int32"`
	WorkerRiNum       int32      `json:"worker_ri_num"       required:"false" column:"field" type:"int32"`
	MasterSpotNum     int32      `json:"master_spot_num"     required:"false" column:"field" type:"int32"`
	WorkerSpotNum     int32      `json:"worker_spot_num"     required:"false" column:"field" type:"int32"`
	MasterOndemandNum int32      `json:"master_ondemand_num" required:"false" column:"field" type:"int32"`
	WorkerOndemandNum int32      `json:"worker_ondemand_num" required:"false" column:"field" type:"int32"`
	Cost              float64    `json:"cost"                required:"false" column:"field" type:"float64"`
	TotalCost         float64    `json:"total_cost"          required:"false" column:"field" type:"float64"`
}

type FedemeterResourceHistoryCostApp struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_history" type:"cost"`
	Metadata       *Metadata  `measurement:"resource_history_cost_app" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time `json:"time"           required:"false" column:"tag"   type:"time"`
	AppName        string     `json:"appname"        required:"true"  column:"tag"   type:"string"`
	NamespaceName  string     `json:"namespacename"  required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"clustername"    required:"true"  column:"tag"   type:"string"`
	Provider       string     `json:"provider"       required:"true"  column:"tag"   type:"string"`
	Type           string     `json:"type"           required:"true"  column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"    required:"true"  column:"tag"   type:"string"`
	CreateTime     int64      `json:"createtime"     required:"false" column:"field" type:"int64"`
	StartTime      int64      `json:"starttime"      required:"false" column:"field" type:"int64"`
	CostPercentage float64    `json:"costpercentage" required:"false" column:"field" type:"float64"`
	WorkloadCost   float64    `json:"workloadcost"   required:"false" column:"field" type:"float64"`
}

type FedemeterResourceHistoryCostNamespace struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_history" type:"cost"`
	Metadata       *Metadata  `measurement:"resource_history_cost_namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time `json:"time"           required:"false" column:"tag"   type:"time"`
	NamespaceName  string     `json:"namespacename"  required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"clustername"    required:"true"  column:"tag"   type:"string"`
	Provider       string     `json:"provider"       required:"true"  column:"tag"   type:"string"`
	Type           string     `json:"type"           required:"true"  column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"    required:"true"  column:"tag"   type:"string"`
	CreateTime     int64      `json:"createtime"     required:"false" column:"field" type:"int64"`
	StartTime      int64      `json:"starttime"      required:"false" column:"field" type:"int64"`
	CostPercentage float64    `json:"costpercentage" required:"false" column:"field" type:"float64"`
	WorkloadCost   float64    `json:"workloadcost"   required:"false" column:"field" type:"float64"`
}

type FedemeterResourcePredictionCostApp struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_prediction" type:"cost"`
	Metadata       *Metadata  `measurement:"resource_prediction_cost_app" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time `json:"time"           required:"false" column:"tag"   type:"time"`
	AppName        string     `json:"appname"        required:"true"  column:"tag"   type:"string"`
	NamespaceName  string     `json:"namespacename"  required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"clustername"    required:"true"  column:"tag"   type:"string"`
	Provider       string     `json:"provider"       required:"true"  column:"tag"   type:"string"`
	Type           string     `json:"type"           required:"true"  column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"    required:"true"  column:"tag"   type:"string"`
	CreateTime     int64      `json:"createtime"     required:"false" column:"field" type:"int64"`
	StartTime      int64      `json:"starttime"      required:"false" column:"field" type:"int64"`
	CostPercentage float64    `json:"costpercentage" required:"false" column:"field" type:"float64"`
	WorkloadCost   float64    `json:"workloadcost"   required:"false" column:"field" type:"float64"`
}

type FedemeterResourcePredictionCostNamespace struct {
	DatahubEntity  `scope:"fedemeter" category:"resource_prediction" type:"cost"`
	Metadata       *Metadata  `measurement:"resource_prediction_cost_namespace" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"true"`
	Time           *time.Time `json:"time"           required:"false" column:"tag"   type:"time"`
	NamespaceName  string     `json:"namespacename"  required:"true"  column:"tag"   type:"string"`
	ClusterName    string     `json:"clustername"    required:"true"  column:"tag"   type:"string"`
	Provider       string     `json:"provider"       required:"true"  column:"tag"   type:"string"`
	Type           string     `json:"type"           required:"true"  column:"tag"   type:"string"`
	Granularity    string     `json:"granularity"    required:"true"  column:"tag"   type:"string"`
	CreateTime     int64      `json:"createtime"     required:"false" column:"field" type:"int64"`
	StartTime      int64      `json:"starttime"      required:"false" column:"field" type:"int64"`
	CostPercentage float64    `json:"costpercentage" required:"false" column:"field" type:"float64"`
	WorkloadCost   float64    `json:"workloadcost"   required:"false" column:"field" type:"float64"`
}
