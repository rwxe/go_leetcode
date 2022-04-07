package leetcode

func maxArea(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	area := 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		area = max(area, min(lMax, rMax)*(r-l))
		if lMax < rMax {
			l++
		} else {
			r--
		}
	}
	return area
}
