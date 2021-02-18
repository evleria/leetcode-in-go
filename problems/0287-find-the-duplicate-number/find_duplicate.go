package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	tortoise, hare := nums[0], nums[nums[0]]
	for tortoise != hare {
		tortoise, hare = nums[tortoise], nums[nums[hare]]
	}

	tortoise = 0
	for tortoise != hare {
		tortoise, hare = nums[tortoise], nums[hare]
	}
	return tortoise
}
