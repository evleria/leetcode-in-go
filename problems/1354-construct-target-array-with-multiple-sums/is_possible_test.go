package _354_construct_target_array_with_multiple_sums

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPossible(t *testing.T) {
	testCases := []struct {
		gotTarget []int
		want      bool
	}{
		{
			gotTarget: []int{9, 3, 5},
			want:      true,
		},
		{
			gotTarget: []int{1, 1, 1, 2},
			want:      false,
		},
		{
			gotTarget: []int{8, 5},
			want:      true,
		},
	}

	for _, testCase := range testCases {
		actual := isPossible(testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
