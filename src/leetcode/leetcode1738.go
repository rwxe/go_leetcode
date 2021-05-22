package leetcode

import (
	"sort"
)

func getSum(matrix [][]int, a, b int) int {
	sum := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			sum ^= matrix[i][j]
		}

	}
	return sum
}

func KthLargestValue(matrix [][]int, k int) int {
	valueMatrix := make([][]int, len(matrix))
	valueList:=make([]int,0)
	for i := range valueMatrix {
		valueMatrix[i] = make([]int, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// v[i][j]=v[i-1][j-1]^v[i][j-1]^v[i-1][j]
			if i > 0 && j > 0 {
				valueMatrix[i][j] = matrix[i][j] ^ valueMatrix[i-1][j-1] ^ valueMatrix[i][j-1] ^ valueMatrix[i-1][j]
			} else if i == 0 && j > 0 {
				valueMatrix[i][j] = matrix[i][j] ^ valueMatrix[i][j-1]
			} else if i > 0 && j == 0 {
				valueMatrix[i][j] = matrix[i][j] ^ valueMatrix[i-1][j]
			} else if i == 0 && j == 0 {
				valueMatrix[i][j] = matrix[i][j]
			}
			valueList=append(valueList,valueMatrix[i][j])
		}
	}
	sort.Ints(valueList)
	return valueList[len(valueList)-k]

}
