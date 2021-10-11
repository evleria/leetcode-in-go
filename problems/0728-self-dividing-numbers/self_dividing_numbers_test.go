package _728_self_dividing_numbers

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSelfDividingNumbers(t *testing.T) {
	testCases := []struct {
		gotLeft  int
		gotRight int
		want     []int
	}{
		{
			gotLeft:  1,
			gotRight: 22,
			want:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			gotLeft:  47,
			gotRight: 85,
			want:     []int{48, 55, 66, 77},
		},
	}

	for _, testCase := range testCases {
		actual := selfDividingNumbers(testCase.gotLeft, testCase.gotRight)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
