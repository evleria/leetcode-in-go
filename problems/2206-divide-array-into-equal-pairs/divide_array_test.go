package _206_divide_array_into_equal_pairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDivideArray(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    bool
	}{
		{
			gotNums: []int{3, 2, 3, 2, 2, 2},
			want:    true,
		},
		{
			gotNums: []int{1, 2, 3, 4},
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := divideArray(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
