package leetcode

func uniquePaths(m int, n int) int {
	row := make([]int, n)
	row[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				row[j] = row[j-1] + row[j]
			} else if j > 0 {
				row[j] = row[j-1]
			}
		}
	}
	return row[n-1]
}
