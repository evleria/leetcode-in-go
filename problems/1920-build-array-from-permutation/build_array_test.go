package _920_build_array_from_permutation

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBuildArray(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{5, 0, 1, 2, 3, 4},
			want: []int{4, 5, 0, 1, 2, 3},
		},
		{
			got:  []int{0, 2, 1, 5, 3, 4},
			want: []int{0, 1, 2, 4, 5, 3},
		},
	}

	for _, testCase := range testCases {
		actual := buildArray(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
