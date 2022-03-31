package _410_split_array_largest_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSplitArray(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotM    int
		want    int
	}{
		{
			gotNums: []int{7, 2, 5, 10, 8},
			gotM:    2,
			want:    18,
		},
		{
			gotNums: []int{1, 2, 3, 4, 5},
			gotM:    2,
			want:    9,
		},
		{
			gotNums: []int{1, 4, 4},
			gotM:    3,
			want:    4,
		},
	}

	for _, testCase := range testCases {
		actual := splitArray(testCase.gotNums, testCase.gotM)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
