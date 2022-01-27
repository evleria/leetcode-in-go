package _421_maximum_xor_of_two_numbers_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMaximumXOR(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70},
			want: 127,
		},
		{
			got:  []int{3, 10, 5, 25, 2, 8},
			want: 28,
		},
	}

	for _, testCase := range testCases {
		actual := findMaximumXOR(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
