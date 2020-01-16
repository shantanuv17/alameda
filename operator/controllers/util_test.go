package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

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
		testCase{
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
		testCase{
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
	now_1m := metav1.NewTime(time.Now().Add(1 * time.Minute))
	now_2m := metav1.NewTime(time.Now().Add(2 * time.Minute))

	testCases := []testCase{
		testCase{
			have: []metav1.ObjectMeta{
				metav1.ObjectMeta{
					Namespace:         "test",
					Name:              "o1",
					UID:               "1",
					CreationTimestamp: now,
				},
				metav1.ObjectMeta{
					Namespace:         "test",
					Name:              "o2",
					UID:               "2",
					CreationTimestamp: now_1m,
				},
				metav1.ObjectMeta{
					Namespace:         "test",
					Name:              "o3",
					UID:               "3",
					CreationTimestamp: now_2m,
				},
			},
			want: metav1.ObjectMeta{
				Namespace:         "test",
				Name:              "o1",
				UID:               "1",
				CreationTimestamp: now,
			},
		},
		testCase{
			have: []metav1.ObjectMeta{
				metav1.ObjectMeta{
					Namespace:         "test",
					Name:              "o1",
					UID:               "1",
					CreationTimestamp: now_2m,
				},
				metav1.ObjectMeta{
					Namespace:         "test",
					Name:              "o2",
					UID:               "2",
					CreationTimestamp: now_1m,
				},
				metav1.ObjectMeta{
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
