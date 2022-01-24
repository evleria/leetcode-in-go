package _832_flipping_an_image

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		for l, r := 0, len(row)-1; l < r; l, r = l+1, r-1 {
			row[l], row[r] = row[r], row[l]
		}
		for i, n := range row {
			row[i] = (n + 1) % 2
		}
	}

	return image
}
