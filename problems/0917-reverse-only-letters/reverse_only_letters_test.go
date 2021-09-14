package _917_reverse_only_letters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseOnlyLetters(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "ab-cd",
			want: "dc-ba",
		},
		{
			got:  "a-bC-dEf-ghIj",
			want: "j-Ih-gfE-dCba",
		},
		{
			got:  "Test1ng-Leet=code-Q!",
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}

	for _, testCase := range testCases {
		actual := reverseOnlyLetters(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
