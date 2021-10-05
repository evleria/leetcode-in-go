package _464_maximum_product_of_two_elements_in_an_array

func maxProduct(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			v := (nums[i] - 1) * (nums[j] - 1)
			if v > max {
				max = v
			}
		}
	}
	return max
}
