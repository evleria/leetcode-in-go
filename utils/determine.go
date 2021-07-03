package utils

import "sort"

func DetermineSliceOfSlicesOfStrings(data [][]string) [][]string {
	sort.Slice(data, func(i, j int) bool {
		li, lj := len(data[i]), len(data[j])
		if li != lj {
			return li < lj
		}

		sort.Strings(data[i])
		sort.Strings(data[j])

		for k := 0; k < li; k++ {
			if data[i][k] != data[j][k] {
				return data[i][k] < data[j][k]
			}
		}

		return false
	})

	return data
}
