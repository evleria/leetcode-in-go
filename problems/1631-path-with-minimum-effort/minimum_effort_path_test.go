package _631_path_with_minimum_effort

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumEffortPath(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			want:    2,
		},
		{
			gotNums: [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			want:    1,
		},
		{
			gotNums: [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := minimumEffortPath(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
