package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestSetResourceRequirements(t *testing.T) {
	type testCase struct {
		have corev1.ResourceRequirements
		want KubernetesMeta
	}

	cpuLimit := resource.MustParse("200m")
	cpuRequest := resource.MustParse("100m")
	memLimit := resource.MustParse("2M")
	memRequest := resource.MustParse("1M")

	testCases := []testCase{
		{
			have: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceCPU:    cpuLimit,
					corev1.ResourceMemory: memLimit,
				},
				Requests: corev1.ResourceList{
					corev1.ResourceCPU:    cpuRequest,
					corev1.ResourceMemory: memRequest,
				},
			},
			want: KubernetesMeta{
				CPULimit:      "200",
				CPURequest:    "100",
				MemoryLimit:   "2000000",
				MemoryRequest: "1000000",
			},
		},
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		actual := KubernetesMeta{}
		actual.SetResourceRequirements(testCase.have)
		assert.Equal(testCase.want, actual)
	}
}
