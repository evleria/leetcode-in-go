package _554_brick_wall

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLeastBricks(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{1, 2, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 4},
				{3, 1, 2},
				{1, 3, 1, 1}},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := leastBricks(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
