package _074_search_a_2d_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		gotMatrix [][]int
		gotTarget int
		want      bool
	}{
		{
			gotMatrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			gotTarget: 3,
			want:      true,
		},
		{
			gotMatrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			gotTarget: 13,
			want:      false,
		},
	}

	for _, testCase := range testCases {
		actual := searchMatrix(testCase.gotMatrix, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
