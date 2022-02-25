package _165_compare_version_numbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	r1, r2 := getRevisions(version1), getRevisions(version2)
	for i := 0; i < len(r1) && i < len(r2); i++ {
		if r1[i] < r2[i] {
			return -1
		} else if r1[i] > r2[i] {
			return 1
		}
	}

	if len(r1) < len(r2) {
		return -1
	} else if len(r1) > len(r2) {
		return 1
	}

	return 0
}

func getRevisions(version string) []int {
	parts := strings.Split(version, ".")
	result := make([]int, 0, len(parts))
	lsi := 0
	for i, part := range parts {
		n, _ := strconv.Atoi(part)
		if n > 0 {
			lsi = i
		}
		result = append(result, n)
	}
	return result[:lsi+1]
}
