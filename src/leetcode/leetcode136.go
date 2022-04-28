package leetcode

import "sort"

//位运算
func SingleNumber_2(nums []int) int {
	x := 0
	for _, v := range nums {
		x ^= v
	}
	return x
}

//滑动窗口
func SingleNumber_1(nums []int) int {
	window := make([]int, 0)
	showedUp := make(map[int]int)
	for i, v := range nums {
		if _, ok := showedUp[v]; !ok {
			window = append(window, v)
			showedUp[v] = i
		} else {
			showedUp[v] = -1
			for len(window) != 0 && showedUp[window[0]] == -1 {
				window = window[1:]
			}
		}
	}
	return window[0]
}

//排序
func SingleNumber_0(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]

}
