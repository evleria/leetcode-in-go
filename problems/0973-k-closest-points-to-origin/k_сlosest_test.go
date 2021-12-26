package _973_k_closest_points_to_origin

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKClosest(t *testing.T) {
	testCases := []struct {
		gotPoints [][]int
		gotK      int
		want      [][]int
	}{
		{
			gotPoints: [][]int{{1, 3}, {-2, 2}},
			gotK:      1,
			want:      [][]int{{-2, 2}},
		},
		{
			gotPoints: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			gotK:      2,
			want:      [][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, testCase := range testCases {
		actual := kClosest(testCase.gotPoints, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
