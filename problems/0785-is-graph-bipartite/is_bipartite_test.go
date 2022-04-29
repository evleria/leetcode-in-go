package _785_is_graph_bipartite

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsBipartite(t *testing.T) {
	testCases := []struct {
		gotGraph [][]int
		want     bool
	}{
		{
			gotGraph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			want:     false,
		},
		{
			gotGraph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			want:     true,
		},
	}

	for _, testCase := range testCases {
		actual := isBipartite(testCase.gotGraph)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
