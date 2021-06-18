package _307_range_sum_query_mutable

type NumArray struct {
	arr  []int
	tree []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{}
	tree := make([]int, len(nums)*5)
	buildTree(nums, tree, 0, 0, len(nums)-1)
	res.arr = nums
	res.tree = tree
	return res
}

func (this *NumArray) Update(index int, val int) {
	updateTree(this.tree, 0, 0, len(this.arr)-1, index, val)
	return
}

func (this *NumArray) SumRange(left int, right int) int {
	res := sumTree(this.tree, 0, 0, len(this.arr)-1, left, right)
	return res
}

func sumTree(tree []int, node, start, end, L, R int) int {
	if end < L || start > R {
		return 0
	}

	if start == end {
		return tree[node]
	}

	if L <= start && end <= R {
		return tree[node]
	}

	mid := (start + end) / 2
	leftNode := node*2 + 1
	rightNode := node*2 + 2
	leftSum := sumTree(tree, leftNode, start, mid, L, R)
	rightSum := sumTree(tree, rightNode, mid+1, end, L, R)
	return leftSum + rightSum
}

func updateTree(tree []int, node, start, end, index, val int) {
	if start == end {
		tree[node] = val
		return
	}
	mid := (start + end) / 2
	leftNode := node*2 + 1
	rightNode := node*2 + 2
	if index <= mid {
		updateTree(tree, leftNode, start, mid, index, val)
	} else {
		updateTree(tree, rightNode, mid+1, end, index, val)
	}
	tree[node] = tree[leftNode] + tree[rightNode]
	return
}

func buildTree(arr, tree []int, node, start, end int) {
	if start == end {
		tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	leftNode := node*2 + 1
	rightNode := node*2 + 2
	buildTree(arr, tree, leftNode, start, mid)
	buildTree(arr, tree, rightNode, mid+1, end)
	tree[node] = tree[leftNode] + tree[rightNode]
	return
}
