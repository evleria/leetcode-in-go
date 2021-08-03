package _090_subsets_II

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSubsetsWithDup(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    [][]int
	}{
		{
			gotNums: []int{1, 2, 2},
			want:    [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			gotNums: []int{0},
			want:    [][]int{{}, {0}},
		},
	}

	for _, testCase := range testCases {
		actual := subsetsWithDup(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", actual))
	}
}
