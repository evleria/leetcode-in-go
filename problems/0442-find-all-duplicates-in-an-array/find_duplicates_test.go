package _442_find_all_duplicates_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{2, 3},
		},
		{
			got:  []int{1, 1, 2},
			want: []int{1},
		},
		{
			got:  []int{1},
			want: []int{},
		},
	}

	for _, testCase := range testCases {
		actual := findDuplicates(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
