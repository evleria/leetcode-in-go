package _207_unique_number_of_occurrences

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUniqueOccurrence(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1, 2, 2, 1, 1, 3},
			want: true,
		},
		{
			got:  []int{1, 2},
			want: false,
		},
		{
			got:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := uniqueOccurrences(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
