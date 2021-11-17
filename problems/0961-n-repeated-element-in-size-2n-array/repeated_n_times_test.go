package _961_n_repeated_element_in_size_2n_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRepeatedNTimes(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3, 3},
			want:    3,
		},
		{
			gotNums: []int{2, 1, 2, 5, 3, 2},
			want:    2,
		},
		{
			gotNums: []int{5, 1, 5, 2, 5, 3, 5, 4},
			want:    5,
		},
	}

	for _, testCase := range testCases {
		actual := repeatedNTimes(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
