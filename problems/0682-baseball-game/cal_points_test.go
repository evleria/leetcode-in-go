package _682_baseball_game

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCalPoints(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"5", "2", "C", "D", "+"},
			want: 30,
		},
		{
			got:  []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			want: 27,
		},
		{
			got:  []string{"1"},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := calPoints(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
