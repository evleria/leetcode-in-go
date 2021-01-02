package _027_remove_element

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		gotNums    []int
		gotVal     int
		wantNums   []int
		wantLength int
	}{
		{
			gotNums:    []int{3, 2, 2, 3},
			gotVal:     3,
			wantNums:   []int{2, 2},
			wantLength: 2,
		},
		{
			gotNums:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			gotVal:     2,
			wantNums:   []int{0, 1, 3, 0, 4},
			wantLength: 5,
		},
	}

	for _, testCase := range testCases {
		actual := removeElement(testCase.gotNums, testCase.gotVal)

		assert.Check(t, is.DeepEqual(testCase.gotNums[:testCase.wantLength], testCase.wantNums))
		assert.Check(t, is.Equal(actual, testCase.wantLength))
	}
}
