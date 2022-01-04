package _009_complement_of_base_10_integer

import "math/bits"

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	return n ^ ((1 << bits.Len(uint(n))) - 1)
}
