package leetcode

func swap2End(ary []int,index int){
	for i:=index;i<len(ary)-1;i++{
		ary[i],ary[i+1]=ary[i+1],ary[i]

	}

}

func MoveZeroes(nums []int) []int {
	left,right:=0,0
	for right<len(nums){
		if nums[right]!=0{
			nums[left],nums[right]=nums[right],nums[left]
			left+=1
		}
		right+=1
	}
	return nums

}
