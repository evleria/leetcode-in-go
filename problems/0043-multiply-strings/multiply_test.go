package _043_multiply_strings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want string
	}{
		{
			got1: "2",
			got2: "3",
			want: "6",
		},
		{
			got1: "19",
			got2: "0",
			want: "0",
		},
		{
			got1: "6",
			got2: "8",
			want: "48",
		},
		{
			got1: "123",
			got2: "456",
			want: "56088",
		},
	}

	for _, testCase := range testCases {
		actual := multiply(testCase.got1, testCase.got2)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got1, testCase.got2)
	}
}
