package leetcode

func luckyNumbers(matrix [][]int) []int {
	//[纵坐标,数值]
	theMins := make([][2]int, 0)
	for i := range matrix {
		tempMin := [2]int{0, matrix[i][0]}
		for j := range matrix[i] {
			if matrix[i][j] < tempMin[1] {
				tempMin = [2]int{j, matrix[i][j]}
			}
		}
		theMins = append(theMins, tempMin)
	}
	result := make([]int, 0)
	for _, min := range theMins {
		for i := range matrix {
			if matrix[i][min[0]] > min[1] {
				goto NextLoop
			}
		}
		result = append(result, min[1])
	NextLoop:
	}
	return result
}
