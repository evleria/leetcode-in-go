package _065_valid_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsNumber(t *testing.T) {
	testCases := []struct {
		gotNums string
		want    bool
	}{
		{
			gotNums: "0",
			want:    true,
		},
		{
			gotNums: "e",
			want:    false,
		},
		{
			gotNums: ".",
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := isNumber(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
