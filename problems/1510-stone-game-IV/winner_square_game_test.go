package _510_stone_game_IV

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWinnerSquareGame(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  1,
			want: true,
		},
		{
			got:  2,
			want: false,
		},
		{
			got:  4,
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := winnerSquareGame(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
