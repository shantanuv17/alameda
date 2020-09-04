package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type testCaseHave struct {
		src  VolumeCapacity
		dest VolumeCapacity
	}

	type testCase struct {
		have testCaseHave
		want VolumeCapacity
	}

	testCases := []testCase{
		{
			have: testCaseHave{
				src: VolumeCapacity{
					Total: 10,
					PVC:   10,
				},
				dest: VolumeCapacity{
					Total: 20,
					PVC:   20,
				},
			},
			want: VolumeCapacity{
				Total: 30,
				PVC:   30,
			},
		},
	}
	assert := assert.New(t)
	for _, testCase := range testCases {
		testCase.have.dest.Add(testCase.have.src)
		assert.Equal(testCase.have.dest, testCase.want)
	}
}
