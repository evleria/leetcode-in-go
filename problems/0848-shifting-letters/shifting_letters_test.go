package _848_shifting_letters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestShiftingLetters(t *testing.T) {
	testCases := []struct {
		gotS      string
		gotShifts []int
		want      string
	}{
		{
			gotS:      "aaa",
			gotShifts: []int{1, 2, 3},
			want:      "gfd",
		},
		{
			gotS:      "abc",
			gotShifts: []int{3, 5, 9},
			want:      "rpl",
		},
	}

	for _, testCase := range testCases {
		actual := shiftingLetters(testCase.gotS, testCase.gotShifts)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
