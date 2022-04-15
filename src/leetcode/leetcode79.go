package leetcode

//扩展性校弱的DFS
//会修改原来的矩阵
//性能更好
func exist79_4(board [][]byte, word string) bool {
	for x := range board {
		for y := range board[x] {
			if board[x][y] == word[0] {
				if dfs79_4(x, y, 0, word, board) {
					return true
				}
			}
		}
	}
	return false

}
func dfs79_4(cx, cy, cwi int, word string, board [][]byte) bool {
	if 0 <= cx &&
		cx < len(board) &&
		0 <= cy &&
		cy < len(board[cx]) &&
		board[cx][cy] == word[cwi] {
		if cwi == len(word)-1 {
			return true
		}
		original := board[cx][cy]
		board[cx][cy] = byte(0)
		if dfs79_4(cx+1, cy, cwi+1, word, board) ||
			dfs79_4(cx, cy+1, cwi+1, word, board) ||
			dfs79_4(cx-1, cy, cwi+1, word, board) ||
			dfs79_4(cx, cy-1, cwi+1, word, board) {
			return true
		}
		board[cx][cy] = original
	}
	return false
}

//扩展性校弱的DFS
//会修改原来的矩阵
func exist79_3(board [][]byte, word string) bool {
	dirs := [][2]int{
		{-1, 0}, //上
		{+1, 0}, //下
		{0, -1}, //左
		{0, +1}, //右
	}
	var dfs func(cx, cy, cwi int) bool
	dfs = func(cx, cy, cwi int) bool {

		if board[cx][cy] != word[cwi] {
			return false
		} else if cwi == len(word)-1 {
			return true
		}

		original := board[cx][cy]
		board[cx][cy] = byte(0)
		defer func() {
			board[cx][cy] = original
		}()

		for _, d := range dirs {
			nx, ny := cx+d[0], cy+d[1]
			if 0 <= nx && nx < len(board) && 0 <= ny && ny < len(board[nx]) {
				if dfs(nx, ny, cwi+1) {
					return true
				}
			}
		}
		return false
	}
	for x := range board {
		for y := range board[x] {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false

}

//稍微没这么直观的DFS
//visited判断在调用时
func exist79_2(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	dirs := [][2]int{
		{-1, 0}, //上
		{+1, 0}, //下
		{0, -1}, //左
		{0, +1}, //右
	}

	var dfs func(cx, cy, cwi int) bool
	dfs = func(cx, cy, cwi int) bool {
		if board[cx][cy] != word[cwi] {
			return false
		} else if cwi == len(word)-1 {
			return true
		}

		visited[cx][cy] = true
		defer func() {
			visited[cx][cy] = false
		}()

		for _, d := range dirs {
			nx, ny := cx+d[0], cy+d[1]
			if 0 <= nx && nx < len(board) && 0 <= ny && ny < len(board[nx]) {
				if !visited[nx][ny] {
					if dfs(nx, ny, cwi+1) {
						return true
					}
				}
			}
		}
		return false
	}
	for x := range board {
		for y := range board[x] {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false

}

//直观的DFS
//对于每个位置都DFS
func exist79_1(board [][]byte, word string) bool {
	//遍历记录
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	//方向数组
	dirs := [][2]int{
		{-1, 0}, //上
		{+1, 0}, //下
		{0, -1}, //左
		{0, +1}, //右
	}

	var dfs func(cx, cy, cwi int) bool
	dfs = func(cx, cy, cwi int) bool {

		if visited[cx][cy] {
			//如果之前遍历过了，就及时回退
			return false
		}
		visited[cx][cy] = true
		defer func() {
			//如果遍历了当前位置
			//就需要回退时将遍历标志置为false
			visited[cx][cy] = false
		}()

		if board[cx][cy] != word[cwi] {
			// 当前位置不符合要求直接回退
			return false
		} else if cwi == len(word)-1 {
			//来到这里，意味着当前位置符合要求
			//即board[cx][cy]==word[cwi]
			//同时这还是目标字符串的最后一个字符
			//因此可以退出递归返回结果了
			return true
		}

		for _, d := range dirs {
			nx, ny := cx+d[0], cy+d[1]
			if 0 <= nx && nx < len(board) && 0 <= ny && ny < len(board[nx]) {
				if dfs(nx, ny, cwi+1) {
					return true
				}
			}
		}
		return false
	}
	for x := range board {
		for y := range board[x] {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false
}

//直观的DFS
//搜索所有起点再DFS
func exist79_0(board [][]byte, word string) bool {
	//[x0,y0]
	starts := make([][2]int, 0)
	for x := range board {
		for y := range board[x] {
			if board[x][y] == word[0] {
				starts = append(starts, [2]int{x, y})
			}
		}
	}
	dirs := [][2]int{
		{-1, 0}, //上
		{+1, 0}, //下
		{0, -1}, //左
		{0, +1}, //右
	}
	var dfs func(visited map[[2]int]bool, nextCharI, cx, cy, px, py int) bool

	dfs = func(visited map[[2]int]bool, nextCharI, cx, cy, px, py int) bool {

		if visited[[2]int{cx, cy}] {
			return false
		}
		visited[[2]int{cx, cy}] = true

		if nextCharI == len(word) {
			return true
		}

		//fmt.Printf("Into %d %d %d %d %c\n", cx, cy, px, py, word[nextCharI])

		for _, d := range dirs {
			tx, ty := cx+d[0], cy+d[1]

			//fmt.Printf("Try %d %d %d %d %c\n", tx, ty, cx, cy, word[nextCharI])
			if tx != px || ty != py {

				if 0 <= tx && tx < len(board) && 0 <= ty && ty < len(board[tx]) {

					if word[nextCharI] == board[tx][ty] {

						if dfs(visited, nextCharI+1, tx, ty, cx, cy) {
							return true
						}
					}
				}
			}
		}
		visited[[2]int{cx, cy}] = false
		return false

	}
	for _, start := range starts {
		x, y := start[0], start[1]
		visited := make(map[[2]int]bool)
		if dfs(visited, 1, x, y, x, y) {
			return true
		}
	}
	return false
}
