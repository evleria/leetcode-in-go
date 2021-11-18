package _791_find_center_of_star_graph

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindCenter(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{{1, 2}, {2, 3}, {4, 2}},
			want:    2,
		},
		{
			gotNums: [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := findCenter(testCase.gotNums)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
