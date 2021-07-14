package _162_find_peak_element

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindPeakElement(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2},
			want:    1,
		},
		{
			gotNums: []int{1, 2, 3},
			want:    2,
		},
		{
			gotNums: []int{3, 2, 1},
			want:    0,
		},
		{
			gotNums: []int{1, 2, 3, 1},
			want:    2,
		},
		{
			gotNums: []int{1, 2, 1, 3, 5, 6, 4},
			want:    5,
		},
	}

	for _, testCase := range testCases {
		actual := findPeakElement(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
