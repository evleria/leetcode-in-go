package _096_unique_binary_search_trees

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumTrees(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  3,
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := numTrees(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}

func BenchmarkNumTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numTrees(19)
	}
}
