package _007_minimum_domino_rotations_for_equal_row

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinDominoRotations(t *testing.T) {
	testCases := []struct {
		gotTops    []int
		gotBottoms []int
		want       int
	}{
		{
			gotTops:    []int{2, 1, 2, 4, 2, 2},
			gotBottoms: []int{5, 2, 6, 2, 3, 2},
			want:       2,
		},
		{
			gotTops:    []int{3, 5, 1, 2, 3},
			gotBottoms: []int{3, 6, 3, 3, 4},
			want:       -1,
		},
		{
			gotTops:    []int{1, 2, 2},
			gotBottoms: []int{2, 1, 1},
			want:       1,
		},
	}

	for _, testCase := range testCases {
		actual := minDominoRotations(testCase.gotTops, testCase.gotBottoms)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
