package autoscaling

import (
	"testing"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

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
		actual := GetTotalResourceFromContainers(testCase.have)
		for resourceName, quantity := range testCase.want.Limits {
			assert.Equal(quantity.Cmp(actual.Limits[resourceName]), 0)
		}
		for resourceName, quantity := range testCase.want.Requests {
			assert.Equal(quantity.Cmp(actual.Requests[resourceName]), 0)
		}
	}
}
