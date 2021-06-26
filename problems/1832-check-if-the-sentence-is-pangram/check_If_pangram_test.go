package _832_check_if_the_sentence_is_pangram

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCheckIfPangram(t *testing.T) {
	testCases := []struct {
		gotS string
		want bool
	}{
		{
			gotS: "thequickbrownfoxjumpsoverthelazydog",
			want: true,
		},
		{
			gotS: "leetcode",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := checkIfPangram(testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
