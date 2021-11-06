package _260_single_number_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 2, 1, 3, 2, 5},
			want:    []int{3, 5},
		},
		{
			gotNums: []int{-1, 0},
			want:    []int{-1, 0},
		},
		{
			gotNums: []int{0, 1},
			want:    []int{1, 0},
		},
	}

	for _, testCase := range testCases {
		actual := singleNumber(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
