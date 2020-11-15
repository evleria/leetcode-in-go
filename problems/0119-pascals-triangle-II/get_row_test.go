package _119_pascals_triangle_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetRow(t *testing.T) {
	testCases := []struct {
		got  int
		want []int
	}{
		{
			got:  3,
			want: []int{1, 3, 3, 1},
		},
	}
	for _, testCase := range testCases {
		actual := getRow(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
