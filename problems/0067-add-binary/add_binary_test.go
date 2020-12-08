package _067_add_binary

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want string
	}{
		{
			got1: "11",
			got2: "1",
			want: "100",
		},
		{
			got1: "1010",
			got2: "1011",
			want: "10101",
		},
	}

	for _, testCase := range testCases {
		actual := addBinary(testCase.got1, testCase.got2)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got1, testCase.got2)
	}
}
