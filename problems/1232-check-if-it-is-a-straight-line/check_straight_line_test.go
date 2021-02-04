package _232_check_if_it_is_a_straight_line

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestChekStraightLine(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want bool
	}{
		{
			got: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
				{6, 7},
			},
			want: true,
		},
		{
			got: [][]int{
				{1, 1},
				{2, 2},
				{3, 4},
				{4, 5},
				{5, 6},
				{7, 7},
			},
			want: false,
		},
		{
			got: [][]int{
				{0, 0},
				{0, 1},
				{0, -1},
			},
			want: true,
		},
		{
			got: [][]int{
				{2, 1},
				{4, 2},
				{6, 3},
			},
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := checkStraightLine(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
