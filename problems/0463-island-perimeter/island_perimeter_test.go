package _463_island_perimeter

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIslandPerimeter(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			want: 16,
		},
		{
			got:  [][]int{{1}},
			want: 4,
		},
		{
			got:  [][]int{{1, 0}},
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := islandPerimeter(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
