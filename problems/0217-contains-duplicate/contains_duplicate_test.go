package _217_contains_duplicate

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1, 2, 3, 1},
			want: true,
		},
		{
			got:  []int{1, 2, 3},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := containsDuplicate(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
