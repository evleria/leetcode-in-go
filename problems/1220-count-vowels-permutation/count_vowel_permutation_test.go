package _220_count_vowels_permutation

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountVowelsPermutation(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  1,
			want: 5,
		},
		{
			got:  2,
			want: 10,
		},
		{
			got:  5,
			want: 68,
		},
	}

	for _, testCase := range testCases {
		actual := countVowelPermutation(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
