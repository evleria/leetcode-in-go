package _667_beautiful_arrangement_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConstructArray(t *testing.T) {
	testCases := []struct {
		gotN int
		gotK int
		want []int
	}{
		{
			gotN: 3,
			gotK: 1,
			want: []int{1, 2, 3},
		},
		{
			gotN: 3,
			gotK: 2,
			want: []int{1, 3, 2},
		},
	}

	for _, testCase := range testCases {
		actual := constructArray(testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
