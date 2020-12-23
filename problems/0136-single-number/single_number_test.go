package _136_single_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1},
			want: 1,
		},
		{
			got:  []int{5, 6, 3, 6, 5},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := singleNumber(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
