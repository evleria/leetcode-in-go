package _383_maximum_performance_of_a_team

import (
	"fmt"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestMaxPerformance(t *testing.T) {
	testCases := []struct {
		gotN          int
		gotSpeed      []int
		gotEfficienty []int
		gotK          int
		want          int
	}{
		{
			gotN:          6,
			gotSpeed:      []int{2, 10, 3, 1, 5, 8},
			gotEfficienty: []int{5, 4, 3, 9, 7, 2},
			gotK:          2,
			want:          60,
		},
		{
			gotN:          6,
			gotSpeed:      []int{2, 10, 3, 1, 5, 8},
			gotEfficienty: []int{5, 4, 3, 9, 7, 2},
			gotK:          3,
			want:          68,
		},
		{
			gotN:          6,
			gotSpeed:      []int{2, 10, 3, 1, 5, 8},
			gotEfficienty: []int{5, 4, 3, 9, 7, 2},
			gotK:          4,
			want:          72,
		},
	}

	for _, testCase := range testCases {
		actual := maxPerformance(testCase.gotN, testCase.gotSpeed, testCase.gotEfficienty, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
