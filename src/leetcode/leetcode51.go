package leetcode

func isValid51(board [][]byte, row, col int) bool {
	// 纵向
	for row := row; row >= 0; row-- {
		if board[row][col] == 'Q' {
			return false
		}
	}
	// 左上方斜向
	for row, col := row-1, col-1; row >= 0 && col >= 0; row, col = row-1, col-1 {
		if board[row][col] == 'Q' {
			return false
		}

	}
	// 右上角斜向
	for row, col := row-1, col+1; row >= 0 && col <= len(board[row]); row, col = row-1, col+1 {
		if board[row][col] == 'Q' {
			return false
		}

	}
	return true
}

func bt51(board [][]byte, row int, result *[][]string) {
	if row >= len(board) {
		return
	}
	for col := 0; col < len(board[col]); col++ {
		if !isValid51(board, row, col) {
			continue
		}
		board[row][col] = 1
		bt51(board, row+1, result)
		board[row][col] = 0

	}

}

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
	}
	result := make([][]string, 0)
	bt51(board, 0, &result)

	return result

}
