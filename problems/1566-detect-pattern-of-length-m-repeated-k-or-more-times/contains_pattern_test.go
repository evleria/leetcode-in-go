package _566_detect_pattern_of_length_m_repeated_k_or_more_times

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestContainsPattern(t *testing.T) {
	testCases := []struct {
		gotArr []int
		gotM   int
		gotK   int
		want   bool
	}{
		{
			gotArr: []int{1, 2, 4, 4, 4, 4},
			gotM:   1,
			gotK:   3,
			want:   true,
		},
		{
			gotArr: []int{1, 2, 1, 2, 1, 1, 1, 3},
			gotM:   2,
			gotK:   2,
			want:   true,
		},
		{
			gotArr: []int{1, 2, 1, 2, 1, 3},
			gotM:   2,
			gotK:   3,
			want:   false,
		},
		{
			gotArr: []int{1, 2, 3, 1, 2},
			gotM:   2,
			gotK:   2,
			want:   false,
		},
	}

	for _, testCase := range testCases {
		actual := containsPattern(testCase.gotArr, testCase.gotM, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
