package controllers

import (
	"context"
	"testing"
	"time"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestIsWorkloadControllerCanBeMonitoredByAlamedaScaler(t *testing.T) {

	testEnv, mgr, err := createTestEnvAndManager()
	if err != nil {
		t.Error(err)
	}
	defer stopTestEnv(testEnv)

	now_1m_before := metav1.NewTime(time.Now().Add(-1 * time.Minute))
	now_2m_before := metav1.NewTime(time.Now().Add(-2 * time.Minute))
	var (
		alamedaScaler1 = autoscalingv1alpha1.AlamedaScaler{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:         "test",
				Name:              "as1",
				CreationTimestamp: now_1m_before,
			},
			Spec: autoscalingv1alpha1.AlamedaScalerSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"k1": "v1",
					},
				},
			},
		}
		alamedaScaler2 = autoscalingv1alpha1.AlamedaScaler{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:         "test",
				Name:              "as2",
				CreationTimestamp: now_2m_before,
			},
			Spec: autoscalingv1alpha1.AlamedaScalerSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"k1": "v1",
					},
				},
			},
		}
		deploymentCG1 = appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test",
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			Spec: appsv1.DeploymentSpec{
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
	)

	items := []runtime.Object{
		&alamedaScaler1,
		&alamedaScaler2,
		&deploymentCG1,
	}
	createKubernetesResources(mgr.GetClient(), items...)

	r := AlamedaScalerReconciler{
		Client: mgr.GetClient(),
	}

	type testCaseHave struct {
		ctx                context.Context
		workloadController metav1.ObjectMeta
		alamedaScaler      autoscalingv1alpha1.AlamedaScaler
	}
	type testCase struct {
		have testCaseHave
		want bool
	}
	ctx := context.Background()
	testCases := []testCase{
		testCase{
			have: testCaseHave{
				ctx:                ctx,
				workloadController: deploymentCG1.ObjectMeta,
				alamedaScaler:      alamedaScaler1,
			},
			want: true,
		},
		testCase{
			have: testCaseHave{
				ctx:                ctx,
				workloadController: deploymentCG1.ObjectMeta,
				alamedaScaler:      alamedaScaler2,
			},
			want: false,
		},
	}
	assert := assert.New(t)
	for i, tc := range testCases {
		actual, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(
			context.TODO(),
			tc.have.workloadController,
			tc.have.alamedaScaler,
		)
		assert.NoError(err)
		assert.Equalf(tc.want, actual, "case #%d", i)
	}
}
