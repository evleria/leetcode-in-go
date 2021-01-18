package _057_insert_interval

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		gotIntervals   [][]int
		gotNewInterval []int
		want           [][]int
	}{
		{
			gotIntervals:   [][]int{{0, 3}, {5, 8}, {9, 11}},
			gotNewInterval: []int{9, 18},
			want:           [][]int{{0, 3}, {5, 8}, {9, 18}},
		},

		{
			gotIntervals:   [][]int{},
			gotNewInterval: []int{5, 7},
			want:           [][]int{{5, 7}},
		},
		{
			gotIntervals:   [][]int{{1, 5}},
			gotNewInterval: []int{2, 3},
			want:           [][]int{{1, 5}},
		},
		{
			gotIntervals:   [][]int{{1, 5}},
			gotNewInterval: []int{0, 3},
			want:           [][]int{{0, 5}},
		},
		{
			gotIntervals:   [][]int{{1, 5}},
			gotNewInterval: []int{2, 7},
			want:           [][]int{{1, 7}},
		},
		{
			gotIntervals:   [][]int{{1, 3}, {6, 9}},
			gotNewInterval: []int{2, 5},
			want:           [][]int{{1, 5}, {6, 9}},
		},
		{
			gotIntervals:   [][]int{{1, 3}, {6, 9}},
			gotNewInterval: []int{4, 10},
			want:           [][]int{{1, 3}, {4, 10}},
		},
		{
			gotIntervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			gotNewInterval: []int{4, 8},
			want:           [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, testCase := range testCases {
		actual := insert(testCase.gotIntervals, testCase.gotNewInterval)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
