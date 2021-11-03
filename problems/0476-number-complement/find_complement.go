package _476_number_complement

import "math/bits"

func findComplement(num int) int {
	return (1<<(bits.Len(uint(num))) - 1) ^ num
}
