package _045_jump_game_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestJump(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{2, 3, 1, 1, 4},
			want:    2,
		},
		{
			gotNums: []int{2, 3, 0, 1, 4},
			want:    2,
		},
	}

	for _, testCase := range testCases {
		actual := jump(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
