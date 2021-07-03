package _363_max_sum_of_rectangle_no_larger_than_k

import (
	"math"
	"sort"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	vals := make([]int, n)
	prefixSums := make(sortedIntSet, n)
	bestResult := math.MinInt32

	for startRow := range matrix {
		for i := range vals {
			vals[i] = 0
		}
		for row := startRow; row < m; row++ {
			for col, val := range matrix[row] {
				vals[col] += val
			}

			currSum := 0
			maxSum := vals[0]
			for _, val := range vals {
				currSum = max(currSum+val, val)
				maxSum = max(maxSum, currSum)
			}
			if maxSum <= k {
				bestResult = max(bestResult, maxSum)
				continue
			}

			currSum = 0
			prefixSums = prefixSums[:1]
			prefixSums[0] = 0
			for _, val := range vals {
				currSum += val
				partnerIdx := sort.SearchInts(prefixSums, currSum-k)
				if partnerIdx != len(prefixSums) {
					bestResult = max(bestResult, currSum-prefixSums[partnerIdx])
				}

				prefixSums.insert(currSum)
			}
		}
	}
	return bestResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type sortedIntSet []int

func (s *sortedIntSet) insert(x int) {
	i := sort.SearchInts(*s, x)
	if i == len(*s) {
		*s = append(*s, x)
	} else if (*s)[i] != x {
		*s = append(*s, 0)
		copy((*s)[i+1:], (*s)[i:])
		(*s)[i] = x
	}
}
