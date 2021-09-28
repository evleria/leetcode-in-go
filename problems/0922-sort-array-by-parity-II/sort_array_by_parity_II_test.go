package _922_sort_array_by_parity_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortArrayByParityII(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{4, 2, 5, 7},
			want: []int{4, 5, 2, 7},
		},
		{
			got:  []int{2, 3},
			want: []int{2, 3},
		},
	}

	for _, testCase := range testCases {
		actual := sortArrayByParityII(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
