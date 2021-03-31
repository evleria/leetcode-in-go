package _936_stamping_the_sequence

func movesToStamp(stamp string, target string) []int {
	letters, toErase := []byte(target), len(target)
	contains := func(offset int) bool {
		empty := true
		for i := 0; i < len(stamp); i++ {
			if l := letters[offset+i]; l != '.' && l != stamp[i] {
				return false
			} else if l != '.' {
				empty = false
			}
		}
		return !empty
	}

	erase := func(offset int) {
		for i := 0; i < len(stamp); i++ {
			if letters[offset+i] != '.' {
				letters[offset+i] = '.'
				toErase--
			}
		}
	}

	var result []int
	for toErase != 0 {
		erased := false
		for i := 0; i < len(letters)-len(stamp)+1; i++ {
			if contains(i) {
				erase(i)
				result = append(result, i)
				erased = true
				i += len(stamp)
				if toErase == 0 {
					break
				}
			}
		}
		if !erased {
			return []int{}
		}
	}

	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return result
}
