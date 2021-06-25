package leetcode

import ()

func IsVaild37(i, j int, num byte, board [][]byte) bool {
	// 检测同行
	for col := 0; col < len(board[i]); col++ {
		if board[i][col] == num {
			//			fmt.Printf("同行:%q,%q", board[i][col], num)
			return false
		}
	}
	// 检测同列
	for row := 0; row < len(board); row++ {
		if board[row][j] == num {
			//			fmt.Printf("同列:%q,%q", board[row][j], num)
			return false
		}
	}
	// 检测同宫格
	palaceI := i / 3
	palaceJ := j / 3
	for row := palaceI * 3; row < palaceI*3+3; row++ {
		for col := palaceJ * 3; col < palaceJ*3+3; col++ {
			if board[row][col] == num {
				//				fmt.Printf("同宫:%q,%q", board[row][col], num)
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

func SolveSudoku(board [][]byte) {
	defer func() { recover() }()
	bt37(0, 0, board)

}

