package kafka

type config struct {
	kafka kafkaConfig
}

type kafkaConfig struct {
	topic         actionSchema
	consumerGroup actionSchema
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
	kafka: kafkaConfig{
		topic: actionSchema{
			create: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "kafka",
					type_:    "topic",
				},
				measurements: []measurement{
					measurement{
						name: "kafka_topic",
					},
				},
			},
			list: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "kafka",
					type_:    "topic",
				},
				measurements: []measurement{
					measurement{
						name: "kafka_topic",
					},
				},
			},
			delete: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "kafka",
					type_:    "topic",
				},
				measurements: []measurement{
					measurement{
						name: "kafka_topic",
					},
				},
			},
		},
		consumerGroup: actionSchema{
			create: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "kafka",
					type_:    "consumer_group",
				},
				measurements: []measurement{
					measurement{
						name: "kafka_consumer_group",
					},
				},
			},
			list: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "kafka",
					type_:    "consumer_group",
				},
				measurements: []measurement{
					measurement{
						name: "kafka_consumer_group",
					},
				},
			},
			delete: schema{
				schemaMeta: schemaMeta{
					scope:    "SCOPE_APPLICATION",
					category: "kafka",
					type_:    "consumer_group",
				},
				measurements: []measurement{
					measurement{
						name: "kafka_consumer_group",
					},
				},
			},
		},
	},
}
