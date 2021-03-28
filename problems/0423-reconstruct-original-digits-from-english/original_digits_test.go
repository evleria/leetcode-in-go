package _423_reconstruct_original_digits_from_english

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOriginalDigits(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "owoztneoer",
			want: "012",
		},
		{
			got:  "fviefuro",
			want: "45",
		},
	}

	for _, testCase := range testCases {
		actual := originalDigits(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
