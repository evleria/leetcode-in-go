package _882_reachable_nodes_in_subdivided_graph

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReachableNodes(t *testing.T) {
	testCases := []struct {
		gotEdges    [][]int
		gotMaxMoves int
		gotN        int
		want        int
	}{
		{
			gotEdges:    [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}},
			gotMaxMoves: 6,
			gotN:        3,
			want:        13,
		},
		{
			gotEdges:    [][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}},
			gotMaxMoves: 10,
			gotN:        4,
			want:        23,
		},
		{
			gotEdges:    [][]int{{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}},
			gotMaxMoves: 17,
			gotN:        5,
			want:        1,
		},
	}

	for _, testCase := range testCases {
		actual := reachableNodes(testCase.gotEdges, testCase.gotMaxMoves, testCase.gotN)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
