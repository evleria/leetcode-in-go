package _209_remove_all_adjacent_duplicates_in_string_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		gotS string
		gotK int
		want string
	}{
		{
			gotS: "abcd",
			gotK: 2,
			want: "abcd",
		},
		{
			gotS: "deeedbbcccbdaa",
			gotK: 3,
			want: "aa",
		},
		{
			gotS: "pbbcggttciiippooaais",
			gotK: 2,
			want: "ps",
		},
	}

	for _, testCase := range testCases {
		actual := removeDuplicates(testCase.gotS, testCase.gotK)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotS)
	}
}
