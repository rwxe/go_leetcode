package leetcode

import "sort"

func bt47(list, nums []int, length int, result *[][]int) {
	if len(list) >= length {
		*result = append(*result, append([]int(nil), list...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		list = append(list, nums[i])
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		bt47(list, newNums, length, result)
		list = list[:len(list)-1]
	}
}
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	length:=len(nums)
	result:=make([][]int,0)
	bt47([]int{},nums,length,&result)
	return result


}
