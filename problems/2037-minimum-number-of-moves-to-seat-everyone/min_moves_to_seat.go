package _037_minimum_number_of_moves_to_seat_everyone

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	moves := 0
	for i := 0; i < len(seats); i++ {
		moves += abs(seats[i] - students[i])
	}
	return moves
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
