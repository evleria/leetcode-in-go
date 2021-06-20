package _771_jewels_and_stones

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		gotJ string
		gotS string
		want int
	}{
		{
			gotJ: "aA",
			gotS: "aAAbbbb",
			want: 3,
		},
		{
			gotJ: "z",
			gotS: "ZZ",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := numJewelsInStones(testCase.gotJ, testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
