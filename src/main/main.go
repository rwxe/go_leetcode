package main

import (
	"fmt"
	"go_leetcode/src/leetcode"
	"sort"
)

type set map[int]bool
func (s *set)Add(n int){
	(*s)[n]=true
}

func main() {
	fmt.Println(leetcode.TwoSum([]int{2,15,11,7},9))
	fmt.Println(leetcode.TwoSum([]int{2,7,11,15},9))
	fmt.Println(leetcode.TwoSum([]int{-10,-1,-18,-19},9))

	//key:=[]int{-10,-1,-18}
	key:=[]int{3,5,1}
	sortedIndex:=[]int{0,1,2}
	fmt.Println("sssss")
	fmt.Println(key,sortedIndex)
	sort.Slice(sortedIndex,func(i,j int)bool{return key[i]<key[j]})
	fmt.Println(key,sortedIndex)
	sort.Ints(key)
	fmt.Println(key,sortedIndex)
}
