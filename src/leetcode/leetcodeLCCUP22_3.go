package leetcode

func BallGame(num int, plate []string) [][]int {
	const (
		UP = iota
		DOWN
		RIGHT
		LEFT
	)
	//di,dj
	D := [][2]int{
		{-1, 0}, //UP
		{+1, 0}, //DOWN
		{0, +1}, //RIGHT
		{0, -1}, //LEFT
	}
	W := map[int]int{
		UP:    LEFT,
		DOWN:  RIGHT,
		LEFT:  DOWN,
		RIGHT: UP,
	}
	E := map[int]int{
		UP:    RIGHT,
		DOWN:  LEFT,
		LEFT:  UP,
		RIGHT: DOWN,
	}
	result := make([][]int, 0)
	// W<- ->E
	playTheGame := func(si, sj int, d int) bool {
		dir := D[d]
		if plate[si][sj] == 'O' {
			return false
		}
		for t := 0; t <= num; t++ {
			if si < 0 ||
				sj < 0 ||
				si >= len(plate) ||
				sj >= len(plate[0]) {
				return false
			}
			if block := plate[si][sj]; block == 'O' {
				return true
			} else if block == 'W' {
				d = W[d]
				dir = D[d]
			} else if block == 'E' {
				d = E[d]
				dir = D[d]
			}
			si += dir[0]
			sj += dir[1]
		}
		return false
	}
	//上边
	for i, j := 0, 1; j < len(plate[0])-1; j++ {
		if playTheGame(i, j, DOWN) {
			result = append(result, []int{i, j})
		}
	}
	//下边
	for i, j := len(plate)-1, 1; j < len(plate[0])-1; j++ {
		if playTheGame(i, j, UP) {
			result = append(result, []int{i, j})
		}
	}
	//左边
	for i, j := 1, 0; i < len(plate)-1; i++ {
		if playTheGame(i, j, RIGHT) {
			result = append(result, []int{i, j})
		}
	}
	//右边
	for i, j := 1, len(plate[0])-1; i < len(plate)-1; i++ {
		if playTheGame(i, j, LEFT) {
			result = append(result, []int{i, j})
		}
	}
	return result

}
