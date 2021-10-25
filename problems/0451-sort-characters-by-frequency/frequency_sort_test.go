package _451_sort_characters_by_frequency

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFrequencySort(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "tree",
			want: "eetr",
		},
		{
			got:  "cccaaa",
			want: "aaaccc",
		},
		{
			got:  "Aabb",
			want: "bbAa",
		},
	}

	for _, testCase := range testCases {
		actual := frequencySort(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
