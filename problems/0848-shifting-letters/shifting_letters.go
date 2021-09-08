package _848_shifting_letters

func shiftingLetters(s string, shifts []int) string {
	bytes, acc := []byte(s), 0
	for i := len(s) - 1; i >= 0; i-- {
		acc += shifts[i]
		bytes[i] = byte((int(bytes[i])-'a'+acc)%26 + 'a')
	}
	return string(bytes)
}
