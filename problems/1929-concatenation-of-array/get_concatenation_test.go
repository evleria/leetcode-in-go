package _929_concatenation_of_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetConcatenation(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 1},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			got:  []int{1, 3, 2, 1},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}

	for _, testCase := range testCases {
		actual := getConcatenation(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
