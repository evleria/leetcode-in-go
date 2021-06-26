package _315_count_of_smaller_numbers_after_self

type Node struct {
	Val, Count, Less int
	Left, Right      *Node
}

func countSmaller(nums []int) []int {
	var root *Node
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	insertAll(&root, sorted)

	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = pop(root, n)
	}
	return result
}

func insertAll(rootAddr **Node, nums []int) {
	if len(nums) == 0 {
		return
	}

	mid := len(nums) / 2
	insert(rootAddr, nums[mid])
	insertAll(rootAddr, nums[:mid])
	insertAll(rootAddr, nums[mid+1:])
}

func insert(rootAddr **Node, n int) {
	if *rootAddr == nil {
		*rootAddr = &Node{Val: n, Count: 1}
		return
	}

	root := *rootAddr
	switch {
	case root.Val > n:
		root.Less++
		insert(&(root.Left), n)
	case root.Val < n:
		insert(&(root.Right), n)
	default:
		root.Count++
	}
}

func pop(root *Node, n int) int {
	switch {
	case root.Val > n:
		root.Less--
		return pop(root.Left, n)
	case root.Val < n:
		return root.Count + root.Less + pop(root.Right, n)
	default:
		root.Count--
		return root.Less
	}
}
