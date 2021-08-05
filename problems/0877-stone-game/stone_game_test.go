package _877_stone_game

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestStoneGame(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    bool
	}{
		{
			gotNums: []int{5, 3, 4, 5},
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := stoneGame(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
