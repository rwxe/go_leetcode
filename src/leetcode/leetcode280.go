package leetcode

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)
	result := make([]int, 0, len(nums))
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		result = append(result, nums[i])
		result = append(result, nums[j])
	}
	copy(nums, result)
}

func wiggleSort_1(nums []int) {
	//TODO:not yet
}
