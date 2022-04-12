package leetcode

func maximalSquare(matrix [][]byte) int {
	min := func(items ...int) int {
		theMin := items[0]
		for i := range items {
			if items[i] < theMin {
				theMin = items[i]
			}
		}
		return theMin
	}
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}

	}
	for i := range dp {
		for j := range dp[i] {
			if i > 0 && j > 0 && dp[i][j] > 0 {
				dp[i][j] = min(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				) + 1
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}
