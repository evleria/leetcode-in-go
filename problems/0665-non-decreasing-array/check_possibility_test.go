package _665_non_decreasing_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCheckPossibility(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    bool
	}{
		{
			gotNums: []int{3, 4, 2, 3},
			want:    false,
		},
		{
			gotNums: []int{4, 2, 3},
			want:    true,
		},
		{
			gotNums: []int{3, 2, 1},
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := checkPossibility(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotNums)
	}
}
