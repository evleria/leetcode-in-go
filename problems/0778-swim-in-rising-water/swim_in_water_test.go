package _778_swim_in_rising_water

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSwimInWater(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{0, 2}, {1, 3}},
			want: 3,
		},
		{
			got: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			want: 16,
		},
		{
			got:  [][]int{{3, 2}, {0, 1}},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := swimInWater(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
