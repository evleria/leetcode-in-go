package _350_intersection_of_two_arrays_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		got1 []int
		got2 []int
		want []int
	}{
		{
			got1: []int{1, 2, 2, 1},
			got2: []int{2, 2},
			want: []int{2, 2},
		},
		{
			got1: []int{4, 9, 5},
			got2: []int{9, 4, 9, 8, 4},
			want: []int{4, 9},
		},
	}

	for _, testCase := range testCases {
		actual := intersect(testCase.got1, testCase.got2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
