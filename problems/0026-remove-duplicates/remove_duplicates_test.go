package _026_remove_duplicates

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
			got:       []int{1, 1, 2},
			wantLen:   2,
			wantSlice: []int{1, 2},
		},
		{
			got:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantLen:   5,
			wantSlice: []int{0, 1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		l := removeDuplicates(testCase.got)

		assert.Check(t, is.Equal(l, testCase.wantLen), testCase.got)
		assert.Check(t, is.DeepEqual(testCase.got[:l], testCase.wantSlice))
	}
}
