package _283_move_zeroes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
	}

	for _, testCase := range testCases {
		moveZeroes(testCase.got)

		assert.Check(t, is.DeepEqual(testCase.got, testCase.want), testCase.got)
	}
}
