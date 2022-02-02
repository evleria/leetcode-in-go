package _438_find_all_anagrams_in_a_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		gotS string
		gotP string
		want []int
	}{
		{
			gotS: "aaaaaaaaaa",
			gotP: "aaaaaaaaaaaaa",
			want: []int{},
		},
		{
			gotS: "cbaebabacd",
			gotP: "abc",
			want: []int{0, 6},
		},
		{
			gotS: "abab",
			gotP: "ab",
			want: []int{0, 1, 2},
		},
	}

	for _, testCase := range testCases {
		actual := findAnagrams(testCase.gotS, testCase.gotP)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
