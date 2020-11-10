package linked_list

import "errors"

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

func (list *ListNode) ToSlice() []int {
	slice := make([]int, 0)

	current := list
	for current != nil {
		slice = append(slice, current.Val)
		current = current.Next
	}

	return slice
}

func (list *ListNode) NodeAt(index int) (*ListNode, error) {
	current := list
	for i := 0; i < index; i++ {
		if current.Next == nil {
			return nil, errors.New("out of bound")
		}
		current = current.Next
	}

	return current, nil
}
