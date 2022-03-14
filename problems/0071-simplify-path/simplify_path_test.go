package _071_simplify_path

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "/home/",
			want: "/home",
		},
		{
			got:  "/../",
			want: "/",
		},
		{
			got:  "/home//foo/",
			want: "/home/foo",
		},
	}

	for _, testCase := range testCases {
		actual := simplifyPath(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
