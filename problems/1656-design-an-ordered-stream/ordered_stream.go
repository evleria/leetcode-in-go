package _656_design_an_ordered_stream

type OrderedStream struct {
	strs []string
	idx  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{strs: make([]string, n)}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	os.strs[idKey-1] = value
	if os.idx == idKey-1 {
		start := os.idx
		for ; os.idx < len(os.strs) && os.strs[os.idx] != ""; os.idx++ {
		}
		return os.strs[start:os.idx]
	}

	return []string{}
}
