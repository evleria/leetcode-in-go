package _461_check_if_a_string_contains_all_binary_codes_of_size_k

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHasAllCodes(t *testing.T) {
	testCases := []struct {
		gotS string
		gotK int
		want bool
	}{
		{
			gotS: "00110110",
			gotK: 2,
			want: true,
		},
		{
			gotS: "00110",
			gotK: 2,
			want: true,
		},
		{
			gotS: "0110",
			gotK: 2,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := hasAllCodes(testCase.gotS, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
