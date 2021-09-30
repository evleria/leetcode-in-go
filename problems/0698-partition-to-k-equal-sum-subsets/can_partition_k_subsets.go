package _698_partition_to_k_equal_sum_subsets

func canPartitionKSubsets(nums []int, k int) bool {
	if nums == nil || len(nums) < k {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%k != 0 {
		return false
	}

	visited := make([]bool, len(nums))

	return dfs(nums, 0, len(nums)-1, visited, sum/k, k)
}

func dfs(nums []int, sum, st int, visited []bool, target, k int) bool {
	if k == 0 {
		return true
	}

	if sum == target && dfs(nums, 0, len(nums)-1, visited, target, k-1) {
		return true
	}

	for i := st; i >= 0; i-- {
		if !visited[i] && sum+nums[i] <= target {
			visited[i] = true
			if dfs(nums, sum+nums[i], i-1, visited, target, k) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}
