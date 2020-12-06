package _205_isomorphic_strings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsIsomorphic(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want bool
	}{
		{
			got1: "egg",
			got2: "add",
			want: true,
		},
		{
			got1: "foo",
			got2: "bar",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isIsomorphic(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
