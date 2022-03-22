package _663_smallest_string_with_a_given_numeric_value

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetSmallestString(t *testing.T) {
	testCases := []struct {
		gotN int
		gotK int
		want string
	}{
		{
			gotN: 5,
			gotK: 130,
			want: "zzzzz",
		},
		{
			gotN: 3,
			gotK: 27,
			want: "aay",
		},
		{
			gotN: 5,
			gotK: 73,
			want: "aaszz",
		},
	}

	for _, testCase := range testCases {
		actual := getSmallestString(testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
