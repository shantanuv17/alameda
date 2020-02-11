package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containers-ai/alameda/operator/datahub/client/kafka/entity"
	"github.com/containers-ai/alameda/operator/pkg/kafka"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func TestNewWriteDataRequestByTopics(t *testing.T) {

	type testCase struct {
		have []kafka.Topic
		want data.WriteDataRequest
	}

	testCases := []testCase{
		testCase{
			have: []kafka.Topic{
				kafka.Topic{
					Name:              "topic1",
					ExporterNamespace: "exporter-namespace-1",
					ClusterName:       "cluster-name-1",
					AlamedaScalerName: "alamedascaler-namespace-1",
				},
				kafka.Topic{
					Name:              "topic2",
					ExporterNamespace: "exporter-namespace-2",
					ClusterName:       "cluster-name-2",
					AlamedaScalerName: "alamedascaler-namespace-2",
				},
			},
			want: data.WriteDataRequest{
				SchemaMeta: &schemas.SchemaMeta{
					Scope:    schemas.Scope_SCOPE_APPLICATION,
					Category: "kafka",
					Type:     "topic",
				},
				WriteData: []*data.WriteData{
					&data.WriteData{
						Measurement: "kafka_topic",
						Columns: []string{
							"name",
							"namespace",
							"cluster_name",
							"alameda_scaler_name",
						},
						Rows: []*common.Row{
							&common.Row{
								Values: []string{
									"topic1",
									"exporter-namespace-1",
									"cluster-name-1",
									"alamedascaler-namespace-1",
								},
							},
							&common.Row{
								Values: []string{
									"topic2",
									"exporter-namespace-2",
									"cluster-name-2",
									"alamedascaler-namespace-2",
								},
							},
						},
					},
				},
			},
		},
	}
	kr := KafkaRepository{
		schemaConfig: defaultConfig,
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := kr.newWriteDataRequestByTopics(testCase.have)
		assert.NoError(err)
		assert.Equal(testCase.want, actual)
	}
}

func TestNewWriteDataRequestByConsumerGroups(t *testing.T) {

	type testCase struct {
		have []kafka.ConsumerGroup
		want data.WriteDataRequest
	}

	testCases := []testCase{
		testCase{
			have: []kafka.ConsumerGroup{
				kafka.ConsumerGroup{
					Name:              "consumer-group-1",
					ExporterNamespace: "exporter-namespace-1",
					ClusterName:       "cluster-name-1",
					AlamedaScalerName: "alamedascaler-namespace-1",
					Policy:            "stable",
					EnableExecution:   true,
					ConsumeTopic:      "topic1",
					ResourceMeta: kafka.ResourceMeta{
						KubernetesMeta: kafka.KubernetesMeta{
							Namespace:     "namespace-1",
							Name:          "deployment-1",
							Kind:          "Deployment",
							ReadyReplicas: 1,
							SpecReplicas:  2,
						},
						CustomName: "custom-name-1",
					},
					MinReplicas: 1,
					MaxReplicas: 10,
				},
				kafka.ConsumerGroup{
					Name:              "consumer-group-2",
					ExporterNamespace: "exporter-namespace-2",
					ClusterName:       "cluster-name-2",
					AlamedaScalerName: "alamedascaler-namespace-2",
					Policy:            "stable",
					EnableExecution:   true,
					ConsumeTopic:      "topic2",
					ResourceMeta: kafka.ResourceMeta{
						CustomName: "custom-name-2",
					},
				},
			},
			want: data.WriteDataRequest{
				SchemaMeta: &schemas.SchemaMeta{
					Scope:    schemas.Scope_SCOPE_APPLICATION,
					Category: "kafka",
					Type:     "consumer_group",
				},
				WriteData: []*data.WriteData{
					&data.WriteData{
						Measurement: "kafka_consumer_group",
						Columns: []string{
							"name",
							"namespace",
							"cluster_name",
							"alameda_scaler_name",
							"policy",
							"enable_execution",
							"topic_name",
							"resource_custom_name",
							"resource_k8s_namespace",
							"resource_k8s_name",
							"resource_k8s_kind",
							"resource_k8s_replicas",
							"resource_k8s_spec_replicas",
							"resource_k8s_min_replicas",
							"resource_k8s_max_replicas",
							"resource_cpu_limit",
							"resource_cpu_request",
							"resource_memory_limit",
							"resource_memory_request",
							"volumes_size",
							"volumes_pvc_size",
						},
						Rows: []*common.Row{
							&common.Row{
								Values: []string{
									"consumer-group-1",
									"exporter-namespace-1",
									"cluster-name-1",
									"alamedascaler-namespace-1",
									"stable",
									"true",
									"topic1",
									"custom-name-1",
									"namespace-1",
									"deployment-1",
									"Deployment",
									"1",
									"2",
									"1",
									"10",
									"",
									"",
									"",
									"",
									"",
									"",
								},
							},
							&common.Row{
								Values: []string{
									"consumer-group-2",
									"exporter-namespace-2",
									"cluster-name-2",
									"alamedascaler-namespace-2",
									"stable",
									"true",
									"topic2",
									"custom-name-2",
									"",
									"",
									"",
									"0",
									"0",
									"0",
									"0",
									"",
									"",
									"",
									"",
									"",
									"",
								},
							},
						},
					},
				},
			},
		},
	}
	kr := KafkaRepository{
		schemaConfig: defaultConfig,
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := kr.newWriteDataRequesByConsumerGroups(testCase.have)
		assert.NoError(err)
		for i, data := range testCase.want.WriteData {
			assert.Equal(data.Measurement, actual.WriteData[i].Measurement)
			assert.ElementsMatch(data.Columns, actual.WriteData[i].Columns)
			for j, row := range data.Rows {
				assert.ElementsMatch(row.Values, actual.WriteData[i].Rows[j].Values)
			}
		}
	}
}

