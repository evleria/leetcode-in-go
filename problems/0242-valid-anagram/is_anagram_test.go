package _242_valid_anagram

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want bool
	}{
		{
			got1: "anagram",
			got2: "nagaram",
			want: true,
		},
		{
			got1: "rat",
			got2: "car",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isAnagram(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
