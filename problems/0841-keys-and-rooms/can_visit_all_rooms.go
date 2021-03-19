package _841_keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
	open, toOpen := make([]bool, len(rooms)), len(rooms)-1
	open[0] = true

	for keys := append([]int{}, rooms[0]...); len(keys) > 0; keys = keys[1:] {
		k := keys[0]
		if !open[k] {
			open[k], toOpen = true, toOpen-1
			if toOpen == 0 {
				return true
			}
			keys = append(keys, rooms[k]...)
		}
	}
	return toOpen == 0
}
