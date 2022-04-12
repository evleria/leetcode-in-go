package _289_game_of_life

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWordPattern(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
	}

	for _, testCase := range testCases {
		gameOfLife(testCase.got)

		assert.Check(t, is.DeepEqual(testCase.got, testCase.want))
	}
}
