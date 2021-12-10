package _790_domino_and_tromino_tiling

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumTilings(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    int
	}{
		{
			gotNums: 3,
			want:    5,
		},
		{
			gotNums: 1,
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := numTilings(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
