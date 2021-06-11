package _690_stone_game_VII

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestStoneGameVII(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{5, 3, 1, 4, 2},
			want:    6,
		},
		{
			gotNums: []int{7, 90, 5, 1, 100, 10, 10, 2},
			want:    122,
		},
	}

	for _, testCase := range testCases {
		actual := stoneGameVII(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
