package leetcode

import "fmt"


func sum560(nums []int)int{
	sum:=0
	for _,v:=range nums{
		sum+=v
	}
	return sum
}

func SubarraySum(nums []int, k int) int {
	count:=0
	prefixSum:=make([]int,len(nums)+1)

	for i:=0;i<len(nums);i++{
		prefixSum[i+1]=prefixSum[i]+nums[i]
	}
	fmt.Println(prefixSum)

	for end:=1;end<=len(nums);end++{
		for start:=0;start<end;start++{
			if prefixSum[end]-prefixSum[start]==k{
				count++
			}
		}

	}
	return count

}
