package _291_sequential_digits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSequentialDigits(t *testing.T) {
	testCases := []struct {
		got  int
		gotH int
		want []int
	}{
		{
			got:  100,
			gotH: 300,
			want: []int{123, 234},
		},
		{
			got:  1000,
			gotH: 13000,
			want: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
	}

	for _, testCase := range testCases {
		actual := sequentialDigits(testCase.got, testCase.gotH)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
