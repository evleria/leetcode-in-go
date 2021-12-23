package _210_course_schedule_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindOrder(t *testing.T) {
	testCases := []struct {
		gotNumCourses int
		gotP          [][]int
		want          []int
	}{
		{
			gotNumCourses: 2,
			gotP:          [][]int{{1, 0}},
			want:          []int{0, 1},
		},
		{
			gotNumCourses: 4,
			gotP:          [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 1, 2, 3},
		},
		{
			gotNumCourses: 1,
			gotP:          [][]int{},
			want:          []int{0},
		},
	}

	for _, testCase := range testCases {
		actual := findOrder(testCase.gotNumCourses, testCase.gotP)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
