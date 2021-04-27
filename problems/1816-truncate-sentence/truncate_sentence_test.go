package _816_truncate_sentence

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestTruncateSentence(t *testing.T) {
	testCases := []struct {
		got  string
		gotK int
		want string
	}{
		{
			got:  "Hello how are you Contestant",
			gotK: 4,
			want: "Hello how are you",
		},
		{
			got:  "What is the solution to this problem",
			gotK: 7,
			want: "What is the solution to this problem",
		},
	}

	for _, testCase := range testCases {
		actual := truncateSentence(testCase.got, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