func TestNewDeleteDataRequestByTopics(t *testing.T) {

	type testCase struct {
		have []kafka.Topic
		want data.DeleteDataRequest
	}

	testCases := []testCase{
		testCase{
			have: []kafka.Topic{
				kafka.Topic{
					Name:              "topic1",
					ExporterNamespace: "exporter-namespace-1",
					ClusterName:       "cluster-name-1",
					AlamedaScalerName: "alamedascaler-namespace-1",
				},
				kafka.Topic{
					ExporterNamespace: "exporter-namespace-2",
					ClusterName:       "cluster-name-2",
					AlamedaScalerName: "alamedascaler-namespace-2",
				},
			},
			want: data.DeleteDataRequest{
				SchemaMeta: &schemas.SchemaMeta{
					Scope:    schemas.Scope_SCOPE_APPLICATION,
					Category: "kafka",
					Type:     "topic",
				},
				DeleteData: []*data.DeleteData{
					&data.DeleteData{
						Measurement: "kafka_topic",
						QueryCondition: &common.QueryCondition{
							WhereCondition: []*common.Condition{
								&common.Condition{
									Keys: []string{
										"name",
										"namespace",
										"cluster_name",
									},
									Values: []string{
										"topic1",
										"exporter-namespace-1",
										"cluster-name-1",
									},
									Operators: []string{
										"=",
										"=",
										"=",
									},
									Types: []common.DataType{
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
									},
								},
								&common.Condition{
									Keys: []string{
										"namespace",
										"cluster_name",
									},
									Values: []string{
										"exporter-namespace-2",
										"cluster-name-2",
									},
									Operators: []string{
										"=",
										"=",
									},
									Types: []common.DataType{
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	kr := KafkaRepository{
		schemaConfig: defaultConfig,
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := kr.newDeleteDataRequestByTopics(testCase.have)
		assert.NoError(err)
		for i, data := range testCase.want.DeleteData {
			for j, cond := range data.QueryCondition.WhereCondition {
				assert.ElementsMatch(cond.Keys, actual.DeleteData[i].QueryCondition.WhereCondition[j].Keys)
				assert.ElementsMatch(cond.Values, actual.DeleteData[i].QueryCondition.WhereCondition[j].Values)
				assert.ElementsMatch(cond.Operators, actual.DeleteData[i].QueryCondition.WhereCondition[j].Operators)
				assert.ElementsMatch(cond.Types, actual.DeleteData[i].QueryCondition.WhereCondition[j].Types)
			}
		}
	}
}

func TestNewDeleteDataRequestByConsumerGroups(t *testing.T) {
	type testCase struct {
		have []kafka.ConsumerGroup
		want data.DeleteDataRequest
	}

	testCases := []testCase{
		testCase{
			have: []kafka.ConsumerGroup{
				kafka.ConsumerGroup{
					Name:              "consumer-group-1",
					ExporterNamespace: "exporter-namespace-1",
					ClusterName:       "cluster-name-1",
					AlamedaScalerName: "alamedascaler-namespace-1",
					Policy:            "stable",
					EnableExecution:   true,
					ConsumeTopic:      "topic1",
					ResourceMeta: kafka.ResourceMeta{
						KubernetesMeta: kafka.KubernetesMeta{
							Namespace:     "namespace-1",
							Name:          "deployment-1",
							Kind:          "Deployment",
							ReadyReplicas: 1,
							SpecReplicas:  2,
						},
						CustomName: "custom-name-1",
					},
				},
				kafka.ConsumerGroup{
					Name:              "consumer-group-2",
					ExporterNamespace: "exporter-namespace-2",
					ClusterName:       "cluster-name-2",
					AlamedaScalerName: "alamedascaler-namespace-2",
					Policy:            "stable",
					EnableExecution:   true,
					ConsumeTopic:      "topic2",
					ResourceMeta: kafka.ResourceMeta{
						CustomName: "custom-name-2",
					},
				},
			},
			want: data.DeleteDataRequest{
				SchemaMeta: &schemas.SchemaMeta{
					Scope:    schemas.Scope_SCOPE_APPLICATION,
					Category: "kafka",
					Type:     "consumer_group",
				},
				DeleteData: []*data.DeleteData{
					&data.DeleteData{
						Measurement: "kafka_consumer_group",
						QueryCondition: &common.QueryCondition{
							WhereCondition: []*common.Condition{
								&common.Condition{
									Keys: []string{
										"name",
										"namespace",
										"cluster_name",
										"topic_name",
									},

									Values: []string{
										"consumer-group-1",
										"exporter-namespace-1",
										"cluster-name-1",
										"topic1",
									},
									Operators: []string{
										"=",
										"=",
										"=",
										"=",
									},
									Types: []common.DataType{
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
									},
								},
								&common.Condition{
									Keys: []string{
										"name",
										"namespace",
										"cluster_name",
										"topic_name",
									},

									Values: []string{
										"consumer-group-2",
										"exporter-namespace-2",
										"cluster-name-2",
										"topic2",
									},
									Operators: []string{
										"=",
										"=",
										"=",
										"=",
									},
									Types: []common.DataType{
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	kr := KafkaRepository{
		schemaConfig: defaultConfig,
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := kr.newDeleteDataRequestByConsumerGroups(testCase.have)
		assert.NoError(err)
		for i, data := range testCase.want.DeleteData {
			for j, cond := range data.QueryCondition.WhereCondition {
				assert.ElementsMatch(cond.Keys, actual.DeleteData[i].QueryCondition.WhereCondition[j].Keys)
				assert.ElementsMatch(cond.Values, actual.DeleteData[i].QueryCondition.WhereCondition[j].Values)
				assert.ElementsMatch(cond.Operators, actual.DeleteData[i].QueryCondition.WhereCondition[j].Operators)
				assert.ElementsMatch(cond.Types, actual.DeleteData[i].QueryCondition.WhereCondition[j].Types)
			}
		}
	}
}

func TestNewReadDataRequestForTopics(t *testing.T) {
	type testCase struct {
		have ListTopicsOption
		want data.ReadDataRequest
	}

	testCases := []testCase{
		testCase{
			have: ListTopicsOption{
				ClusterName:       "test-cluster",
				ExporterNamespace: "test-namespace",
				AlamedaScalerName: "test-alameda-scaler-name",
			},
			want: data.ReadDataRequest{
				SchemaMeta: &schemas.SchemaMeta{
					Scope:    schemas.Scope_SCOPE_APPLICATION,
					Category: "kafka",
					Type:     "topic",
				},
				ReadData: []*data.ReadData{
					&data.ReadData{
						Measurement: "kafka_topic",
						QueryCondition: &common.QueryCondition{
							WhereCondition: []*common.Condition{
								&common.Condition{
									Keys: []string{
										"cluster_name",
										"namespace",
										"alameda_scaler_name",
									},
									Values: []string{
										"test-cluster",
										"test-namespace",
										"test-alameda-scaler-name",
									},
									Operators: []string{
										"=",
										"=",
										"=",
									},
									Types: []common.DataType{
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	r := KafkaRepository{
		schemaConfig: defaultConfig,
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := r.newReadDataRequestForTopics(testCase.have)
		assert.NoError(err)
		assert.Equal(testCase.want, actual)
	}
}

func TestNewReadDataRequestForConsumerGroups(t *testing.T) {
	type testCase struct {
		have ListConsumerGroupsOption
		want data.ReadDataRequest
	}

	testCases := []testCase{
		testCase{
			have: ListConsumerGroupsOption{
				ClusterName:       "test-cluster",
				ExporterNamespace: "test-namespace",
				AlamedaScalerName: "test-alameda-scaler-name",
			},
			want: data.ReadDataRequest{
				SchemaMeta: &schemas.SchemaMeta{
					Scope:    schemas.Scope_SCOPE_APPLICATION,
					Category: "kafka",
					Type:     "consumer_group",
				},
				ReadData: []*data.ReadData{
					&data.ReadData{
						Measurement: "kafka_consumer_group",
						QueryCondition: &common.QueryCondition{
							WhereCondition: []*common.Condition{
								&common.Condition{
									Keys: []string{
										"cluster_name",
										"namespace",
										"alameda_scaler_name",
									},
									Values: []string{
										"test-cluster",
										"test-namespace",
										"test-alameda-scaler-name",
									},
									Operators: []string{
										"=",
										"=",
										"=",
									},
									Types: []common.DataType{
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
										common.DataType_DATATYPE_STRING,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	r := KafkaRepository{
		schemaConfig: defaultConfig,
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := r.newReadDataRequestForConsumerGroups(testCase.have)
		assert.NoError(err)
		assert.Equal(testCase.want, actual)
	}
}

func TestDecodeSlice(t *testing.T) {
	// type testCaseHave struct {
	// 	data  data.Data
	// 	items []entity.KafkaTopic
	// }

	// type testCase struct {
	// 	have testCaseHave
	// 	want []entity.KafkaTopic
	// }

	// testCases := []testCase{
	// 	testCase{
	// 		have: testCaseHave{
	// 			data: data.Data{
	// 				SchemaMeta: &schemas.SchemaMeta{
	// 					Scope:    schemas.Scope_SCOPE_APPLICATION,
	// 					Category: "kafka",
	// 					Type:     "topic",
	// 				},
	// 				Rawdata: []*data.Rawdata{
	// 					&data.Rawdata{
	// 						Measurement: "kafka_topic",
	// 						Groups: []*common.Group{
	// 							&common.Group{
	// 								Columns: []string{
	// 									"name",
	// 									"namespace",
	// 									"cluster_name",
	// 									"alameda_scaler_name",
	// 								},
	// 								Rows: []*common.Row{
	// 									&common.Row{
	// 										Values: []string{
	// 											"topic1",
	// 											"exporter-namespace-1",
	// 											"cluster-name-1",
	// 											"alamedascaler-namespace-1",
	// 										},
	// 									},
	// 									&common.Row{
	// 										Values: []string{
	// 											"topic2",
	// 											"exporter-namespace-2",
	// 											"cluster-name-2",
	// 											"alamedascaler-namespace-2",
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			items: make([]entity.KafkaTopic, 0),
	// 		},
	// 		want: []entity.KafkaTopic{
	// 			entity.KafkaTopic{
	// 				Name:              "topic1",
	// 				ExporterNamespace: "exporter-namespace-1",
	// 				ClusterName:       "cluster-name-1",
	// 				AlamedaScalerName: "alamedascaler-namespace-1",
	// 			},
	// 			entity.KafkaTopic{
	// 				Name:              "topic2",
	// 				ExporterNamespace: "exporter-namespace-2",
	// 				ClusterName:       "cluster-name-2",
	// 				AlamedaScalerName: "alamedascaler-namespace-2",
	// 			},
	// 		},
	// 	},
	// }

	type testCaseHave struct {
		data  data.Data
		items []entity.KafkaConsumerGroup
	}

	type testCase struct {
		have testCaseHave
		want []entity.KafkaConsumerGroup
	}

	testCases := []testCase{
		testCase{
			have: testCaseHave{
				data: data.Data{
					SchemaMeta: &schemas.SchemaMeta{
						Scope:    schemas.Scope_SCOPE_APPLICATION,
						Category: "kafka",
						Type:     "consumer_group",
					},
					Rawdata: []*data.Rawdata{
						&data.Rawdata{
							Measurement: "kafka_consumer_group",
							Groups: []*common.Group{
								&common.Group{
									Columns: []string{
										"alameda_scaler_name",
										"cluster_name",
										"enable_execution",
										"name",
										"namespace",
										"policy",
										"resource_cpu_limit",
										"resource_cpu_request",
										"resource_custom_name",
										"resource_k8s_kind",
										"resource_k8s_max_replicas",
										"resource_k8s_min_replicas",
										"resource_k8s_name",
										"resource_k8s_namespace",
										"resource_k8s_replicas",
										"resource_k8s_spec_replicas",
										"resource_memory_limit",
										"resource_memory_request",
										"topic_name",
										"volumes_pvc_size",
										"volumes_size",
									},
									Rows: []*common.Row{
										&common.Row{
											Values: []string{
												"federatorai",
												"e4bd7717-d9f4-11e9-b897-005056bcf4ef",
												"true",
												"consumergroup-1",
												"kafka",
												"STABLE",
												"0",
												"50",
												"",
												"Deployment",
												"3",
												"1",
												"alameda-datahub",
												"federatorai",
												"1",
												"1",
												"0",
												"104857600",
												"topic2022",
												"",
												"",
											},
										},
										&common.Row{
											Values: []string{
												"federatorai",
												"e4bd7717-d9f4-11e9-b897-005056bcf4ef",
												"true",
												"consumergroup-3",
												"kafka",
												"STABLE",
												"",
												"",
												"test-custom-name",
												"",
												"1",
												"1",
												"",
												"",
												"0",
												"0",
												"",
												"",
												"topic2020",
												"",
												"",
											},
										},
									},
								},
							},
						},
					},
				},
				items: make([]entity.KafkaConsumerGroup, 0),
			},
			want: []entity.KafkaConsumerGroup{
				entity.KafkaConsumerGroup{
					AlamedaScalerName:         "federatorai",
					ClusterName:               "e4bd7717-d9f4-11e9-b897-005056bcf4ef",
					EnableExecution:           true,
					Name:                      "consumergroup-1",
					ExporterNamespace:         "kafka",
					Policy:                    "STABLE",
					ResourceK8SCPULimit:       "0",
					ResourceK8SCPURequest:     "50",
					ResourceCustomName:        "",
					ResourceK8SKind:           "Deployment",
					MaxReplicas:               3,
					MinReplicas:               1,
					ResourceK8SName:           "alameda-datahub",
					ResourceK8SNamespace:      "federatorai",
					ReadyReplicas:             1,
					SpecReplicas:              1,
					ResourceK8SMemoryLimit:    "0",
					ResourceK8SMemoryRequest:  "104857600",
					ConsumeTopic:              "topic2022",
					ResourceK8SVolumesPVCSize: "",
					ResourceK8SVolumesSize:    "",
				},
				entity.KafkaConsumerGroup{
					AlamedaScalerName:         "federatorai",
					ClusterName:               "e4bd7717-d9f4-11e9-b897-005056bcf4ef",
					EnableExecution:           true,
					Name:                      "consumergroup-3",
					ExporterNamespace:         "kafka",
					Policy:                    "STABLE",
					ResourceK8SCPULimit:       "",
					ResourceK8SCPURequest:     "",
					ResourceCustomName:        "test-custom-name",
					ResourceK8SKind:           "",
					MaxReplicas:               1,
					MinReplicas:               1,
					ResourceK8SName:           "",
					ResourceK8SNamespace:      "",
					ReadyReplicas:             0,
					SpecReplicas:              0,
					ResourceK8SMemoryLimit:    "",
					ResourceK8SMemoryRequest:  "",
					ConsumeTopic:              "topic2020",
					ResourceK8SVolumesPVCSize: "",
					ResourceK8SVolumesSize:    "",
				},
			},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		err := decodeSlice(testCase.have.data, &testCase.have.items)
		assert.NoError(err)
		assert.ElementsMatch(testCase.have.items, testCase.want)
	}
}
