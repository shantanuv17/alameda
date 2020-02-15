package controllers

import (
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	// "github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka/mock"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	"github.com/containers-ai/alameda/operator/pkg/kafka"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// func TestListConsumerGroupDetails(t *testing.T) {
// 	testEnv, mgr, err := createTestEnvAndManager()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	defer stopTestEnv(testEnv)

// 	var (
// 		vTrue          = true
// 		topic2         = "topic2"
// 		alamedaScaler1 = autoscalingv1alpha1.AlamedaScaler{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Namespace: "test",
// 				Name:      "test",
// 			},
// 			Spec: autoscalingv1alpha1.AlamedaScalerSpec{
// 				EnableExecution: &vTrue,
// 				Policy:          autoscalingv1alpha1.RecommendationPolicySTABLE,
// 				Type:            autoscalingv1alpha1.AlamedaScalerTypeKafka,
// 				Kafka: &autoscalingv1alpha1.KafkaSpec{
// 					ExporterNamespace: "test",
// 					Topics:            []string{"topic1", "topic2", "topic3"},
// 					ConsumerGroups: []autoscalingv1alpha1.KafkaConsumerGroupSpec{
// 						autoscalingv1alpha1.KafkaConsumerGroupSpec{
// 							Name: "cg1",
// 							Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceSpec{
// 								Custom: "cg1-custom-name",
// 								Kubernetes: &autoscalingv1alpha1.KubernetesResourceSpec{
// 									Selector: &metav1.LabelSelector{
// 										MatchLabels: map[string]string{
// 											"cg1": "cg1",
// 										},
// 									},
// 								},
// 							},
// 							MajorTopic: nil,
// 						},
// 						autoscalingv1alpha1.KafkaConsumerGroupSpec{
// 							Name: "cg2",
// 							Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceSpec{
// 								Custom: "cg2-custom-name",
// 								Kubernetes: &autoscalingv1alpha1.KubernetesResourceSpec{
// 									Selector: &metav1.LabelSelector{
// 										MatchLabels: map[string]string{
// 											"cg2": "cg2",
// 										},
// 									},
// 								},
// 							},
// 							MajorTopic: &topic2,
// 						},
// 						autoscalingv1alpha1.KafkaConsumerGroupSpec{
// 							Name: "cg7",
// 							Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceSpec{
// 								Custom: "cg7-custom-name",
// 							},
// 							MajorTopic: nil,
// 						},
// 					},
// 				},
// 			},
// 		}
// 		deploymentCG1 = appsv1.Deployment{
// 			TypeMeta: metav1.TypeMeta{
// 				Kind: "Deployment",
// 			},
// 			ObjectMeta: metav1.ObjectMeta{
// 				Namespace: "test",
// 				Name:      "cg1",
// 				Labels: map[string]string{
// 					"cg1": "cg1",
// 				},
// 			},
// 			Spec: appsv1.DeploymentSpec{
// 				Selector: &metav1.LabelSelector{
// 					MatchLabels: map[string]string{
// 						"cg1": "cg1",
// 					},
// 				},
// 				Template: corev1.PodTemplateSpec{
// 					ObjectMeta: metav1.ObjectMeta{
// 						Labels: map[string]string{
// 							"cg1": "cg1",
// 						},
// 					},
// 					Spec: corev1.PodSpec{
// 						Containers: []corev1.Container{
// 							corev1.Container{
// 								Name:  "c1",
// 								Image: "c1",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}
// 		deploymentCG2 = appsv1.Deployment{
// 			TypeMeta: metav1.TypeMeta{
// 				Kind: "Deployment",
// 			},
// 			ObjectMeta: metav1.ObjectMeta{
// 				Namespace: "test",
// 				Name:      "cg2",
// 				Labels: map[string]string{
// 					"cg2": "cg2",
// 				},
// 			},
// 			Spec: appsv1.DeploymentSpec{
// 				Selector: &metav1.LabelSelector{
// 					MatchLabels: map[string]string{
// 						"cg2": "cg2",
// 					},
// 				},
// 				Template: corev1.PodTemplateSpec{
// 					ObjectMeta: metav1.ObjectMeta{
// 						Labels: map[string]string{
// 							"cg2": "cg2",
// 						},
// 					},
// 					Spec: corev1.PodSpec{
// 						Containers: []corev1.Container{
// 							corev1.Container{
// 								Name:  "c2",
// 								Image: "c2",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}
// 	)
// 	items := []runtime.Object{
// 		&alamedaScaler1,
// 		&deploymentCG1,
// 		&deploymentCG2,
// 	}
// 	createKubernetesResources(mgr.GetClient(), items...)

// 	type testCaseHave struct {
// 		alamedaScaler                 autoscalingv1alpha1.AlamedaScaler
// 		consumerGroupConsumeTopicsMap map[string][]string
// 	}
// 	type testCase struct {
// 		have testCaseHave
// 		want []kafka.ConsumerGroup
// 	}
// 	testCases := []testCase{
// 		testCase{
// 			have: testCaseHave{
// 				alamedaScaler: alamedaScaler1,
// 				consumerGroupConsumeTopicsMap: map[string][]string{
// 					"cg1": []string{"topic1", "topic2", "topic3"},
// 					"cg2": []string{"topic1", "topic2", "topic3"},
// 				},
// 			},
// 			want: []kafka.ConsumerGroup{
// 				kafka.ConsumerGroup{
// 					Name: "cg1",
// 					ResourceMeta: kafka.ResourceMeta{
// 						CustomName: "cg1-custom-name",
// 						KubernetesMeta: kafka.KubernetesMeta{
// 							Namespace: deploymentCG1.GetNamespace(),
// 							Name:      deploymentCG1.GetName(),
// 							Kind:      "Deployment",
// 						},
// 					},
// 					ConsumeTopic:      "topic1",
// 					ExporterNamespace: "test",
// 					Policy:            datahubresources.RecommendationPolicy_STABLE.String(),
// 					EnableExecution:   true,
// 					AlamedaScalerName: "test",
// 				},
// 				kafka.ConsumerGroup{
// 					Name: "cg2",
// 					ResourceMeta: kafka.ResourceMeta{
// 						CustomName: "cg2-custom-name",
// 						KubernetesMeta: kafka.KubernetesMeta{
// 							Namespace: deploymentCG2.GetNamespace(),
// 							Name:      deploymentCG2.GetName(),
// 							Kind:      "Deployment",
// 						},
// 					},
// 					ConsumeTopic:      "topic2",
// 					ExporterNamespace: "test",
// 					Policy:            datahubresources.RecommendationPolicy_STABLE.String(),
// 					EnableExecution:   true,
// 					AlamedaScalerName: "test",
// 				},
// 				kafka.ConsumerGroup{
// 					Name: "cg7",
// 					ResourceMeta: kafka.ResourceMeta{
// 						CustomName:     "cg7-custom-name",
// 						KubernetesMeta: kafka.KubernetesMeta{},
// 					},
// 					ConsumeTopic:      "",
// 					ExporterNamespace: "test",
// 					Policy:            datahubresources.RecommendationPolicy_STABLE.String(),
// 					EnableExecution:   true,
// 					AlamedaScalerName: "test",
// 				},
// 			},
// 		},
// 	}
// 	reconciler := AlamedaScalerKafkaReconciler{
// 		K8SClient: mgr.GetClient(),
// 		Scheme:    mgr.GetScheme(),
// 	}
// 	ctx := context.Background()
// 	assert := assert.New(t)
// 	for _, testCase := range testCases {
// 		actual, err := reconciler.listConsumerGroupDetails(ctx, testCase.have.alamedaScaler, testCase.have.consumerGroupConsumeTopicsMap)
// 		assert.NoError(err)
// 		assert.ElementsMatch(testCase.want, actual)
// 	}
// }

// TODO: To pass
func TestGetFirstCreatedMatchedKubernetesMetadata(t *testing.T) {
	testEnv, mgr, err := createTestEnvAndManager()
	if err != nil {
		t.Error(err)
	}
	defer stopTestEnv(testEnv)

	replicas1 := int32(1)
	replicas5 := int32(5)
	var (
		deploymentCG1 = appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test",
				Name:      "cg1",
				Labels: map[string]string{
					"cg1": "cg1",
				},
				Annotations: map[string]string{
					"alamedascalers.autoscaling.containers.ai/type": "kafka",
				},
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &replicas1,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"cg1": "cg1",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"cg1": "cg1",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							corev1.Container{
								Name:  "c1",
								Image: "c1",
							},
						},
					},
				},
			},
		}
		deploymentCG2 = appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test",
				Name:      "cg2",
				Labels: map[string]string{
					"cg2": "cg2",
				},
				Annotations: map[string]string{
					"alamedascalers.autoscaling.containers.ai/type": "kafka",
				},
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &replicas5,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"cg2": "cg2",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"cg2": "cg2",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							corev1.Container{
								Name:  "c2",
								Image: "c2",
							},
						},
					},
				},
			},
		}
		deploymentCG3 = appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test",
				Name:      "cg3",
				Labels: map[string]string{
					"cg3": "cg3",
				},
				Annotations: map[string]string{
					"alamedascalers.autoscaling.containers.ai/type": "kafka",
				},
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &replicas5,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"cg3": "cg3",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"cg3": "cg3",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							corev1.Container{
								Name:  "c3",
								Image: "c3",
							},
						},
					},
				},
			},
		}
		deploymentCG4 = appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test",
				Name:      "cg4",
				Labels: map[string]string{
					"cg4": "cg4",
				},
				Annotations: map[string]string{
					"alamedascalers.autoscaling.containers.ai/type": "kafka",
				},
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &replicas5,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"cg4": "cg4",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"cg4": "cg4",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							corev1.Container{
								Name:  "c3",
								Image: "c3",
							},
						},
					},
				},
			},
		}
		deploymentCG5 = appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test",
				Name:      "cg3-cg4-mix",
				Labels: map[string]string{
					"cg3": "cg3",
					"cg4": "cg4",
				},
				Annotations: map[string]string{
					"alamedascalers.autoscaling.containers.ai/type": "kafka",
				},
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &replicas5,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"cg3": "cg3",
						"cg4": "cg4",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"cg3": "cg3",
							"cg4": "cg4",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							corev1.Container{
								Name:  "cg3-cg4-mix",
								Image: "cg3-cg4-mix",
							},
						},
					},
				},
			},
		}
	)

	items := []runtime.Object{
		&deploymentCG1,
		&deploymentCG2,
		&deploymentCG5,
		&deploymentCG3,
		&deploymentCG4,
	}
	createKubernetesResources(mgr.GetClient(), items...)

	assert := assert.New(t)

	type testCaseHave struct {
		namespace              string
		kubernetesResourceSpec autoscalingv1alpha1.KubernetesResourceSpec
	}
	type testCase struct {
		have testCaseHave
		want kafka.KubernetesMeta
	}
	testCases := []testCase{
		testCase{
			have: testCaseHave{
				namespace: "test",
				kubernetesResourceSpec: autoscalingv1alpha1.KubernetesResourceSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"cg1": "cg1",
						},
					},
				},
			},
			want: kafka.KubernetesMeta{
				Namespace:    "test",
				Name:         "cg1",
				Kind:         "Deployment",
				SpecReplicas: 1,
			},
		},
		testCase{
			have: testCaseHave{
				namespace: "test",
				kubernetesResourceSpec: autoscalingv1alpha1.KubernetesResourceSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"cg2": "cg2",
						},
					},
				},
			},
			want: kafka.KubernetesMeta{
				Namespace:    "test",
				Name:         "cg2",
				Kind:         "Deployment",
				SpecReplicas: 5,
			},
		},
		testCase{
			have: testCaseHave{
				namespace: "test",
				kubernetesResourceSpec: autoscalingv1alpha1.KubernetesResourceSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"cg3": "cg3",
						},
					},
				},
			},
			want: kafka.KubernetesMeta{
				Namespace:    "test",
				Name:         deploymentCG5.GetName(),
				Kind:         "Deployment",
				SpecReplicas: 5,
			},
		},
	}
	reconciler := AlamedaScalerKafkaReconciler{
		K8SClient: mgr.GetClient(),
		Scheme:    mgr.GetScheme(),
	}
	ctx := context.Background()
	for _, testCase := range testCases {
		actual, err := reconciler.getFirstCreatedMatchedKubernetesMetadata(ctx, testCase.have.namespace, testCase.have.kubernetesResourceSpec)
		assert.NoError(err)
		assert.Equal(testCase.want, actual)
	}

}

