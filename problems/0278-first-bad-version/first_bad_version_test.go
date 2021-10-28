package _278_first_bad_version

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGuessNumber(t *testing.T) {
	testCases := []struct {
		gotBad int
		gotN   int
	}{
		{
			gotBad: 4,
			gotN:   5,
		},
		{
			gotBad: 1,
			gotN:   1,
		},
	}

	for _, testCase := range testCases {
		Bad = testCase.gotBad
		actual := firstBadVersion(testCase.gotN)

		assert.Check(t, is.Equal(actual, Bad))
	}
}
