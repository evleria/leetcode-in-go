package _046_last_stone_weight

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLastStoneWeight(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 7, 2},
			want: 2,
		},
		{
			got:  []int{1, 3},
			want: 2,
		},
		{
			got:  []int{2, 7, 4, 1, 8, 1},
			want: 1,
		},
		{
			got:  []int{1},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := lastStoneWeight(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
