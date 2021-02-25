package _567_permutation_in_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCheckInclusion(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want bool
	}{
		{
			got1: "ab",
			got2: "eidbaooo",
			want: true,
		},
		{
			got1: "ab",
			got2: "eidboaoo",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := checkInclusion(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
