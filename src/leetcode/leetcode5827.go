package leetcode

func checkLine(board [][]byte, rMove int, cMove int, color byte, dr, dc int) bool {
	r := rMove + dr
	c := cMove + dc
	step := 1
	for ; 0 <= r && r < 8 && 0 <= c && c < 8; r, c, step = r+dr, c+dc, step+1 {
		if step == 1 {
			if board[r][c] == '.' || board[r][c] == color {
				return false
			}
		} else {
			if board[r][c] == '.' {
				return false
			}
			if board[r][c] == color {
				return true
			}

		}

	}
	return false

}

func CheckMove(board [][]byte, rMove int, cMove int, color byte) bool {
	drs := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	dcs := []int{0, 0, -1, 1, -1, -1, 1, 1}
	for i:=range drs{
		if checkLine(board,rMove,cMove,color,drs[i],dcs[i]){
			return true
		}
	}
	return false

}
