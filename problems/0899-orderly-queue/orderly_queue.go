package _899_orderly_queue

import (
	"bytes"
	"sort"
)

func orderlyQueue(s string, k int) string {
	b := []byte(s)
	if k == 1 {
		min := []byte(s)
		for i := 0; i < len(s)-1; i++ {
			b = append(b[1:], b[0])
			if bytes.Compare(b, min) < 0 {
				copy(min, b)
			}
		}
		return string(min)
	} else {
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		return string(b)
	}
}
