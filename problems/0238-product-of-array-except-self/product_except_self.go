package _238_product_of_array_except_self

func productExceptSelf(nums []int) []int {
	product, zeroIdx := 1, -1
	result := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			if zeroIdx >= 0 {
				return result
			}
			zeroIdx = i
		} else {
			product *= num
		}
	}
	if zeroIdx >= 0 {
		result[zeroIdx] = product
	} else {
		for i := 0; i < len(nums); i++ {
			result[i] = product / nums[i]
		}
	}
	return result
}
