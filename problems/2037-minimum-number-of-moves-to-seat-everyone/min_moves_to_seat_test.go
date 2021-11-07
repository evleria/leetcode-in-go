package _037_minimum_number_of_moves_to_seat_everyone

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinMovesToSeat(t *testing.T) {
	testCases := []struct {
		gotSeats    []int
		gotStudents []int
		want        int
	}{
		{
			gotSeats:    []int{3, 1, 5},
			gotStudents: []int{2, 7, 4},
			want:        4,
		},
		{
			gotSeats:    []int{4, 1, 5, 9},
			gotStudents: []int{1, 3, 2, 6},
			want:        7,
		},
		{
			gotSeats:    []int{2, 2, 6, 6},
			gotStudents: []int{1, 3, 2, 6},
			want:        4,
		},
	}

	for _, testCase := range testCases {
		actual := minMovesToSeat(testCase.gotSeats, testCase.gotStudents)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
