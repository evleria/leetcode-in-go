package _454_4Sum_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFourSumCount(t *testing.T) {
	testCases := []struct {
		gotA []int
		gotB []int
		gotC []int
		gotD []int
		want int
	}{
		{
			gotA: []int{1, 2},
			gotB: []int{-2, -1},
			gotC: []int{-1, 2},
			gotD: []int{0, 2},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := fourSumCount(testCase.gotA, testCase.gotB, testCase.gotC, testCase.gotD)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
