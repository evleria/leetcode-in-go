package _448_find_all_numbers_disappeared_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindDisappearedNumbers(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
	}

	for _, testCase := range testCases {
		actual := findDisappearedNumbers(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
