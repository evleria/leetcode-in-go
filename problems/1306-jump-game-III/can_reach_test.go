package _306_jump_game_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanReach(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      bool
	}{
		{
			gotNums:   []int{4, 2, 3, 0, 3, 1, 2},
			gotTarget: 5,
			want:      true,
		},
		{
			gotNums:   []int{4, 2, 3, 0, 3, 1, 2},
			gotTarget: 0,
			want:      true,
		},
		{
			gotNums:   []int{3, 0, 2, 1, 2},
			gotTarget: 2,
			want:      false,
		},
	}

	for _, testCase := range testCases {
		actual := canReach(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
