package _566_reshape_the_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMatrixReshape(t *testing.T) {
	testCases := []struct {
		got  [][]int
		gotR int
		gotC int
		want [][]int
	}{
		{
			got: [][]int{
				{1, 2},
				{3, 4},
			},
			gotR: 1,
			gotC: 4,
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			got: [][]int{
				{1, 2},
				{3, 4},
			},
			gotR: 2,
			gotC: 4,
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			got: [][]int{
				{1, 2, 3, 4},
			},
			gotR: 2,
			gotC: 2,
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}

	for _, testCase := range testCases {
		actual := matrixReshape(testCase.got, testCase.gotR, testCase.gotC)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
