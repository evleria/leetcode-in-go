package _207_course_schedule

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanFinish(t *testing.T) {
	testCases := []struct {
		gotNumCourses    int
		gotPrerequisites [][]int
		want             bool
	}{
		/*{
			gotNumCourses: 2,
			gotPrerequisites: [][]int{
				{1, 0},
			},
			want: true,
		},
		{
			gotNumCourses: 2,
			gotPrerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			want: false,
		},*/
		{
			gotNumCourses: 4,
			gotPrerequisites: [][]int{
				{2, 0},
				{1, 0},
				{3, 1},
				{3, 2},
				{1, 3},
			},
			want: false,
		},
		{
			gotNumCourses: 3,
			gotPrerequisites: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := canFinish(testCase.gotNumCourses, testCase.gotPrerequisites)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
