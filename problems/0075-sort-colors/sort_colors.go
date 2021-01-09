package _075_sort_colors

func sortColors(nums []int) {
	zeroes, twos := 0, len(nums)-1
	for i := 0; i <= twos; {
		switch nums[i] {
		case 0:
			nums[i], nums[zeroes] = nums[zeroes], nums[i]
			zeroes++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[twos] = nums[twos], nums[i]
			twos--
		}
	}
}
