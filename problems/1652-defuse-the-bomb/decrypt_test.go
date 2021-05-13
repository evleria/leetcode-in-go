package _652_defuse_the_bomb

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDecrypt(t *testing.T) {
	testCases := []struct {
		gotCode []int
		gotK    int
		want    []int
	}{
		{
			gotCode: []int{5, 7, 1, 4},
			gotK:    3,
			want:    []int{12, 10, 16, 13},
		},
		{
			gotCode: []int{1, 2, 3, 4},
			gotK:    0,
			want:    []int{0, 0, 0, 0},
		},
		{
			gotCode: []int{2, 4, 9, 3},
			gotK:    -2,
			want:    []int{12, 5, 6, 13},
		},
	}

	for _, testCase := range testCases {
		actual := decrypt(testCase.gotCode, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
