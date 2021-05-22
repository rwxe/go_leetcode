package leetcode

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	n:=len(nums)
	inNums:=make(map[int]bool)
	result:=make([]int,0)
	for _,num:=range nums{
		inNums[num]=true
	}
	for i:=0;i<n;i++{
		if !inNums[i]{
			result=append(result,0)
		}

	}
	return result

}

func FindDisappearedNumbers1(nums []int) []int {
	result:=make([]int,0)
	for i:=0;i<len(nums);i++{
		nums[(nums[i]-1)%len(nums)]+=len(nums)
	}
	fmt.Println(nums)
	for i:=0;i<len(nums);i++{
		if nums[i]<=len(nums){
			result=append(result,i+1)
		}
	}

	return result
}
