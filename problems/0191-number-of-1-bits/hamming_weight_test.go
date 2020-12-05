package _191_number_of_1_bits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHammingWeight(t *testing.T) {
	testCases := []struct {
		got  uint32
		want int
	}{
		{
			got:  0b1011,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := hammingWeight(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
