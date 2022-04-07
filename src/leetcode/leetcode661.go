package leetcode

func imageSmoother(img [][]int) [][]int {
	result := make([][]int, len(img))
	for i := range result {
		result[i] = make([]int, len(img[i]))
	}
	directions := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, +1},
		{0, -1},
		{0, 0},
		{0, +1},
		{+1, -1},
		{+1, 0},
		{+1, +1},
	}
	for x := range img {
		for y := range img[x] {
			grayscale := 0
			count := 0
			for _, d := range directions {
				dx := x + d[0]
				dy := y + d[1]
				if 0 <= dx && dx < len(img) && 0 <= dy && dy < len(img[dx]) {
					count++
					grayscale += img[dx][dy]
				}
			}
			result[x][y] = grayscale / count

		}
	}
	return result
}
