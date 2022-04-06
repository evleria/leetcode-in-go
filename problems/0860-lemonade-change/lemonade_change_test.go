package _860_lemonade_change

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLemonadeChange(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{5, 5, 5, 10, 20},
			want: true,
		},
		{
			got:  []int{5, 5, 10, 10, 20},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := lemonadeChange(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
