package leetcode

import "sort"


func SequentialDigits(low int, high int) []int {
	result:=[]int{}
	for i:=1;i<10;i++{
		num:=i
		for j:=i+1;j<10;j++{
			num=num*10+j
			if low<=num && num<=high{
				result=append(result,num)
			}
		}
	}
	sort.Ints(result)
	return result
	
}
