package machinegroup

type config struct {
	machineGroup machineGroupConfig
}

type machineGroupConfig struct {
	machineGroup actionSchema
}

type actionSchema struct {
	create schema
	list   schema
	delete schema
}

type schema struct {
	schemaMeta   schemaMeta
	measurements []measurement
}

type schemaMeta struct {
	scope    string
	category string
	type_    string
}

type measurement struct {
	name string
}

// TODO
var defaultConfig = config{
	machineGroup: machineGroupConfig{
		machineGroup: actionSchema{
			create: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_RESOURCE",
					category: "cluster_autoscaler",
					type_:    "machinegroup",
				},
				measurements: []measurement{
					{
						name: "cluster_autoscaler_machinegroup",
					},
				},
			},
			list: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_RESOURCE",
					category: "cluster_autoscaler",
					type_:    "machinegroup",
				},
				measurements: []measurement{
					{
						name: "cluster_autoscaler_machinegroup",
					},
				},
			},
			delete: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_RESOURCE",
					category: "cluster_autoscaler",
					type_:    "machinegroup",
				},
				measurements: []measurement{
					{
						name: "cluster_autoscaler_machinegroup",
					},
				},
			},
		},
	},
}
