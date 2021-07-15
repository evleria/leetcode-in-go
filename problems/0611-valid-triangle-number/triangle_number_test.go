package _611_valid_triangle_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTriangleNumber(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{24, 3, 82, 22, 35, 84, 19},
			want:    10,
		},
		{
			gotNums: []int{2, 2, 3, 4},
			want:    3,
		},
		{
			gotNums: []int{4, 2, 3, 4},
			want:    4,
		},
		{
			gotNums: []int{0, 1, 0},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := triangleNumber(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
