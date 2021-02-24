package _241_different_ways_to_add_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDiffWaysToCompute(t *testing.T) {
	testCases := []struct {
		got  string
		want []int
	}{
		{
			got:  "2*3-4*5",
			want: []int{-34, -10, -14, -10, 10},
		},
	}

	for _, testCase := range testCases {
		actual := diffWaysToCompute(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
