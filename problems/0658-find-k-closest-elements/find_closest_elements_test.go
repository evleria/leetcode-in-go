package _658_find_k_closest_elements

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindClosestElements(t *testing.T) {
	testCases := []struct {
		gotArr []int
		gotK   int
		gotX   int
		want   []int
	}{
		{
			gotArr: []int{1, 5, 10},
			gotK:   1,
			gotX:   4,
			want:   []int{5},
		},
		{
			gotArr: []int{1, 3},
			gotK:   1,
			gotX:   2,
			want:   []int{1},
		},
		{
			gotArr: []int{1, 2, 3, 4, 5},
			gotK:   4,
			gotX:   3,
			want:   []int{1, 2, 3, 4},
		},
		{
			gotArr: []int{1, 2, 3, 4, 5},
			gotK:   4,
			gotX:   -1,
			want:   []int{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		actual := findClosestElements(testCase.gotArr, testCase.gotK, testCase.gotX)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
