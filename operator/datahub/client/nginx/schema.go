package nginx

type config struct {
	nginx nginxConfig
}

type nginxConfig struct {
	nginx actionSchema
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
	nginx: nginxConfig{
		nginx: actionSchema{
			create: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "nginx",
					type_:    "nginx",
				},
				measurements: []measurement{
					{
						name: "nginx",
					},
				},
			},
			list: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "nginx",
					type_:    "nginx",
				},
				measurements: []measurement{
					{
						name: "nginx",
					},
				},
			},
			delete: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "nginx",
					type_:    "nginx",
				},
				measurements: []measurement{
					{
						name: "nginx",
					},
				},
			},
		},
	},
}
