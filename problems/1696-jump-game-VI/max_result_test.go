package _696_jump_game_VI

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxResult(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    int
	}{
		{
			gotNums: []int{1, -1, -2, 4, -7, 3},
			gotK:    2,
			want:    7,
		},
		{
			gotNums: []int{10, -5, -2, 4, 0, 3},
			gotK:    3,
			want:    17,
		},
		{
			gotNums: []int{1, -5, -20, 4, -1, 3, -6, -3},
			gotK:    2,
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := maxResult(testCase.gotNums, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
