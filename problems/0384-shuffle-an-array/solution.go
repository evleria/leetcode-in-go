package _384_shuffle_an_array

import "math/rand"

type Solution struct {
	array []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (s *Solution) Reset() []int {
	return s.array
}

func (s *Solution) Shuffle() []int {
	result := make([]int, len(s.array))
	copy(result, s.array)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
