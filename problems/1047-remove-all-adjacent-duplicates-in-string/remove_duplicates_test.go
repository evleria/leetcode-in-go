package _047_remove_all_adjacent_duplicates_in_string

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"

	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "abbaca",
			want: "ca",
		},
		{
			got:  "azxxz",
			want: "a",
		},
	}

	for _, testCase := range testCases {
		actual := removeDuplicates(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
