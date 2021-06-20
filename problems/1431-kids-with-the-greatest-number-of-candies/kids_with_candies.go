package _431_kids_with_the_greatest_number_of_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, n := range candies {
		if n > max {
			max = n
		}
	}

	result := make([]bool, len(candies))
	for i, n := range candies {
		result[i] = n+extraCandies >= max
	}
	return result
}
