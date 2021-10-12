package _374_guess_number_higher_or_lower

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGuessNumber(t *testing.T) {
	testCases := []struct {
		gotPick int
		gotN    int
	}{
		{
			gotPick: 6,
			gotN:    10,
		},
		{
			gotPick: 1,
			gotN:    1,
		},
		{
			gotPick: 1,
			gotN:    2,
		},
		{
			gotPick: 2,
			gotN:    2,
		},
	}

	for _, testCase := range testCases {
		Pick = testCase.gotPick
		actual := guessNumber(testCase.gotN)

		assert.Check(t, is.Equal(actual, Pick))
	}
}
