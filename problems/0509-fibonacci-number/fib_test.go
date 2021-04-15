package _509_fibonacci_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    int
	}{
		{
			gotNums: 2,
			want:    1,
		},
		{
			gotNums: 3,
			want:    2,
		},
		{
			gotNums: 0,
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := fib(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
