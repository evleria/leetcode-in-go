package _684_redundant_connection

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindRedundantConnection(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want []int
	}{
		/*{
			got:  [][]int{{1, 2}, {1, 3}, {2, 3}},
			want: []int{2, 3},
		},*/
		{
			got:  [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			want: []int{1, 4},
		},
	}

	for _, testCase := range testCases {
		actual := findRedundantConnection(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
