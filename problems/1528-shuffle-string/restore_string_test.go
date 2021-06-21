package _528_shuffle_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRestoreString(t *testing.T) {
	testCases := []struct {
		gotS   string
		gotIdx []int
		want   string
	}{
		{
			gotS:   "codeleet",
			gotIdx: []int{4, 5, 6, 7, 0, 2, 1, 3},
			want:   "leetcode",
		},
		{
			gotS:   "abc",
			gotIdx: []int{0, 1, 2},
			want:   "abc",
		},
		{
			gotS:   "aiohn",
			gotIdx: []int{3, 1, 4, 2, 0},
			want:   "nihao",
		},
	}

	for _, testCase := range testCases {
		actual := restoreString(testCase.gotS, testCase.gotIdx)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotS)
	}
}
