package _474_ones_and_zeroes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMaxForm(t *testing.T) {
	testCases := []struct {
		gotStrings []string
		gotZeroes  int
		gotOnes    int
		want       int
	}{
		{
			gotStrings: []string{"10", "0001", "111001", "1", "0"},
			gotZeroes:  5,
			gotOnes:    3,
			want:       4,
		},
		{
			gotStrings: []string{"10", "0", "1"},
			gotZeroes:  1,
			gotOnes:    1,
			want:       2,
		},
	}

	for _, testCase := range testCases {
		actual := findMaxForm(testCase.gotStrings, testCase.gotZeroes, testCase.gotOnes)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
