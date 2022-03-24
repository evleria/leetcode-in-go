package _881_boats_to_save_people

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	count, left, right := 0, 0, len(people)-1
	for left <= right {
		count++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}
	return count
}
