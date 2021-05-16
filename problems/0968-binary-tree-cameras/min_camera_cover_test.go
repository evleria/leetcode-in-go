package _968_binary_tree_cameras

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinCameraCover(t *testing.T) {
	testCases := []struct {
		gotTree []int
		want    int
	}{
		{
			gotTree: []int{0, 0, 0, NULL, NULL, NULL, 0},
			want:    2,
		},
		{
			gotTree: []int{0, 0, NULL, 0, 0},
			want:    1,
		},
		{
			gotTree: []int{0, 0, NULL, 0, NULL, 0, NULL, NULL, 0},
			want:    2,
		},
	}

	for _, testCase := range testCases {
		actual := minCameraCover(BinaryTreeFromSlice(testCase.gotTree))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
