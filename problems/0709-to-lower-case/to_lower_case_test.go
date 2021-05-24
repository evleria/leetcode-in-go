package _709_to_lower_case

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestToLowerCase(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "HELLO",
			want: "hello",
		},
		{
			got:  "Abc",
			want: "abc",
		},
	}

	for _, testCase := range testCases {
		actual := toLowerCase(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
