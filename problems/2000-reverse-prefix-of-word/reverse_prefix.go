package _000_reverse_prefix_of_word

func reversePrefix(word string, ch byte) string {
	w, end := []byte(word), 0
	for i := 0; i < len(w); i++ {
		if w[i] == ch {
			end = i
			break
		}
	}
	for start := 0; start < end; start, end = start+1, end-1 {
		w[start], w[end] = w[end], w[start]
	}
	return string(w)
}