func TestIsAlamedaScalerNeedToBeReconciled(t *testing.T) {
	testEnv, mgr, err := createTestEnvAndManager()
	if err != nil {
		t.Error(err)
	}
	defer stopTestEnv(testEnv)

	alamedaScaler1 := autoscalingv1alpha1.AlamedaScaler{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test",
			Name:      "test",
		},
		Spec: autoscalingv1alpha1.AlamedaScalerSpec{
			Type: autoscalingv1alpha1.AlamedaScalerTypeKafka,
			Kafka: &autoscalingv1alpha1.KafkaSpec{
				ExporterNamespace: "test",
				Topics:            []string{"topic1", "topic2", "topic3"},
				ConsumerGroups: []autoscalingv1alpha1.KafkaConsumerGroupSpec{
					autoscalingv1alpha1.KafkaConsumerGroupSpec{
						Name: "cg1",
						Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceSpec{
							Custom: "cg1-custom-name",
						},
						MajorTopic: "",
					},
				},
			},
		},
	}
	alamedaScaler2 := autoscalingv1alpha1.AlamedaScaler{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test",
			Name:      "test2",
		},
		Spec: autoscalingv1alpha1.AlamedaScalerSpec{
			Type: autoscalingv1alpha1.AlamedaScalerTypeKafka,
			Kafka: &autoscalingv1alpha1.KafkaSpec{
				ExporterNamespace: "test",
				Topics:            []string{"topic1", "topic2", "topic3"},
				ConsumerGroups: []autoscalingv1alpha1.KafkaConsumerGroupSpec{
					autoscalingv1alpha1.KafkaConsumerGroupSpec{
						Name: "cg1",
						Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceSpec{
							Custom: "cg1-custom-name",
						},
						MajorTopic: "",
					},
				},
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := createKubernetesResources(mgr.GetClient(), &alamedaScaler1, &alamedaScaler2); err != nil {
		panic(err)
	}

	r := AlamedaScalerKafkaReconciler{
		K8SClient: mgr.GetClient(),
		Scheme:    mgr.GetScheme(),
	}

	type testCaseWant struct {
		ok  bool
		err error
	}
	type testCase struct {
		have autoscalingv1alpha1.AlamedaScaler
		want testCaseWant
	}

	testCases := []testCase{
		testCase{
			have: alamedaScaler1,
			want: testCaseWant{
				ok:  true,
				err: nil,
			},
		},
	}
	ctx = context.Background()
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual, err := r.isAlamedaScalerNeedToBeReconciled(ctx, testCase.have)
		assert.Equal(testCase.want.ok, actual)
		assert.NoError(err)
	}

	failedTestCases := []testCase{
		testCase{
			have: alamedaScaler2,
			want: testCaseWant{
				ok: false,
			},
		},
	}
	for _, testCase := range failedTestCases {
		actual, err := r.isAlamedaScalerNeedToBeReconciled(ctx, testCase.have)
		assert.Equal(testCase.want.ok, actual)
		assert.NotNil(err)
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

// func TestListConsumerGroupDetailsFromAlamedaScaler(t *testing.T) {
// 	gmockCtrl := gomock.NewController(t)
// 	defer gmockCtrl.Finish()

// 	mockKafkaClient := mock_kafka.NewMockClient(gmockCtrl)
// 	r := AlamedaScalerKafkaReconciler{
// 		KafkaClient: mockKafkaClient,
// 	}

// 	type testCase struct {
// 	}

// 	testCases := []testCase{}

// 	ctx := context.Background()
// 	assert := assert.New(t)
// 	for _, testCase := range testCases {
// 		r.listConsumerGroupDetailsFromAlamedaScaler()
// 	}
// }

// func TestGetMatchedKubernetesMetadata(t *testing.T) {
// 	gmockCtrl := gomock.NewController(t)
// 	defer gmockCtrl.Finish()

// 	mockKafkaClient := mock_kafka.NewMockClient(gmockCtrl)
// 	r := AlamedaScalerKafkaReconciler{
// 		KafkaClient: mockKafkaClient,
// 	}

// 	type testCase struct {
// 	}

// 	testCases := []testCase{}

// 	ctx := context.Background()
// 	assert := assert.New(t)
// 	for _, testCase := range testCases {
// 		r.getMatchedKubernetesMetadata()
// 	}
// }

func TestGetKafkaStatus(t *testing.T) {
	r := AlamedaScalerKafkaReconciler{}
	type testCaseHave struct {
		alamedaScaler  autoscalingv1alpha1.AlamedaScaler
		consumerGroups []kafka.ConsumerGroup
	}

	type testCase struct {
		have testCaseHave
		want autoscalingv1alpha1.KafkaStatus
	}

	testCases := []testCase{
		testCase{
			have: testCaseHave{
				alamedaScaler: autoscalingv1alpha1.AlamedaScaler{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "test",
						Name:      "test",
					},
					Spec: autoscalingv1alpha1.AlamedaScalerSpec{
						Type: autoscalingv1alpha1.AlamedaScalerTypeKafka,
						Kafka: &autoscalingv1alpha1.KafkaSpec{
							ExporterNamespace: "test",
							Topics:            []string{"topic1", "topic2", "topic3"},
							ConsumerGroups: []autoscalingv1alpha1.KafkaConsumerGroupSpec{
								autoscalingv1alpha1.KafkaConsumerGroupSpec{
									Name: "cg1",
									Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceSpec{
										Custom: "cg1-custom-name",
									},
									MajorTopic: "",
								},
							},
						},
					},
				},
				consumerGroups: []kafka.ConsumerGroup{
					kafka.ConsumerGroup{
						Name: "cg1",
						ResourceMeta: kafka.ResourceMeta{
							CustomName: "aaa",
						},
						ConsumeTopic:      "topic1",
						ExporterNamespace: "test",
					},
					kafka.ConsumerGroup{
						Name: "cg2",
						ResourceMeta: kafka.ResourceMeta{
							KubernetesMeta: kafka.KubernetesMeta{
								Namespace: "test",
								Name:      "test",
								Kind:      "Deployment",
							},
						},
						ConsumeTopic:      "topic1",
						ExporterNamespace: "test",
					},
					kafka.ConsumerGroup{
						Name: "cg3",
						ResourceMeta: kafka.ResourceMeta{
							CustomName: "ccc",
						},
						ConsumeTopic:      "topic2",
						ExporterNamespace: "test",
					},
				},
			},
			want: autoscalingv1alpha1.KafkaStatus{
				Effective:         true,
				Message:           "",
				ExporterNamespace: "test",
				Topics: []string{
					"topic1",
					"topic2",
				},
				ConsumerGroups: []autoscalingv1alpha1.KafkaConsumerGroupStatus{
					autoscalingv1alpha1.KafkaConsumerGroupStatus{
						Name:  "cg1",
						Topic: "topic1",
						Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
							CustomName: "aaa",
						},
					},
					autoscalingv1alpha1.KafkaConsumerGroupStatus{
						Name:  "cg2",
						Topic: "topic1",
						Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
							Kubernetes: &autoscalingv1alpha1.KubernetesObjectMetadata{
								Namespace: "test",
								Name:      "test",
								Kind:      "Deployment",
							},
						},
					},
					autoscalingv1alpha1.KafkaConsumerGroupStatus{
						Name:  "cg3",
						Topic: "topic2",
						Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
							CustomName: "ccc",
						},
					},
				},
			},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		actual := r.getKafkaStatus(testCase.have.alamedaScaler, testCase.have.consumerGroups)
		assert.Equal(testCase.want, actual)
	}
}

