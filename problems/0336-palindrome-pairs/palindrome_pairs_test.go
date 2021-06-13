package _336_palindrome_pairs

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPalindromePairs(t *testing.T) {
	testCases := []struct {
		gotwords []string
		want     [][]int
	}{
		{
			gotwords: []string{"a", "ba", "b", "abb"},
			want:     [][]int{{0, 1}, {1, 2}, {3, 0}, {3, 1}},
		},
		{
			gotwords: []string{"abcd", "dcba", "lls", "s", "sssll"},
			want:     [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}},
		},
		{
			gotwords: []string{"bat", "tab", "cat"},
			want:     [][]int{{0, 1}, {1, 0}},
		},
		{
			gotwords: []string{"a", ""},
			want:     [][]int{{0, 1}, {1, 0}},
		},
	}

	for _, testCase := range testCases {
		actual := palindromePairs(testCase.gotwords)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
