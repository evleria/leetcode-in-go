package structs

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromSlice(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	first := &ListNode{
		Val: input[0],
	}

	for current, i := first, 1; i < len(input); i++ {
		next := ListNode{
			Val: input[i],
		}
		current.Next = &next
		current = &next
	}

	return first
}

func ToSlice(list *ListNode) []int {
	slice := make([]int, 0)

	for list != nil {
		slice = append(slice, list.Val)
		list = list.Next
	}

	return slice
}
