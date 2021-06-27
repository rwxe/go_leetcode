package leetcode

import "sort"


func productDifference(path []int)int{
	return (path[0]*path[1])-(path[2]*path[3])
}
func bt5797(nums,path []int,max *int){
	if len(path)==4{
		if productDifference(path)>*max{
			*max=productDifference(path)
		}
	}
	for i:=0;i<len(nums);i++{
		path=append(path,nums[i])
		newNums:=append([]int(nil),nums[:i]...)
		newNums=append(newNums,nums[i+1:]...)
		bt5797(newNums,path,max)
		path=path[:len(path)-1]
	}

}

func MaxProductDifference(nums []int) int {
	max:=-1<<31
	bt5797(nums,[]int{},&max)
	return max
}

func MaxProductDifference_2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2]-nums[0]*nums[1]
}
