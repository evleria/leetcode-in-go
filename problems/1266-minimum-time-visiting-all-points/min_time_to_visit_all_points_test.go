package _266_minimum_time_visiting_all_points

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinTimeVisitAllPoints(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{1, 1}, {3, 4}, {-1, 0}},
			want: 7,
		},
		{
			got:  [][]int{{3, 2}, {-2, 2}},
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := minTimeToVisitAllPoints(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
