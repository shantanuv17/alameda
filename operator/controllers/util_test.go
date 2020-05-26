package controllers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsLabelsSelectedBySelector(t *testing.T) {
	type testCaseHave struct {
		selector metav1.LabelSelector
		labels   map[string]string
	}
	type testCase struct {
		have testCaseHave
		want bool
	}

	testCases := []testCase{
		{
			have: testCaseHave{
				selector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						"k1": "v1",
						"k2": "v2",
					},
				},
				labels: map[string]string{
					"k1": "v1",
					"k2": "v2",
					"k3": "v3",
				},
			},
			want: true,
		},
		{
			have: testCaseHave{
				selector: metav1.LabelSelector{
					MatchLabels: map[string]string{
						"k1": "v1",
						"k2": "v2",
					},
				},
				labels: map[string]string{
					"k1": "v1",
					"k3": "v3",
				},
			},
			want: false,
		},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		actual := isLabelsSelectedBySelector(tc.have.selector, tc.have.labels)
		assert.Equal(tc.want, actual)
	}
}

func TestGetFirstCreatedObjectMeta(t *testing.T) {
	type testCase struct {
		have []metav1.ObjectMeta
		want metav1.ObjectMeta
	}

	now := metav1.NewTime(time.Now())
	now1m := metav1.NewTime(time.Now().Add(1 * time.Minute))
	now2m := metav1.NewTime(time.Now().Add(2 * time.Minute))

	testCases := []testCase{
		{
			have: []metav1.ObjectMeta{
				{
					Namespace:         "test",
					Name:              "o1",
					UID:               "1",
					CreationTimestamp: now,
				},
				{
					Namespace:         "test",
					Name:              "o2",
					UID:               "2",
					CreationTimestamp: now1m,
				},
				{
					Namespace:         "test",
					Name:              "o3",
					UID:               "3",
					CreationTimestamp: now2m,
				},
			},
			want: metav1.ObjectMeta{
				Namespace:         "test",
				Name:              "o1",
				UID:               "1",
				CreationTimestamp: now,
			},
		},
		{
			have: []metav1.ObjectMeta{
				{
					Namespace:         "test",
					Name:              "o1",
					UID:               "1",
					CreationTimestamp: now2m,
				},
				{
					Namespace:         "test",
					Name:              "o2",
					UID:               "2",
					CreationTimestamp: now1m,
				},
				{
					Namespace:         "test",
					Name:              "o3",
					UID:               "3",
					CreationTimestamp: now,
				},
			},
			want: metav1.ObjectMeta{
				Namespace:         "test",
				Name:              "o3",
				UID:               "3",
				CreationTimestamp: now,
			},
		},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		actual := getFirstCreatedObjectMeta(tc.have)
		assert.Equal(tc.want, actual)
	}
}

func TestGetTotalResourceFromContainers(t *testing.T) {
	type testCase struct {
		have []corev1.Container
		want corev1.ResourceRequirements
	}

	cpu100 := resource.MustParse("100m")
	cpu200 := resource.MustParse("200m")
	cpu300 := resource.MustParse("300m")
	cpu400 := resource.MustParse("400m")

	mem1M := resource.MustParse("1Mi")
	mem2M := resource.MustParse("2Mi")
	mem3M := resource.MustParse("3Mi")
	mem4M := resource.MustParse("4Mi")

	testCases := []testCase{
		{
			have: []corev1.Container{
				{
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    cpu100,
							corev1.ResourceMemory: mem1M,
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    cpu100,
							corev1.ResourceMemory: mem1M,
						},
					},
				},
				{
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    cpu200,
							corev1.ResourceMemory: mem2M,
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    cpu300,
							corev1.ResourceMemory: mem3M,
						},
					},
				},
			},
			want: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceCPU:    cpu300,
					corev1.ResourceMemory: mem3M,
				},
				Requests: corev1.ResourceList{
					corev1.ResourceCPU:    cpu400,
					corev1.ResourceMemory: mem4M,
				},
			},
		},
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual := getTotalResourceFromContainers(testCase.have)
		for resourceName, quantity := range testCase.want.Limits {
			assert.Equal(quantity.Cmp(actual.Limits[resourceName]), 0)
		}
		for resourceName, quantity := range testCase.want.Requests {
			assert.Equal(quantity.Cmp(actual.Requests[resourceName]), 0)
		}
	}
}
