package _131_palindrome_partitioning

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		gotN string
		want [][]string
	}{
		{
			gotN: "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			gotN: "a",
			want: [][]string{{"a"}},
		},
	}

	for _, testCase := range testCases {
		actual := partition(testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
