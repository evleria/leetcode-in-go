package _847_shortest_path_visiting_all_nodes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestShortestPathLength(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{0}},
			want: 0,
		},
		{
			got:  [][]int{{1, 2, 3}, {0}, {0}, {0}},
			want: 4,
		},
		{
			got:  [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}},
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := shortestPathLength(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
