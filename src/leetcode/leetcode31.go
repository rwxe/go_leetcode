package leetcode

import "sort"

func nextPermutation(nums []int) {
	//right most decr index
	//right most bigger index
	var RMDI, RMBI int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			RMDI = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[RMDI] {
			RMBI = i
			break
		}
	}
	nums[RMBI], nums[RMDI] = nums[RMDI], nums[RMBI]
	if RMDI != RMBI {
		sort.Ints(nums[RMDI+1:])
	} else {
		sort.Ints(nums)
	}
}
