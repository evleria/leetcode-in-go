package _764_largest_plus_sign

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOrderOfLargestPlusSign(t *testing.T) {
	testCases := []struct {
		gotNums   int
		gotTarget [][]int
		want      int
	}{
		{
			gotNums:   5,
			gotTarget: [][]int{{4, 2}},
			want:      2,
		},
		{
			gotNums:   1,
			gotTarget: [][]int{{0, 0}},
			want:      0,
		},
	}

	for _, testCase := range testCases {
		actual := orderOfLargestPlusSign(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
