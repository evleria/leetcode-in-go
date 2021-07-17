package _927_three_equal_parts

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestThreeEqualParts(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0},
			want: []int{15, 32},
		},
		{
			got:  []int{0, 0, 0, 0, 0},
			want: []int{0, 4},
		},
		{
			got:  []int{1, 1, 1},
			want: []int{0, 2},
		},
		{
			got:  []int{0, 1, 1, 0, 0, 1},
			want: []int{1, 3},
		},
		{
			got:  []int{1, 0, 1, 0, 1},
			want: []int{0, 3},
		},
		{
			got:  []int{1, 1, 0, 1, 1},
			want: []int{-1, -1},
		},
	}

	for _, testCase := range testCases {
		actual := threeEqualParts(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintln(testCase.want))
	}
}
