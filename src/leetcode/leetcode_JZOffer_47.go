package leetcode

func maxValue(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i > 0 && j > 0 {
				grid[i][j] = max(
					grid[i][j]+grid[i-1][j],
					grid[i][j]+grid[i][j-1],
				)
			} else if j > 0 {
				grid[i][j] += grid[i][j-1]

			} else if i > 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
