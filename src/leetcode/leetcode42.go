package leetcode

func Trap42_2(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	water := 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		if lMax < rMax {
			water += lMax - height[l]
			l++
		} else {
			water += rMax - height[r]
			r--
		}
	}
	return water
}
func Trap42_1(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	min := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	lmax := make([]int, len(height))
	rmax := make([]int, len(height))
	lmax[0] = height[0]
	rmax[len(height)-1] = height[len(height)-1]
	water := 0
	for i := 1; i < len(height)-1; i++ {
		lmax[i] = max(height[i], lmax[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = max(height[i], rmax[i-1])
	}
	for i := 1; i < len(height); i++ {
		water += min(lmax[i], rmax[i]) - height[i]
	}
	return water
}
func Trap42_0(height []int) int {
	max := func(items ...int) int {
		maxItem := items[0]
		for i := range items {
			if items[i] > maxItem {
				maxItem = items[i]
			}
		}
		return maxItem
	}
	min := func(items ...int) int {
		maxItem := items[0]
		for i := range items {
			if items[i] < maxItem {
				maxItem = items[i]
			}
		}
		return maxItem
	}
	water := 0
	for i := 1; i < len(height)-1; i++ {
		currWater := min(max(height[:i+1]...), max(height[i:]...)) - height[i]
		//如果当前位置是最高的，则currWater为0
		water += currWater
	}
	return water

}
