package _413_minimum_value_to_get_positive_step_by_step_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinStartValue(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{-3, 2, -3, 4, 2},
			want: 5,
		},
		{
			got:  []int{1, 2},
			want: 1,
		},
		{
			got:  []int{1, -2, -3},
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := minStartValue(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
