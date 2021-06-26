package _688_count_of_matches_in_tournament

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumberOfMatches(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    int
	}{
		{
			gotNums: 7,
			want:    6,
		},
		{
			gotNums: 14,
			want:    13,
		},
	}

	for _, testCase := range testCases {
		actual := numberOfMatches(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
