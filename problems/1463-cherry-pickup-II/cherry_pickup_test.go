package _463_cherry_pickup_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCherryPickup(t *testing.T) {
	testCases := []struct {
		gotGrid [][]int
		want    int
	}{
		{
			gotGrid: [][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}},
			want:    24,
		},
		{
			gotGrid: [][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}},
			want:    28,
		},
	}

	for _, testCase := range testCases {
		actual := cherryPickup(testCase.gotGrid)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
