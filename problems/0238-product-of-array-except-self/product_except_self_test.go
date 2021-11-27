package _238_product_of_array_except_self

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			got:  []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			got:  []int{0, 1, 0, -3, 3},
			want: []int{0, 0, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := productExceptSelf(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
