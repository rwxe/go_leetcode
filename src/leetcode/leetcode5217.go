package leetcode

import "sort"

func SortJumbled(mapping []int, nums []int) []int {
	transer := func(num int) int {
		result := 0
		multiple := 1
		firstLoop := true
		for num != 0 || firstLoop {
			firstLoop = false
			result += mapping[num%10] * multiple
			num /= 10
			multiple *= 10
		}
		return result
	}
	lessFunc := func(i, j int) bool {
		if transer(nums[i]) == transer(nums[j]) {
			return i < j
		} else {
			return transer(nums[i]) < transer(nums[j])
		}
	}
	sort.Slice(nums, lessFunc)
	return nums

}
