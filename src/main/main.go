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
	fmt.Println(leetcode.Subsets0([]int{1,2,3}))
	fmt.Println(leetcode.Subsets1([]int{1,2,3}))
	fmt.Println(leetcode.Subsets2([]int{1,2,3}))

}
