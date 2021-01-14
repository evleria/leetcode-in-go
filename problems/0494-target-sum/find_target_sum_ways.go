package _494_target_sum

func findTargetSumWays(nums []int, s int) int {
	const max = 1000
	const l = max*2 + 1
	cur := [l]int{}
	cur[-nums[0]+max] = 1
	cur[nums[0]+max] += 1
	for i := 0; i < len(nums)-1; i++ {
		next := [l]int{}
		for j := 0; j < l; j++ {
			if n := cur[j]; n != 0 {
				next[j+nums[i+1]] += n
				next[j-nums[i+1]] += n
			}
		}
		cur = next
	}

	if s > max {
		return 0
	}

	return cur[s+max]
}
