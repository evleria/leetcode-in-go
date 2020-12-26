package _349_intersection_of_two_arrays

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		gotNums1 []int
		gotNums2 []int
		want     []int
	}{
		{
			gotNums1: []int{1, 2, 2, 1},
			gotNums2: []int{2, 2},
			want:     []int{2},
		},
		{
			gotNums1: []int{4, 9, 5},
			gotNums2: []int{9, 4, 9, 8, 4},
			want:     []int{9, 4},
		},
	}

	for _, testCase := range testCases {
		actual := intersection(testCase.gotNums1, testCase.gotNums2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
