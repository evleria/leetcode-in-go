package _201_bitwise_and_of_numbers_range

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRangeBitwiseAnd(t *testing.T) {
	testCases := []struct {
		gotLeft  int
		gotRight int
		want     int
	}{
		{
			gotLeft:  5,
			gotRight: 7,
			want:     4,
		},
		{
			gotLeft:  0,
			gotRight: 0,
			want:     0,
		},
		{
			gotLeft:  1,
			gotRight: 2147483647,
			want:     0,
		},
	}

	for _, testCase := range testCases {
		actual := rangeBitwiseAnd(testCase.gotLeft, testCase.gotRight)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
