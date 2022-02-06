package _551_student_attendance_record_I

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCheckRecord(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{

		{
			got:  "PPALLP",
			want: true,
		},
		{
			got:  "PPALLL",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := checkRecord(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
