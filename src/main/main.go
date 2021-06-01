package main

import (
	"fmt"
	"go_leetcode/src/leetcode"
)

type set map[int]bool
func (s *set)Add(n int){
	(*s)[n]=true
}

func main() {
	fmt.Println(leetcode.ThreeSum1([]int{-1,0,1,2,-1,-4}))

}
