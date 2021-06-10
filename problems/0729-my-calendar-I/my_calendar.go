package _729_my_calendar_I

type Node struct {
	Start int
	End   int
	Left  *Node
	Right *Node
}

type MyCalendar struct {
	head *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	return tryInsert(&c.head, start, end)
}

func tryInsert(head **Node, start, end int) bool {
	switch {
	case *head == nil:
		*head = &Node{Start: start, End: end}
		return true
	case end <= (*head).Start:
		return tryInsert(&((*head).Left), start, end)
	case start >= (*head).End:
		return tryInsert(&((*head).Right), start, end)
	default:
		return false
	}
}
