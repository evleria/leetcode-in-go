package _122_relative_sort_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRelativeSortArray(t *testing.T) {
	testCases := []struct {
		got1 []int
		got2 []int
		want []int
	}{
		{
			got1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			got2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			got1: []int{28, 6, 22, 8, 44, 17},
			got2: []int{22, 28, 8, 6},
			want: []int{22, 28, 8, 6, 17, 44},
		},
	}

	for _, testCase := range testCases {
		actual := relativeSortArray(testCase.got1, testCase.got2)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got1)
	}
}
