package _986_interval_list_intersections

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIntervalIntersection(t *testing.T) {
	testCases := []struct {
		gotF [][]int
		gotS [][]int
		want [][]int
	}{
		{
			gotF: [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			gotS: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			gotF: [][]int{{1, 3}, {5, 9}},
			gotS: [][]int{},
			want: [][]int{},
		},
		{
			gotF: [][]int{},
			gotS: [][]int{{4, 8}, {10, 12}},
			want: [][]int{},
		},
		{
			gotF: [][]int{{1, 7}},
			gotS: [][]int{{3, 10}},
			want: [][]int{{3, 7}},
		},
	}

	for _, testCase := range testCases {
		actual := intervalIntersection(testCase.gotF, testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
