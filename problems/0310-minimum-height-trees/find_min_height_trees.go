package _310_minimum_height_trees

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	for n > 2 {
		leaves := make([]int, 0, n-1)
		for i, list := range graph {
			if len(list) == 1 {
				leaves = append(leaves, i)
			}
		}

		for _, i := range leaves {
			j := graph[i][0]
			graph[j] = remove(graph[j], i)
			graph[i] = nil
			n--
		}
	}

	result := make([]int, 0, n)
	for i, list := range graph {
		if list != nil {
			result = append(result, i)
		}
	}
	return result
}

func remove(nums []int, target int) []int {
	for i, n := range nums {
		if n == target {
			nums[i] = nums[len(nums)-1]
			return nums[:len(nums)-1]
		}
	}
	return nums
}
