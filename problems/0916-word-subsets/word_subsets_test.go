package _916_word_subsets

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWordSubsets(t *testing.T) {
	testCases := []struct {
		gotA []string
		gotB []string
		want []string
	}{
		{
			gotA: []string{"amazon", "apple", "facebook", "google", "leetcode"},
			gotB: []string{"ec", "oc", "ceo"},
			want: []string{"facebook", "leetcode"},
		},
	}

	for _, testCase := range testCases {
		actual := wordSubsets(testCase.gotA, testCase.gotB)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
