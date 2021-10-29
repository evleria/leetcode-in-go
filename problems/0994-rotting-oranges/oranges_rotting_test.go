package _994_rotting_oranges

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOrangesRotting(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			want: 4,
		},
		{
			got:  [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			want: -1,
		},
		{
			got:  [][]int{{0, 2}},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := orangesRotting(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
