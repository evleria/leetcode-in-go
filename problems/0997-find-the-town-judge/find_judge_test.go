package _997_find_the_town_judge

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindJudge(t *testing.T) {
	testCases := []struct {
		gotN     int
		gotTrust [][]int
		want     int
	}{
		{
			gotN:     2,
			gotTrust: [][]int{{1, 2}},
			want:     2,
		},
		{
			gotN:     3,
			gotTrust: [][]int{{1, 3}, {2, 3}},
			want:     3,
		},
		{
			gotN:     3,
			gotTrust: [][]int{{1, 3}, {2, 3}, {3, 1}},
			want:     -1,
		},
	}

	for _, testCase := range testCases {
		actual := findJudge(testCase.gotN, testCase.gotTrust)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
