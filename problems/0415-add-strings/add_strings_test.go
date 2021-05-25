package _415_add_strings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddStrings(t *testing.T) {
	testCases := []struct {
		gotNum1 string
		gotNum2 string
		want    string
	}{
		{
			gotNum1: "11",
			gotNum2: "123",
			want:    "134",
		},
		{
			gotNum1: "456",
			gotNum2: "77",
			want:    "533",
		},
		{
			gotNum1: "999",
			gotNum2: "10",
			want:    "1009",
		},
	}

	for _, testCase := range testCases {
		actual := addStrings(testCase.gotNum1, testCase.gotNum2)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotNum1, testCase.gotNum2)
	}
}
