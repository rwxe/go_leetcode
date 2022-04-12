package leetcode

func countSquares(matrix [][]int) int {
	min := func(items ...int) int {
		theMin := items[0]
		for _, v := range items {
			if v < theMin {
				theMin = v
			}
		}
		return theMin
	}
	result := 0
	for i := range matrix {
		for j := range matrix[i] {
			if i > 0 && j > 0 && matrix[i][j] > 0 {
				matrix[i][j] = min(
					matrix[i-1][j],
					matrix[i][j-1],
					matrix[i-1][j-1],
				) + 1
			}
			result += matrix[i][j]
		}

	}
	return result
}
