package _414_third_maximum_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestThirdMax(t *testing.T) {
	testCases := []struct {
		gotRows []int
		want    int
	}{
		{
			gotRows: []int{1, 1, 2},
			want:    2,
		},
		{
			gotRows: []int{3, 2, 1},
			want:    1,
		},
		{
			gotRows: []int{1, 2},
			want:    2,
		},
		{
			gotRows: []int{2, 2, 3, 1},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := thirdMax(testCase.gotRows)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotRows)
	}
}
