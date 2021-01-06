package _016_3Sum_closest

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestThreeSumClosest(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{-1, 2, 1, -4},
			gotTarget: 1,
			want:      2,
		},
	}

	for _, testCase := range testCases {
		actual := threeSumClosest(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
