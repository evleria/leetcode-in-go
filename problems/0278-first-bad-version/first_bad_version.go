package _278_first_bad_version

var Bad int

func isBadVersion(version int) bool {
	return version >= Bad
}

func firstBadVersion(n int) int {
	start, end := 1, n
	for start < end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}
