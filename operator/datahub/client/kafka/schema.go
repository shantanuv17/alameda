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
	name    string
	columns []string
}

type listSchema struct {
	schemaMeta   schemaMeta
	measurements []measurement
}

type listMeasurement struct {
	name       string
	conditions []condition
}

type condition struct {
	keys []string
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
						columns: []string{
							"name",
							"namespace",
							"cluster_name",
							"alameda_scaler_name",
						},
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
						columns: []string{
							"name",
							"namespace",
							"cluster_name",
							"alameda_scaler_name",
						},
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
						columns: []string{
							"name",
							"namespace",
							"cluster_name",
							"topic_name",
							"resource_k8s_namespace",
							"resource_k8s_name",
							"resource_k8s_kind",
							"resource_k8s_replicas",
							"resource_k8s_spec_replicas",
							"resource_k8s_min_replicas",
							"resource_k8s_max_replicas",
							"resource_custom_name",
							"policy",
							"enable_execution",
							"alameda_scaler_name",
						},
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
						columns: []string{
							"name",
							"namespace",
							"cluster_name",
							"topic_name",
							"resource_k8s_namespace",
							"resource_k8s_name",
							"resource_k8s_kind",
							"resource_k8s_replicas",
							"resource_k8s_spec_replicas",
							"resource_custom_name",
							"policy",
							"enable_execution",
							"alameda_scaler_name",
						},
					},
				},
			},
		},
	},
}
