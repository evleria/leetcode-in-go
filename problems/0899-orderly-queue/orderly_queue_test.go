package _899_orderly_queue

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOrderlyQueue(t *testing.T) {
	testCases := []struct {
		gotS string
		gotK int
		want string
	}{
		{
			gotS: "cba",
			gotK: 1,
			want: "acb",
		},
		{
			gotS: "baaca",
			gotK: 3,
			want: "aaabc",
		},
	}

	for _, testCase := range testCases {
		actual := orderlyQueue(testCase.gotS, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
