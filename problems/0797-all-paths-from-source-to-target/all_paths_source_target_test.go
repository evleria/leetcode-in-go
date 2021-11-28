package _797_all_paths_from_source_to_target

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAllPathsSourceTarget(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got:  [][]int{{1, 2}, {3}, {3}, {}},
			want: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			got:  [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			want: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			got:  [][]int{{1, 2, 3}, {2}, {3}, {}},
			want: [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}},
		},
		{
			got:  [][]int{{1, 3}, {2}, {3}, {}},
			want: [][]int{{0, 1, 2, 3}, {0, 3}},
		},
	}

	for _, testCase := range testCases {
		actual := allPathsSourceTarget(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
