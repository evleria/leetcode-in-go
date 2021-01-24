package _219_contains_duplicate_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		got  []int
		gotK int
		want bool
	}{
		{
			got:  []int{1, 2, 3, 1},
			gotK: 3,
			want: true,
		},
		{
			got:  []int{1, 2, 3, 1},
			gotK: 4,
			want: true,
		},
		{
			got:  []int{1, 2, 3, 1, 2, 3},
			gotK: 2,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := containsNearbyDuplicate(testCase.got, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
