package leetcode

func searchMatrix_1(matrix [][]int, target int) bool {
	//m[x][y]>t y-1
	//m[x][y]<t x+1
	for x, y := 0, len(matrix[0])-1; 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[0]); {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
func searchMatrix_0(matrix [][]int, target int) bool {
	bs := func(row, l, r, t int) int {
		for l <= r {
			mid := (l + r) / 2
			if matrix[row][mid] == target {
				return mid
			} else if matrix[row][mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}
		return -1
	}
	for i := range matrix {
		if bs(i, 0, len(matrix[i])-1, target) != -1 {
			return true
		}
	}
	return false
}
