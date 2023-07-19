package leetcode

import "math"

func maxDistance_1(arrays [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
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
	result := math.MinInt64
	var maxNum, minNum int
	for i, arr := range arrays {
		first, last := arr[0], arr[len(arr)-1]
		if i == 0 {
			maxNum = last
			minNum = first
			continue
		}
		result = max(result,
			max(abs(maxNum-first), abs(last-minNum)))
		maxNum = max(maxNum, last)
		minNum = min(minNum, first)
	}
	return result
}

func maxDistance(arrays [][]int) int {

	max1 := [2]int{math.MinInt64, 0}
	max2 := [2]int{math.MinInt64, 0}
	min1 := [2]int{math.MaxInt64, 0}
	min2 := [2]int{math.MaxInt64, 0}

	for i, arr := range arrays {
		first, last := arr[0], arr[len(arr)-1]

		if first <= min1[0] {
			min2 = min1
			min1 = [2]int{first, i}
		} else if first <= min2[0] {
			min2 = [2]int{first, i}
		}

		if last >= max1[0] {
			max2 = max1
			max1 = [2]int{last, i}
		} else if last >= max2[0] {
			max2 = [2]int{last, i}
		}
	}
	maxs := [][2]int{max1, max2}
	mins := [][2]int{min1, min2}
	maxDist := math.MinInt64
	for _, max := range maxs {
		for _, min := range mins {
			if max[1] != min[1] {
				if dist := max[0] - min[0]; dist > maxDist {
					maxDist = dist
				}
			}
		}
	}
	return maxDist

}
