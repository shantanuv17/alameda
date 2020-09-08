package entities

import (
	"time"
)

type ConfigTenancyTenant struct {
	DatahubEntity `scope:"config" category:"tenancy" type:"tenant"`
	Measurement   *Measurement `name:"tenancy_tenant" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time          *time.Time   `json:"time"  required:"false" column:"tag"`
	Name          string       `json:"name"  required:"true"  column:"tag"`
	Dummy         string       `json:"dummy" required:"true"  column:"field"`
}

type ConfigTenancyOrganization struct {
	DatahubEntity `scope:"config" category:"tenancy" type:"organization"`
	Measurement   *Measurement `name:"tenancy_organization" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time          *time.Time   `json:"time"   required:"false" column:"tag"`
	Name          string       `json:"name"   required:"true"  column:"tag"`
	Tenant        string       `json:"tenant" required:"true"  column:"tag"`
	Dummy         string       `json:"dummy"  required:"true"  column:"field"`
}

type ConfigTenancyCluster struct {
	DatahubEntity             `scope:"config" category:"tenancy" type:"cluster"`
	Measurement               *Measurement   `name:"tenancy_cluster" metric:"undefined" boundary:"undefined" quota:"undefined" ts:"false"`
	Time                      *time.Time     `json:"time"                        required:"false" column:"tag"`
	Name                      string         `json:"name"                        required:"true"  column:"tag"`
	Organization              string         `json:"organization"                required:"true"  column:"tag"`
	Tenant                    string         `json:"tenant"                      required:"true"  column:"tag"`
	DataSource                DataSource     `json:"data_source"                 required:"true"  column:"field"`
	DataSourceAddress         string         `json:"data_source_address"         required:"true"  column:"field"`
	DataSourceAccount         string         `json:"data_source_account"         required:"true"  column:"field"`
	DataSourceKeys            string         `json:"data_source_keys"            required:"false" column:"field"`
	DataSourceKeysFunctions   string         `json:"data_source_keys_functions"  required:"false" column:"field"`
	ResourcePlanning          bool           `json:"resource_planning"           required:"false" column:"field"`
	ResourcePlanningMode      DataStoredMode `json:"resource_planning_mode"      required:"false" column:"field"`
	CostAnalysis              bool           `json:"cost_analysis"               required:"false" column:"field"`
	CostAnalysisMode          DataStoredMode `json:"cost_analysis_mode"          required:"false" column:"field"`
	WatchedNamespaces         string         `json:"watched_namespaces"          required:"true"  column:"field"`
	WatchedNamespacesOperator LogicOperator  `json:"watched_namespaces_operator" required:"true"  column:"field"`
	RawSpec                   string         `json:"raw_spec"                    required:"true"  column:"field"`
}
