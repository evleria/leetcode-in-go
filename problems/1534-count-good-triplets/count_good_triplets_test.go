package _534_count_good_triplets

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountGoodTriples(t *testing.T) {
	testCases := []struct {
		got  []int
		gotA int
		gotB int
		gotC int
		want int
	}{
		{
			got:  []int{3, 0, 1, 1, 9, 7},
			gotA: 7,
			gotB: 2,
			gotC: 3,
			want: 4,
		},
		{
			got:  []int{1, 1, 2, 2, 3},
			gotA: 0,
			gotB: 0,
			gotC: 1,
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := countGoodTriplets(testCase.got, testCase.gotA, testCase.gotB, testCase.gotC)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
