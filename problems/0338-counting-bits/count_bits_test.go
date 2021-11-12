package _338_counting_bits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountBits(t *testing.T) {
	testCases := []struct {
		got  int
		want []int
	}{
		{
			got:  2,
			want: []int{0, 1, 1},
		},
		{
			got:  5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, testCase := range testCases {
		actual := countBits(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
