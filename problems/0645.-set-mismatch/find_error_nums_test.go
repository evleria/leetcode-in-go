package _645__set_mismatch

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindErrorNums(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 2, 2, 4},
			want:    []int{2, 3},
		},
		{
			gotNums: []int{1, 1},
			want:    []int{1, 2},
		},
		{
			gotNums: []int{2, 2},
			want:    []int{2, 1},
		},
		{
			gotNums: []int{3, 2, 2},
			want:    []int{2, 1},
		},
		{
			gotNums: []int{3, 2, 3, 4, 6, 5},
			want:    []int{3, 1},
		},
	}

	for _, testCase := range testCases {
		actual := findErrorNums(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
