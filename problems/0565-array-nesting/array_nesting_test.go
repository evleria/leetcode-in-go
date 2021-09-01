package _565_array_nesting

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestArrayNesting(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{5, 4, 0, 3, 1, 6, 2},
			want:    4,
		},
		{
			gotNums: []int{0, 1, 2},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := arrayNesting(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
