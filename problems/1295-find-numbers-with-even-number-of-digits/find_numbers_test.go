package _295_find_numbers_with_even_number_of_digits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindNumbers(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{12, 345, 2, 6, 7896},
			want:    2,
		},
		{
			gotNums: []int{555, 901, 482, 1771},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := findNumbers(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
