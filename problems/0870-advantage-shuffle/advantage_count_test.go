package _870_advantage_shuffle

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAdvantageCount(t *testing.T) {
	testCases := []struct {
		gotA []int
		gotB []int
		want []int
	}{
		{
			gotA: []int{2, 7, 11, 15},
			gotB: []int{1, 10, 4, 11},
			want: []int{2, 11, 7, 15},
		},
		{
			gotA: []int{12, 24, 8, 32},
			gotB: []int{13, 25, 32, 11},
			want: []int{24, 32, 8, 12},
		},
	}

	for _, testCase := range testCases {
		actual := advantageCount(testCase.gotA, testCase.gotB)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
