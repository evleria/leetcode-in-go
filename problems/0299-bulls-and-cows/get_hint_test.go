package _299_bulls_and_cows

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetHint(t *testing.T) {
	testCases := []struct {
		gotSecret string
		gotGuess  string
		want      string
	}{
		{
			gotSecret: "1807",
			gotGuess:  "7810",
			want:      "1A3B",
		},
		{
			gotSecret: "1123",
			gotGuess:  "0111",
			want:      "1A1B",
		},
		{
			gotSecret: "0122",
			gotGuess:  "1133",
			want:      "1A0B",
		},
		{
			gotSecret: "1",
			gotGuess:  "0",
			want:      "0A0B",
		},
	}

	for _, testCase := range testCases {
		actual := getHint(testCase.gotSecret, testCase.gotGuess)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
