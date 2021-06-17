package _795_number_of_subarrays_with_bounded_maximum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumSubarrayBoundedMax(t *testing.T) {
	testCases := []struct {
		gotNums  []int
		gotLeft  int
		gotRight int
		want     int
	}{
		{
			gotNums:  []int{2, 1, 4, 3},
			gotLeft:  2,
			gotRight: 3,
			want:     3,
		},
	}

	for _, testCase := range testCases {
		actual := numSubarrayBoundedMax(testCase.gotNums, testCase.gotLeft, testCase.gotRight)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
