package _215_find_the_difference_of_two_arrays

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindDifference(t *testing.T) {
	testCases := []struct {
		gotNums1 []int
		gotNums2 []int
		want     [][]int
	}{
		{
			gotNums1: []int{1, 2, 3, 3},
			gotNums2: []int{1, 1, 2, 2},
			want:     [][]int{{3}, {}},
		},
		{
			gotNums1: []int{1, 2, 3},
			gotNums2: []int{2, 4, 6},
			want:     [][]int{{1, 3}, {4, 6}},
		},
	}

	for _, testCase := range testCases {
		actual := findDifference(testCase.gotNums1, testCase.gotNums2)

		sort.Ints(actual[0])
		sort.Ints(actual[1])

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
