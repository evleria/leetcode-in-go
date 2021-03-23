package _923_3Sum_with_multiplicity

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestThreeSumMulti(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			gotTarget: 8,
			want:      20,
		},
		{
			gotNums:   []int{1, 1, 2, 2, 2, 2},
			gotTarget: 5,
			want:      12,
		},
	}

	for _, testCase := range testCases {
		actual := threeSumMulti(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