func TestGetKafkaConsumerGroupStatusFromConsumerGroupDetail(t *testing.T) {
	r := AlamedaScalerKafkaReconciler{}

	type testCase struct {
		have kafka.ConsumerGroup
		want autoscalingv1alpha1.KafkaConsumerGroupStatus
	}

	testCases := []testCase{
		testCase{
			have: kafka.ConsumerGroup{
				Name: "cg1",
				ResourceMeta: kafka.ResourceMeta{
					CustomName: "aaa",
				},
				ConsumeTopic: "topic1",
			},
			want: autoscalingv1alpha1.KafkaConsumerGroupStatus{
				Name:  "cg1",
				Topic: "topic1",
				Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
					CustomName: "aaa",
				},
			},
		},
		testCase{
			have: kafka.ConsumerGroup{
				Name: "cg2",
				ResourceMeta: kafka.ResourceMeta{
					KubernetesMeta: kafka.KubernetesMeta{
						Namespace: "test",
						Name:      "test",
						Kind:      "Deployment",
					},
				},
				ConsumeTopic: "topic2",
			},
			want: autoscalingv1alpha1.KafkaConsumerGroupStatus{
				Name:  "cg2",
				Topic: "topic2",
				Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
					Kubernetes: &autoscalingv1alpha1.KubernetesObjectMetadata{
						Namespace: "test",
						Name:      "test",
						Kind:      "Deployment",
					},
				},
			},
		},
		testCase{
			have: kafka.ConsumerGroup{
				Name: "cg3",
				ResourceMeta: kafka.ResourceMeta{
					CustomName: "ccc",
				},
				ConsumeTopic: "",
			},
			want: autoscalingv1alpha1.KafkaConsumerGroupStatus{
				Name:  "cg3",
				Topic: "",
				Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
					CustomName: "ccc",
				},
			},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		actual := r.getKafkaConsumerGroupStatusFromConsumerGroupDetail(testCase.have)
		assert.Equal(testCase.want, actual)
	}
}

// TODO: Find a way to seperate from integration test from unit test
// Integration test
// func TestIsMetricsExist(t *testing.T) {
// 	metrics := []string{}

// 	cfg := prometheus.Config{
// 		URL: "http://localhost:9090",
// 	}
// 	prom, err := prometheus.NewClient(&cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	reconile := AlamedaScalerKafkaReconciler{
// 		NeededMetrics:    metrics,
// 		PrometheusClient: *prom,
// 	}

// 	type testCase struct {
// 		have []string
// 		want bool
// 	}

// 	testCases := []testCase{
// 		testCase{
// 			have: metrics,
// 			want: true,
// 		},
// 		testCase{
// 			have: []string{"m1"},
// 			want: false,
// 		},
// 		testCase{
// 			have: []string{"m1", "m2"},
// 			want: false,
// 		},
// 	}
// 	assert := assert.New(t)
// 	for _, testCase := range testCases {
// 		actual, missingMetrics, err := reconile.PrometheusClient.IsMetricsExist(context.TODO(), testCase.have)
// 		assert.NoError(err)
// 		assert.Equal(testCase.want, actual)
// 		if !actual {
// 			t.Logf("missing metrics: %+v", missingMetrics)
// 		}
// 	}
// }
