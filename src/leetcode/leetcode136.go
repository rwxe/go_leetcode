package leetcode

import "sort"

func SingleNumber(nums []int) int {
	sort.Ints(nums)
	for i:=0;i<len(nums)-1;{
		if nums[i]!=nums[i+1]{
			return nums[i]
		}
		i+=2
	}
	return nums[len(nums)-1]

}
