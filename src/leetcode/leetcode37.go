package leetcode

import ()

func IsVaild37(i, j int, num byte, board [][]byte) bool {
	// 检测同行
	for col := 0; col < len(board[i]); col++ {
		if board[i][col] == num {
			return false
		}
	}
	// 检测同列
	for row := 0; row < len(board); row++ {
		if board[row][j] == num {
			return false
		}
	}
	// 检测同宫格
	palaceI := i / 3
	palaceJ := j / 3
	for row := palaceI * 3; row < palaceI*3+3; row++ {
		for col := palaceJ * 3; col < palaceJ*3+3; col++ {
			if board[row][col] == num {
				return false
			}
		}
	}
	return true
}

func nextPosition(i, j int, board [][]byte) (int, int) {
	for row := i; row < len(board); row++ {
		for col := j; col < len(board[row]); col++ {
			if board[row][col] == '.' {
				return row, col
			}
		}
		j = 0

	}
	return -1, -1
}

func bt37(i, j int, board [][]byte) {
	i, j = nextPosition(i, j, board)
	if i == -1 && j == -1 {
		//fmt.Println(board)
		panic("完成了")
	}
	for num := 1; num < 10; num++ {
		if IsVaild37(i, j, byte(num+'0'), board) {
			board[i][j] = byte(num + '0')
			bt37(i, j, board)
			board[i][j] = '.'
		}
	}
}
func bt37_2(pos int, board [][]byte, space [][]int, row, columns [9][9]bool, block [3][3][9]bool) {
	if pos == len(space) {
		panic("完成了")
	}
	i, j := space[pos][0], space[pos][1]
	for num := 0; num < 9; num++ {
		if row[i][num] || columns[j][num] || block[i/3][j/3][num] {
			continue
		} else {
			row[i][num] = true
			columns[j][num] = true
			block[i/3][j/3][num] = true
			board[i][j] = byte(num + '1')
			bt37_2(pos+1, board, space, row, columns, block)
			board[i][j] = '.'
			block[i/3][j/3][num] = false
			columns[j][num] = false
			row[i][num] = false
		}
	}

}
func SolveSudoku(board [][]byte) {
	defer func() { recover() }()
	bt37(0, 0, board)

}

func SolveSudoku_2(board [][]byte) {
	defer func() { recover() }()
	row := [9][9]bool{}
	columns := [9][9]bool{}
	block := [3][3][9]bool{}
	space := make([][]int, 0)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				space = append(space, []int{i, j})
			} else {
				num := board[i][j] - '1'
				row[i][num] = true
				columns[j][num] = true
				block[i/3][j/3][num] = true
			}
		}
	}
	bt37_2(0,board,space,row,columns,block)

}
