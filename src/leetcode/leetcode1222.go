package leetcode

// 数组棋盘模拟
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	const BoardSize int = 8
	result := make([][]int, 0, BoardSize)
	board := [BoardSize][BoardSize]bool{}
	for _, q := range queens {
		board[q[0]][q[1]] = true
	}
	//x,y
	dirs := [][]int{
		{0, -1},
		{0, +1},
		{-1, 0},
		{+1, 0},
		{-1, -1},
		{-1, +1},
		{+1, -1},
		{+1, +1},
	}
	for _, d := range dirs {
		kx, ky := king[0], king[1]
		for 0 <= kx && kx < BoardSize &&
			0 <= ky && ky < BoardSize {
			if board[kx][ky] {
				result = append(result, []int{kx, ky})
				break
			}
			kx, ky = kx+d[0], ky+d[1]
		}

	}
	return result
}

// 位图模拟
func queensAttacktheKing_0(queens [][]int, king []int) [][]int {
	const BoardSize int = 8
	const DirsNum int = 8
	var board uint64 = 0 //8*8
	result := make([][]int, 0, DirsNum)
	getPos := func(x, y int) uint64 {
		return 1 << (x*BoardSize + y)
	}

	for _, q := range queens {
		pos := getPos(q[0], q[1])
		board |= pos
	}

	//dx,dy;
	dirs := [...][2]int{
		{0, -1},
		{0, +1},
		{-1, 0},
		{+1, 0},
		{-1, -1},
		{-1, +1},
		{+1, -1},
		{+1, +1},
	}
	for _, d := range dirs {
		kx, ky := king[0], king[1]
		dx, dy := d[0], d[1]
		for 0 <= kx && kx < BoardSize &&
			0 <= ky && ky < BoardSize {
			if board&getPos(kx, ky) != 0 {
				result = append(result, []int{kx, ky})
				break
			}
			kx, ky = kx+dx, ky+dy
		}
	}
	return result
}

// 哈希表
func queensAttacktheKing_1(queens [][]int, king []int) [][]int {
	const DirsNum int = 8
	result := make([][]int, 0, DirsNum)
	//k:vector,v:dist,index of queens
	attackable := make(map[[2]int][2]int, DirsNum)
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	sign := func(x int) int {
		if x > 0 {
			return 1
		} else if x < 0 {
			return -1
		} else {
			return 0
		}
	}
	dist := func(dx, dy int) int {
		return abs(dx)*abs(dx) + abs(dy)*abs(dy)
	}
	vector := func(dx, dy int) [2]int {
		return [2]int{sign(dx), sign(dy)}
	}

	kx, ky := king[0], king[1]
	for i, q := range queens {
		qx, qy := q[0], q[1]
		dx, dy := qx-kx, qy-ky
		if dx == 0 || dy == 0 || abs(dx) == abs(dy) {
			v := vector(dx, dy)
			currDist := dist(dx, dy)
			if prevQ, ok := attackable[v]; !ok || currDist < prevQ[0] {
				attackable[v] = [2]int{currDist, i}
			}
		}
	}
	for _, q := range attackable {
		result = append(result, queens[q[1]])
	}
	return result
}
