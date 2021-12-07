package _401_binary_watch

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	if turnedOn > 8 {
		return result
	}

	minutes, hours := uint64((1<<6)-1), uint64(((1<<4)-1)<<6)
	var i uint64
	for i = 0; i < (1<<10)-1; i++ {
		if bits.OnesCount64(i) == turnedOn {
			m, h := i&minutes, i&hours>>6
			if m < 60 && h < 12 {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}
