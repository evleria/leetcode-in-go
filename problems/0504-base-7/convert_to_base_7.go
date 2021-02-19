package _504_base_7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var negative bool
	if num < 0 {
		negative = true
		num = -num
	}
	var bytes []byte
	for ; num > 0; num /= 7 {
		bytes = append(bytes, byte(num%7)+'0')
	}
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}
	if negative {
		bytes = append([]byte{'-'}, bytes...)
	}
	return string(bytes)
}
