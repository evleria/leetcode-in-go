package _485_max_consecutive_ones

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			got:  []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := findMaxConsecutiveOnes(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
