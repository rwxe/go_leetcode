package leetcode

func Generate118(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		result = append(result, make([]int, i+1))
		for j := 0; j < len(result[i]); j++ {
			if j == 0 || j == len(result[i])-1 {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j] + result[i-1][j-1]
			}

		}
	}
	return result

}
