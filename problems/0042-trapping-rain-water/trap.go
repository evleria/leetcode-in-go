package _042_trapping_rain_water

func trap(height []int) int {
	left, right, volume := 0, len(height)-1, 0
	maxL, maxR := 0, 0
	for left < right {
		if hl, hr := height[left], height[right]; hl < hr {
			if hl > maxL {
				maxL = hl
			} else {
				volume += maxL - hl
			}
			left++
		} else {
			if hr > maxR {
				maxR = hr
			} else {
				volume += maxR - hr
			}
			right--
		}
	}
	return volume
}
