package _089_gray_code

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGrayCode(t *testing.T) {
	testCases := []struct {
		got  int
		want []int
	}{
		{
			got:  2,
			want: []int{0, 1, 3, 2},
		},
		{
			got:  1,
			want: []int{0, 1},
		},
	}

	for _, testCase := range testCases {
		actual := grayCode(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
