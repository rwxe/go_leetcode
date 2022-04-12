package leetcode

func reachingPoints_3(sx int, sy int, tx int, ty int) bool {
	//反向，辗转相除
	//(tx,ty-tx),(tx-ty,tx)
	//对每轮的较大者取模
	for tx > sx && ty > sy && tx != ty {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	if tx == sx && ty == sy {
		return true
	} else if tx == sx {
		return ty > sy && (ty-sy)%tx == 0
	} else if ty == sy {
		return tx > sx && (tx-sx)%ty == 0
	} else {
		return false
	}

}
func reachingPoints_1(sx int, sy int, tx int, ty int) bool {
	//会爆内存
	//DFS
	var dfs func(currX, currY int) bool
	dfs = func(currX, currY int) bool {
		if currX == tx && currY == ty {
			return true
		} else if currX <= tx && currY <= ty {
			return dfs(currX, currX+currY) || dfs(currX+currY, currY)
		} else {

			return false
		}
	}
	return dfs(sx, sy)
}
func reachingPoints_0(sx int, sy int, tx int, ty int) bool {
	//会爆内存
	//BFS
	//ways:(sx,sy+sx);(sx+sy,sy);
	//back:
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{sx, sy})
	for len(queue) != 0 {
		currX, currY := queue[0][0], queue[0][1]
		queue = queue[1:]
		if currX == tx && currY == ty {
			return true
		} else if currX < tx && currY < ty {
			queue = append(queue, [2]int{currX, currY + currX})
			queue = append(queue, [2]int{currX + currY, currY})
		}
	}
	return false
}
