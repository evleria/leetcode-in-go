package _164_maximum_gap

import "math"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		} else if n < min {
			min = n
		}
	}

	gap := int(math.Ceil(float64(max-min) / float64(len(nums)-1)))
	buckets := make([]*[2]int, len(nums)-1)
	for _, n := range nums {
		if n == min || n == max {
			continue
		}
		idx := (n - min) / gap
		if buckets[idx] == nil {
			buckets[idx] = &[2]int{n, n}
		} else {
			if n < buckets[idx][0] {
				buckets[idx][0] = n
			} else if n > buckets[idx][1] {
				buckets[idx][1] = n
			}
		}
	}

	maxGap, prev := 0, min
	for _, bucket := range buckets {
		if bucket != nil {
			if g := bucket[0] - prev; g > maxGap {
				maxGap = g
			}
			prev = bucket[1]
		}
	}
	if g := max - prev; g > maxGap {
		maxGap = g
	}

	return maxGap
}
