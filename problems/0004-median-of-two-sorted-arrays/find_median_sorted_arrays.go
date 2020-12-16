package _004_median_of_two_sorted_arrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	lo, hi := 0, len(nums1)
	for lo <= hi {
		partitionX := (lo + hi) / 2
		partitionY := (len(nums1)+len(nums2)+1)/2 - partitionX

		maxLeftX, minRightX, maxLeftY, minRightY := math.MinInt64, math.MaxInt64, math.MinInt64, math.MaxInt64

		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX < len(nums1) {
			minRightX = nums1[partitionX]
		}
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY < len(nums2) {
			minRightY = nums2[partitionY]
		}

		switch {
		case maxLeftX > minRightY:
			hi = partitionX - 1
		case maxLeftY > minRightX:
			lo = partitionX + 1
		default:
			if (len(nums1)+len(nums2))%2 == 0 {
				return (math.Max(float64(maxLeftX), float64(maxLeftY)) + math.Min(float64(minRightX), float64(minRightY))) / 2
			} else {
				return math.Max(float64(maxLeftX), float64(maxLeftY))
			}
		}
	}

	return 0
}
