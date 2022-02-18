package _402_remove_k_digits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveKdigits(t *testing.T) {
	testCases := []struct {
		gotNum string
		gotK   int
		want   string
	}{
		{
			gotNum: "1432219",
			gotK:   3,
			want:   "1219",
		},
		{
			gotNum: "10200",
			gotK:   1,
			want:   "200",
		},
		{
			gotNum: "10",
			gotK:   2,
			want:   "0",
		},
	}

	for _, testCase := range testCases {
		actual := removeKdigits(testCase.gotNum, testCase.gotK)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
