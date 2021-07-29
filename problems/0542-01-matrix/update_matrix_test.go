package _542_01_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUpdateMatrix(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			got:  [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}

	for _, testCase := range testCases {
		actual := updateMatrix(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
