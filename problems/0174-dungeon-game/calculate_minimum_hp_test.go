package _174_dungeon_game

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}},
			want: 7,
		},
		{
			got:  [][]int{{0}},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := calculateMinimumHP(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
