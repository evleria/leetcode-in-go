package _089_find_target_indices_after_sorting_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTargetIndices(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      []int
	}{
		{
			gotNums:   []int{1, 2, 5, 2, 3},
			gotTarget: 2,
			want:      []int{1, 2},
		},
		{
			gotNums:   []int{1, 2, 5, 2, 3},
			gotTarget: 3,
			want:      []int{3},
		},
		{
			gotNums:   []int{1, 2, 5, 2, 3},
			gotTarget: 5,
			want:      []int{4},
		},
		{
			gotNums:   []int{1, 2, 5, 2, 3},
			gotTarget: 4,
			want:      []int{},
		},
	}

	for _, testCase := range testCases {
		actual := targetIndices(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
