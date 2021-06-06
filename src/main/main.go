package main

import (
	"fmt"
	"go_leetcode/src/algo"
	"go_leetcode/src/leetcode"
)

type set map[int]bool
func (s *set)Add(n int){
	(*s)[n]=true
}

func main() {
	t1:=&algo.TreeNode{Val:1}
	t1.Left=&algo.TreeNode{Val:2}
	t1.Right=&algo.TreeNode{Val:3}
	t1.Left.Left=&algo.TreeNode{Val:4}
	t1.Left.Right=&algo.TreeNode{Val:5}
	t1.Right.Left=&algo.TreeNode{Val:6}
	t1.Right.Right=&algo.TreeNode{Val:7}
	print(t1,"\n")
	fmt.Println(leetcode.Subset474([]string{"a","b","c"}))
	fmt.Println(leetcode.Subset474([]string{"10", "0001", "111001", "1", "0"}))

}
