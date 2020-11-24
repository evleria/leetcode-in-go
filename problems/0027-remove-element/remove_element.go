package _027_remove_element

func removeElement(nums []int, val int) int {
	out := 0
	for in := 0; in < len(nums); in++ {
		if n := nums[in]; n != val {
			nums[out] = n
			out++
		}
	}

	return out
}
