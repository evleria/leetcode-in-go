package _905_sort_array_by_parity

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortArrayByParity(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{3, 1, 2, 4},
			want:    []int{4, 2, 1, 3},
		},
	}

	for _, testCase := range testCases {
		actual := sortArrayByParity(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
