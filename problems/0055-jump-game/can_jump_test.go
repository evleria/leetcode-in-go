package _055_jump_game

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			got:  []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			got:  []int{0, 4},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := canJump(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
