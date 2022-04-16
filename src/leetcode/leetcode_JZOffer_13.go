package leetcode

func movingCount_1(m int, n int, k int) int {

	sumK := func(x, y int) int {
		sum := 0
		for ; x != 0; x /= 10 {
			sum += x % 10
		}
		for ; y != 0; y /= 10 {
			sum += y % 10
		}
		return sum
	}
	visitable := make([][]bool, m)
	for x := range visitable {
		visitable[x] = make([]bool, n)
	}
	result := 0
	visitable[0][0] = true
	for x := range visitable {
		for y := range visitable[x] {
			if sumK(x, y) <= k {
				if x-1 >= 0 {
					visitable[x][y] = visitable[x][y] || visitable[x-1][y]
				}
				if y-1 >= 0 {
					visitable[x][y] = visitable[x][y] || visitable[x][y-1]
				}
				if visitable[x][y] {
					result++
				}
			}
		}
	}
	return result

}

// 普通的BFS
func movingCount_0(m int, n int, k int) int {
	sumK := func(x, y int) int {
		sum := 0
		for ; x != 0; x /= 10 {
			sum += x % 10
		}
		for ; y != 0; y /= 10 {
			sum += y % 10
		}
		return sum
	}
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	dirs := [][2]int{
		{+1, 0},
		{-1, 0},
		{0, -1},
		{0, +1},
	}
	result := 0
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	matrix[0][0] = 1
	for len(queue) != 0 {
		currX, currY := queue[0][0], queue[0][1]
		queue = queue[1:]
		result++
		for _, d := range dirs {
			nextX, nextY := currX+d[0], currY+d[1]
			if 0 <= nextX && nextX < m && 0 <= nextY && nextY < n {
				if matrix[nextX][nextY] == 0 {
					if nextK := sumK(nextX, nextY); nextK <= k {
						queue = append(queue, [2]int{nextX, nextY})
						matrix[nextX][nextY] = 1
					}
				}
			}
		}
	}
	return result
}
