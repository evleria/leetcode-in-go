package _028_implement_strstr

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		gotHaystack string
		gotNeedle   string
		want        int
	}{
		{
			gotHaystack: "hello",
			gotNeedle:   "ll",
			want:        2,
		},
		{
			gotHaystack: "abaaa",
			gotNeedle:   "bba",
			want:        -1,
		},
		{
			gotHaystack: "",
			gotNeedle:   "",
			want:        0,
		},
	}

	for _, testCase := range testCases {
		actual := strStr(testCase.gotHaystack, testCase.gotNeedle)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
