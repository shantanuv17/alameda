package machineset

type config struct {
	machineSet machineSetConfig
}

type machineSetConfig struct {
	machineSet actionSchema
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
	machineSet: machineSetConfig{
		machineSet: actionSchema{
			create: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_RESOURCE",
					category: "cluster_autoscaler",
					type_:    "machineset",
				},
				measurements: []measurement{
					{
						name: "cluster_autoscaler_machineset",
					},
				},
			},
			list: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_RESOURCE",
					category: "cluster_autoscaler",
					type_:    "machineset",
				},
				measurements: []measurement{
					{
						name: "cluster_autoscaler_machineset",
					},
				},
			},
			delete: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_RESOURCE",
					category: "cluster_autoscaler",
					type_:    "machineset",
				},
				measurements: []measurement{
					{
						name: "cluster_autoscaler_machineset",
					},
				},
			},
		},
	},
}
