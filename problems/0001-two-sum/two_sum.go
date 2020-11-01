package _001_two_sum

func twoSum(nums []int, target int) []int {

	dict := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := dict[target-n]; ok {
			return []int{j, i}
		}

		dict[n] = i
	}

	return nil
}
