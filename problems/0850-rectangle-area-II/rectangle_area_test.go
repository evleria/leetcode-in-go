package _850_rectangle_area_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRectangleArea(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}},
			want:    6,
		},
		{
			gotNums: [][]int{{0, 0, 1000000000, 1000000000}},
			want:    49,
		},
	}

	for _, testCase := range testCases {
		actual := rectangleArea(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
