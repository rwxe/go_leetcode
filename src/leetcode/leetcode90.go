package leetcode

import "sort"
func SubsetsWithDup(nums []int) [][]int {
	result:=make([][]int,0)
	sort.Ints(nums)
	bt90(nums,[]int{},0,&result)
	return result

}

func bt90(nums,path []int,pos int,result *[][]int) {
	*result=append(*result,append([]int(nil),path...))
	for i:=pos;i<len(nums);i++{
		if i>pos && nums[i]==nums[i-1]{
			print(i,"jump")
			continue
		}
		path=append(path,nums[i])
		bt90(nums,path,i+1,result)
		path=path[:len(path)-1]
	}

}
