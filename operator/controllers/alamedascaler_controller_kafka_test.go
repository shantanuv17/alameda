package controllers

import (
	"context"
	"log"
	"path/filepath"
	"testing"

	mock_kafka "github.com/containers-ai/alameda/internal/pkg/message-queue/kafka/mock"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	gomock "github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func createTestEnvAndManager() (*envtest.Environment, manager.Manager, error) {
	//specify testEnv configuration
	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crds")},
	}
	//start testEnv
	cfg, err := testEnv.Start()
	if err != nil {
		return nil, nil, err
	}
	// setup manager and reconciler
	mgr, err := ctrl.NewManager(cfg, ctrl.Options{})
	err = autoscalingv1alpha1.AddToScheme(mgr.GetScheme())
	if err != nil {
		return nil, nil, err
	}
	go func() {
		err := mgr.Start(ctrl.SetupSignalHandler())
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Wait for sync...")
	ok := mgr.GetCache().WaitForCacheSync(nil)
	if !ok {
		return nil, nil, errors.New("Wait for sync failed")
	}
	log.Println("Wait for sync done.")

	return testEnv, mgr, nil
}

func stopTestEnv(testEnv *envtest.Environment) error {
	if err := testEnv.Stop(); err != nil {
		return errors.Wrap(err, "stop test environment failed")
	}
	return nil
}

func createKubernetesResources(k8sClient client.Client, items ...runtime.Object) error {
	for _, item := range items {
		err := k8sClient.Create(context.Background(), item)
		if err != nil {
			return errors.Wrap(err, "create item failed")
		}
	}
	return nil
}

func TestChooseTopic(t *testing.T) {
	assert := assert.New(t)
	type testCaseHave struct {
		majorTopic           string
		wantToConsumeTopics  []string
		currentConsumeTopics []string
	}
	type testCase struct {
		have testCaseHave
		want string
	}
	testCases := []testCase{
		testCase{
			have: testCaseHave{
				majorTopic:           "t1",
				wantToConsumeTopics:  []string{},
				currentConsumeTopics: []string{"t1", "t2"},
			},
			want: "t1",
		},
		testCase{
			have: testCaseHave{
				majorTopic:           "t3",
				wantToConsumeTopics:  []string{"t1"},
				currentConsumeTopics: []string{"t1", "t2"},
			},
			want: "t1",
		},
		testCase{
			have: testCaseHave{
				majorTopic:           "t1",
				wantToConsumeTopics:  []string{},
				currentConsumeTopics: []string{},
			},
			want: "",
		},
		testCase{
			have: testCaseHave{
				majorTopic:           "t1",
				wantToConsumeTopics:  []string{"t2"},
				currentConsumeTopics: []string{},
			},
			want: "",
		},
		testCase{
			have: testCaseHave{
				majorTopic:           "",
				wantToConsumeTopics:  []string{},
				currentConsumeTopics: []string{},
			},
			want: "",
		},
	}
	for _, testCase := range testCases {
		actual := chooseTopic(testCase.have.majorTopic, testCase.have.wantToConsumeTopics, testCase.have.currentConsumeTopics)
		assert.Equal(testCase.want, actual)
	}
}

func TestGetConsumerGroupToConsumeTopicsMap(t *testing.T) {
	gmockCtrl := gomock.NewController(t)
	defer gmockCtrl.Finish()

	mockKafkaClient := mock_kafka.NewMockClient(gmockCtrl)
	r := AlamedaScalerKafkaReconciler{
		KafkaClient: mockKafkaClient,
	}

	type consumeStatus struct {
		consumerGroup string
		consumeTopics []string
	}

	type testCaseHave struct {
		consumerGroupsToCheck []string
	}

	type testCase struct {
		have testCaseHave
		want map[string][]string
	}

	consumeStatuses := []consumeStatus{
		consumeStatus{
			consumerGroup: "cg1",
			consumeTopics: []string{"topic1", "topic2"},
		},
		consumeStatus{
			consumerGroup: "cg2",
			consumeTopics: []string{"topic2", "topic3"},
		},
		consumeStatus{
			consumerGroup: "cg3",
			consumeTopics: []string{},
		},
	}
	for _, status := range consumeStatuses {
		mockKafkaClient.EXPECT().ListConsumeTopics(context.Background(), gomock.Eq(status.consumerGroup)).Return(status.consumeTopics, nil)
	}

	testCases := []testCase{
		testCase{
			have: testCaseHave{
				consumerGroupsToCheck: []string{"cg1", "cg2", "cg3"},
			},
			want: map[string][]string{
				"cg1": []string{"topic1", "topic2"},
				"cg2": []string{"topic2", "topic3"},
				"cg3": []string{},
			},
		},
	}

	ctx := context.Background()
	assert := assert.New(t)
	for _, testCase := range testCases {
		m, err := r.getConsumerGroupToConsumeTopicsMap(ctx, testCase.have.consumerGroupsToCheck)
		assert.NoError(err)
		assert.Equal(testCase.want, m)
	}
}

func TestAlamedaScalerKafkaReconciler_writeKafkaControllers(t *testing.T) {

	type fields struct {
		ClusterUID    string
		DatahubClient *datahubpkg.Client
	}
	type args struct {
		kafkaStatus *autoscalingv1alpha1.KafkaStatus
		scalerName  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				ClusterUID:    "c1",
				DatahubClient: datahubpkg.NewClient("127.0.0.1:50050"),
			},
			args: args{
				kafkaStatus: &autoscalingv1alpha1.KafkaStatus{
					ConsumerGroups: []autoscalingv1alpha1.KafkaConsumerGroupStatus{
						{
							Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
								Kubernetes: &autoscalingv1alpha1.KubernetesObjectMetadata{
									Namespace: "ns",
									Name:      "name",
									Kind:      "Deployment",
								},
							},
						},
					},
				},
				scalerName: "sc1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := AlamedaScalerKafkaReconciler{
				ClusterUID:    tt.fields.ClusterUID,
				DatahubClient: tt.fields.DatahubClient,
			}
			if err := r.writeKafkaControllers(tt.args.kafkaStatus, tt.args.scalerName); (err != nil) != tt.wantErr {
				t.Errorf("AlamedaScalerKafkaReconciler.writeKafkaControllers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
