package _309_decrypt_string_from_alphabet_to_integer_mapping

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFreqAlphabets(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "10#11#12",
			want: "jkab",
		},
		{
			got:  "1326#1",
			want: "acza",
		},
	}

	for _, testCase := range testCases {
		actual := freqAlphabets(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
