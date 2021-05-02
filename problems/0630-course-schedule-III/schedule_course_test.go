package _630_course_schedule_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestScheduleCourse(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{
				{100, 200},
				{200, 1300},
				{1000, 1250},
				{2000, 3200},
			},
			want: 3,
		},
		{
			gotNums: [][]int{
				{1, 2},
			},
			want: 1,
		},
		{
			gotNums: [][]int{
				{3, 2},
				{4, 3},
			},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := scheduleCourse(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
