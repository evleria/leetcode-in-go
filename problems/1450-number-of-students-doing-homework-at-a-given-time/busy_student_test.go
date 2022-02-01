package _450_number_of_students_doing_homework_at_a_given_time

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBusyStudent(t *testing.T) {
	testCases := []struct {
		gotStart []int
		gotEnd   []int
		gotQuery int
		want     int
	}{
		{
			gotStart: []int{1, 2, 3},
			gotEnd:   []int{3, 2, 7},
			gotQuery: 4,
			want:     1,
		},
		{
			gotStart: []int{4},
			gotEnd:   []int{4},
			gotQuery: 4,
			want:     1,
		},
	}

	for _, testCase := range testCases {
		actual := busyStudent(testCase.gotStart, testCase.gotEnd, testCase.gotQuery)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
