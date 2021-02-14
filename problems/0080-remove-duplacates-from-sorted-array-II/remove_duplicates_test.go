package _080_remove_duplacates_from_sorted_array_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		got       []int
		wantLen   int
		wantSlice []int
	}{
		{
			got:       []int{3},
			wantLen:   1,
			wantSlice: []int{3},
		},
		{
			got:       []int{4, 5},
			wantLen:   2,
			wantSlice: []int{4, 5},
		},
		{
			got:       []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			wantLen:   7,
			wantSlice: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}

	for _, testCase := range testCases {
		l := removeDuplicates(testCase.got)

		assert.Check(t, is.Equal(l, testCase.wantLen), testCase.got)
		assert.Check(t, is.DeepEqual(testCase.got[:l], testCase.wantSlice))
	}
}
