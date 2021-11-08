package _011_final_value_of_variable_after_performing_operations

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFinalValueAfterOperations(t *testing.T) {
	testCases := []struct {
		gotSeats []string
		want     int
	}{
		{
			gotSeats: []string{"--X", "X++", "X++"},
			want:     1,
		},
		{
			gotSeats: []string{"++X", "++X", "X++"},
			want:     3,
		},
		{
			gotSeats: []string{"X++", "++X", "--X", "X--"},
			want:     0,
		},
	}

	for _, testCase := range testCases {
		actual := finalValueAfterOperations(testCase.gotSeats)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
