package _331_verify_preorder_serialization_of_a_binary_tree

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsValidSerialization(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "1,#,#,#",
			want: false,
		},
		{
			got:  "9,3,4,#,#,1,#,#,2,#,6,#,#",
			want: true,
		},
		{
			got:  "1,#",
			want: false,
		},
		{
			got:  "9,#,#,1",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isValidSerialization(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
