package leetcode

func maxPoints(points [][]int) int {
	type linearEquation struct {
		//Ax+By+C=0
		A, B, C int
	}
	gcd := func(m, n int) int {
		for n != 0 {
			m, n = n, m%n
		}
		return m
	}
	genPos := func(p []int) [2]int {
		return [2]int{p[0], p[1]}
	}
	genLinearEquation := func(p1, p2 [2]int) linearEquation {
		//Ax+By+C=0
		//A:y2-y1
		//B:x1-x2
		//C:x2y1-x1y2
		A := p2[1] - p1[1]
		B := p1[0] - p2[0]
		C := p2[0]*p1[1] - p1[0]*p2[1]
		//自动处理符号位，不同顺序p1和p2将会生成同样的方程
		AllGcd := gcd(gcd(A, B), gcd(B, C))
		A /= AllGcd //让ABC互质
		B /= AllGcd
		C /= AllGcd
		return linearEquation{A, B, C}
	}
	n := len(points)
	if n <= 2 {
		return n
	}
	leMap := make(map[linearEquation]map[[2]int]struct{})
	maxCount := 1
	for i := 0; i < n-1; i++ {
		//优化，此时已经获得最大值
		if maxCount >= n-i || maxCount > n/2 {
			break
		}
		for j := i + 1; j < n; j++ {
			pi, pj := genPos(points[i]), genPos(points[j])
			le := genLinearEquation(pi, pj)
			if leMap[le] == nil {
				leMap[le] = make(map[[2]int]struct{})
			}
			leMap[le][pi] = struct{}{}
			leMap[le][pj] = struct{}{}
			if currCount := len(leMap[le]); currCount > maxCount {
				maxCount = currCount
			}
		}
	}
	return maxCount

}

func MaxPoints_2(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	gcd := func(m, n int) int {
		for n != 0 {
			m, n = n, m%n
		}
		return m
	}
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	maxCount := 0
	for i, p := range points {
		//每一轮遍历，count map要重置，否值会将个平行的点一起加入
		count := make(map[[2]int]int)
		//优化，此时已经获得最大值
		if maxCount >= n-i || maxCount > n/2 {
			break
		}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			k := [2]int{y, x}
			count[k]++
			if currCount := count[k] + 1; currCount > maxCount {
				maxCount = currCount
			}
		}
	}
	return maxCount
}
