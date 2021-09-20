package _275_find_winner_on_a_tic_tac_toe_game

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTictactoe(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want string
	}{
		{
			got:  [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}},
			want: "A",
		},
		{
			got:  [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
			want: "B",
		},
		{
			got:  [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}},
			want: "Draw",
		},
		{
			got:  [][]int{{0, 0}, {1, 1}},
			want: "Pending",
		},
	}

	for _, testCase := range testCases {
		actual := tictactoe(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
