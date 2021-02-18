package _287_find_the_duplicate_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			got:  []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			want: 9,
		},
	}

	for _, testCase := range testCases {
		actual := findDuplicate(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
